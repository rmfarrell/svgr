package svgr

import (
	"fmt"
	"image/color"
)

type Screen struct {

	// must be one of 'r,g,b,a,c,m,y,k"
	Color string `json:color`

	// must be >= 0 and <= 100
	Lightness int `json:lightness`

	// must be >= 0 and <= 100
	Saturation int `json:saturation`

	// must be >= 0 and <= 100
	Opacity int `json:opacity`

	// Offset the output dot to be cool
	// must be >= 0 and <= 100
	OffsetX int `json:offsetx`
	OffsetY int `json:offsety`

	// Shift the hue on the output dot
	HueShift int `json:hueshift`
}

type ScreenSet []*Screen

type hsla struct {
	hue        int
	saturation int
	lightness  int
	alpha      float64
}

var (
	allowedColors                = []string{"r", "g", "b", "a", "c", "m", "y", "k"}
	hues          map[string]int = map[string]int{
		"r": 0,
		"g": 120,
		"b": 240,
		"c": 180,
		"m": 300,
		"y": 60,
		"k": 0,
	}
)

func (m *Mosaic) Halftone(sn ScreenSet) (string, error) {

	var isCMYK bool = false

	// -- Normalize input
	for _, c := range sn {

		// color is allowed
		isAllowed := false
		for _, ac := range allowedColors {
			if c.Color == ac {
				// flag CMYK
				if c.Color == "c" || c.Color == "m" || c.Color == "y" || c.Color == "k" {
					isCMYK = true
				}
				isAllowed = true
				break
			}
		}
		if isAllowed == false {
			return "", fmt.Errorf("%s is not allowed for Color. Must be one of r,b,a,c,m,y,k", c.Color)
		}

		// Lightness and Saturation are between 0 and 1
		err := validateInRange(c.Lightness, 0, 100)
		err = validateInRange(c.Opacity, 0, 100)
		err = validateInRange(c.Saturation, 0, 100)
		if err != nil {
			return "", err
		}

		// -- Assign Default opacity
		if c.Opacity == 0 {
			c.Opacity = 33
		}
	}

	return m.render(func() string {

		var out string
		var h = hsla{}
		var size float64
		var r, g, b, a uint32
		var c, ma, y, k uint8

		// sample the pixel
		sample := m.img.At(m.current.X, m.current.Y)
		r, g, b, a = sample.RGBA()

		// get CMYK color if necessary
		if isCMYK {
			c, ma, y, k = color.RGBToCMYK(uint8(r), uint8(g), uint8(b))
		}

		for _, screen := range sn {
			switch screen.Color {
			case "r":
				h.hue = hues["r"]
				size = calcSize(uint8(r))
			case "g":
				h.hue = hues["g"]
				size = calcSize(uint8(g))
			case "b":
				h.hue = hues["b"]
				size = calcSize(uint8(b))
			case "c":
				h.hue = hues["c"]
				size = calcSize(c)
			case "m":
				h.hue = hues["m"]
				size = calcSize(ma)
			case "y":
				h.hue = hues["y"]
				size = calcSize(y)
			case "k":
				h.hue = hues["k"]
				size = calcSize(k)
			}
			h.hue = h.hue + screen.HueShift
			h.saturation = screen.Saturation
			h.lightness = screen.Lightness
			h.alpha = float64(screen.Opacity) * 0.01 * float64(uint8(a)/255)

			// handle offset
			x := offsetCoordByPercent(m.current.X, screen.OffsetX)
			y := offsetCoordByPercent(m.current.Y, screen.OffsetY)

			out = fmt.Sprintf("%s%s",
				out,
				screen.renderDot(&h, size, x, y),
			)
		}

		return out

	}), nil
}

func offsetCoordByPercent(in, pct int) float64 {
	return float64(in)*shapeSize + shapeSize*0.5 + float64(pct)/100*shapeSize
}

// return the size of a dot based on r/g/b/c/m/y/k value
func calcSize(in uint8) float64 {
	return float64(in) / 255 * shapeSize * 0.5
}

func (s *Screen) renderDot(h *hsla, size, x, y float64) string {
	return fmt.Sprintf(
		"<circle r=\"%.2f\" cx=\"%.2f\" cy=\"%.2f\" fill=\"%s\"/>",
		size,
		x,
		y,
		h.toString(),
	)
}

func (h *hsla) toString() string {
	return fmt.Sprintf(
		"hsla(%d,%d%%,%d%%,%.2f)",
		h.hue,
		h.saturation,
		h.lightness,
		h.alpha,
	)
}

func validateInRange(in, min, max int) error {
	if in < min || in > max {
		return fmt.Errorf("value must be between %d and %d (inclusive). Received %d", min, max, in)
	}
	return nil
}
