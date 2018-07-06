# svgr
A dependency-free SVG mosaic generator. Resses down a bitmap image and then crawls it for pixel data, outputting an SVG rendered with various  algorithms

## Usage
```
m, err := svgr.NewMosaic(&svgr.Input{
  Path: "./testdata/doge.png",
  Resolution: &svgr.Resolution{
	  Width:  10,
	  Height: 10,
  },
})
if err != nil {
  t.Error(err)
}
svg := m.Triangles()
b := []byte(svg)
err := ioutil.WriteFile("./out.svg", b, 0644)
if err != nil {
  return err
}
```

## The Constructor
The constructor pulls the image specified by `Path` and shrinks it to the size specified by the `Resolution` in memory.
This virtual image is then used to render SVGs from the below methods.

#### Options
- **Path** path to file
- **Resolution** size of image to be sampled. You really, really want small numbers here :P
-- **Width** max width of sample
-- **Height** max height of sample


## Methods

### Triangles
Outputs interlocking triangles. 
Takes one optional argument which introduces randomization to the rendered pattern.
```
svg := m.Triangles(3)
```


### Hexagons
Outputs interlocking hexagons.
```
svg := m.Hexagons()
```

### Squares
Outputs a grid of squares. 
Takes one optional argument which introduces randomization to the rendered pattern.
```
svg := m.Squares()
```
### Dots
Outputs a grid of dots.
```
svg := m.Dots()
```
### Halftone
A vintage printer halftone effect.
Takes a `ScreenSet` struct which is an array of `Screen`s.
a single `Screen` samples a single channel from the image (red, blue, green, cyan, magenta, yellow, or black).
The output SVG is composed of these screens layered on top of each other (like an old-school printer)

#### Screen Options
- **Color** (string) {"r"|"g"|"b"|"c"|"m"|"y|"k"} the color channel from the original bitmap to be sampled.
- **Lightness** (int) {0-100} the tint of the output channel
- **Saturation** (int) {0-100} the saturation of the output channel
- **Opacity** (int) {0-100} the opacity of the output channel
- **OffsetX** (int) {0-100} offset the output channel horizontally by an `n` percent
- **OffsetY** (int) {0-100} offset the output channel vertically by an `n` percent
- **HueShift** (int) shift the output hue by `n`(+/-) color values.

```
svg, err := m.Halftone(svgr.ScreenSet{
  &svgr.Screen{
	  Color:      "r",
		Saturation: 50,
		Lightness:  50,
	},
	&svgr.Screen{
		Color:      "g",
		Saturation: 50,
		Lightness:  50,
		},
	&svgr.Screen{
		Color:      "b",
		Saturation: 50,
		Lightness:  50,
	},
})
```

TODO: Examples
