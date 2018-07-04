package svgr

import (
	"fmt"
	"image"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/nfnt/resize"
)

const shapeSize = 10

type Input struct {
	Path       string      `json:path`
	Resolution *Resolution `json:resolution`
	ID         string      `json:id`
}

type Resolution struct {
	Width  uint `json:width`
	Height uint `json:height`
}

type Mosaic struct {
	img     image.Image
	id      string
	current *image.Point
	svg     string
	w, h    int
}

type point struct {
	x int
	y int
}

// Mosaic constructor
func NewMosaic(in *Input) (*Mosaic, error) {
	reader, err := os.Open(in.Path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	img = resize.Thumbnail(in.Resolution.Width, in.Resolution.Height, img, resize.Lanczos3)

	// get width and height
	bounds := img.Bounds()
	w := bounds.Max.X
	h := bounds.Max.Y

	// open the svg
	svg := fmt.Sprintf("<svg viewbox=\"0 0 %d %d\" xmlns=\"http://www.w3.org/2000/svg\">", w, h)

	return &Mosaic{img, in.ID, &image.Point{0, 0}, svg, w, h}, nil
}

// Call the render function if there is a next pixel
// Otherwise close the SVG and return
func (m *Mosaic) render(f func() string) string {
	m.svg = fmt.Sprintf("%s%s", m.svg, f())
	if m.next() != nil {
		m.svg = fmt.Sprintf("%s%s", m.svg, f())
		m.render(f)
	}
	return fmt.Sprintf("%s</svg>", m.svg)
}

// Iterate through the image from right to left pixel-by-pixel
// sets the current pixel
func (m *Mosaic) next() *image.Point {
	c := m.current
	if c.X <= m.w-1 {
		m.current.X++
		return m.current
	}
	if c.Y <= m.h-1 {
		m.current.X = 0
		m.current.Y++
		return m.current
	}
	return nil
}

// sample the color of the pixel at m.current
// return the value in SVG-friendly rgba() form
func (m *Mosaic) colorAtCurrent() string {
	sample := m.img.At(m.current.X, m.current.Y)
	r, g, b, a := sample.RGBA()
	return fmt.Sprintf("rgba(%d, %d, %d, %d)", uint8(r), uint8(g), uint8(b), uint8(a))
}
