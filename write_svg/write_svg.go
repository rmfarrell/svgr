package write_svg

import (
  "fmt"
  "bufio"
  _ "image"
  "os"
  "github.com/gographics/imagick/imagick"
  "math/rand"
)

const (
  AdaptiveSharpenVal  float64   = 16
  Funkiness           int       = 6
)

type svg_content struct {
  start, g, end string
}

type pixel_array struct {
  svg_content
  pixel_data  [][]uint8
  w           int
  h           int
  name        string
}

// Constructor
func NewSvgr(imageFiles [][]byte, maxSize int, name string) pixel_array {

  imagick.Initialize()
  defer imagick.Terminate()

  var (
    w           uint
    h           uint
    pixel_data  [][]uint8
  )

  for _, i := range imageFiles {

    wand := imagick.NewMagickWand()

    if err := wand.ReadImageBlob(i); err != nil {
      panic(err.Error())
    }

    w,h = shrink_image(wand, maxSize)

    px, err := wand.ExportImagePixels(0,0,w,h,"RGB", imagick.PIXEL_CHAR)
    if err != nil {
      panic(err.Error())
    }
    pixel_data = append(pixel_data, px.([]uint8))
  }

  return pixel_array {
    svg_content:  writeContainer(w,h),
    pixel_data:   pixel_data,
    w:            int(w),
    h:            int(h),
    name:         name,
  }
}

func (px pixel_array) GetSize() int {

  return len(px.pixel_data)
}

func (px pixel_array) GetName() string {

  return px.name
}

func (px *pixel_array) SetName(name string) {

  px.name = name
  return
}

func (px *pixel_array) Reset() {
  px.svg_content.g = ""
}

/*
* Public Methods
*/

// If no frames are passed, return an array of all frames. Otherwise, do nothing
func normalizeFramesArray(framesIn []int, framesLength int) (framesOut []int) {
  if len(framesIn) >= 1 {
    framesOut = framesIn
  } else {
    for f:=0; f < framesLength; f++ {
      framesOut = append(framesOut, f)
    }
  }
  return
}

func (px *pixel_array) Pixels(frames ...int) {
  // TODO: error/recover from pushing frames that exceed size of pixel_array

  frames = normalizeFramesArray(frames, px.GetSize())

  for _, frame := range frames {
    writeGroup(px, frame, func(rgb []uint8, x,y int) string {
      return fmt.Sprintf(
        "<rect height=\"10px\" width=\"10px\" y=\"%d\" x=\"%d\" fill=\"#%x\"/>",
        y*10, 
        x*10, 
        rgb,
      )
    })
  }
  return
}

func (px *pixel_array) Dots(frames ...int) {

  frames = normalizeFramesArray(frames, px.GetSize())

  for _, frame := range frames {
    writeGroup(px, frame, func(rgb []uint8, x,y int) string {
      return fmt.Sprintf(
        "<circle r=\"5px\" cy=\"%d\" cx=\"%d\" fill=\"#%x\"/>",
        y*10,
        x*10, 
        rgb,
      )
    })
  }
  return
}

func (px *pixel_array) Triangles(frames ...int) (svg string, error error) {

  for _, frame := range normalizeFramesArray(frames, px.GetSize()) {

    writeGroup(px, frame, func(rgb []uint8, x,y int) string {

      x = x*10
      y = y*10
      g := [][]int{}

      if x/10 % 2 == 0 {
        
        // Draw down-pointing triangle
        g = [][]int {
          []int {x-4,y+0},
          []int {x+16,y+0},
          []int {x+6,y+10},
        }

      } else {

        // Draw up-pointing triangle
        g = [][]int {
          []int {x-4,y+10},
          []int {x+16,y+10},
          []int {x+6,y+0},
        }
      }

      return fmt.Sprintf(
        "<polygon points=\"%d,%d %d,%d %d,%d\" fill=\"#%x\"/>",
        g[0][0],
        g[0][1],
        g[1][0],
        g[1][1],
        g[2][0],
        g[2][1],
        rgb,
      )
    })
  }
  return
}

func (px *pixel_array) FunkyTriangles(frames ...int) (svg string, error error) {

  for _, frame := range normalizeFramesArray(frames, px.GetSize()) {

    writeGroup(px, frame, func(rgb []uint8, x,y int) string {

      x = x*10
      y = y*10
      g := [][]int{}
      f := rand.Intn(Funkiness)

      if x/10 % 2 == 0 {
        
        // Draw down-pointing triangle
        g = [][]int {
          []int {x-f,y+f},
          []int {x+16,y},
          []int {x+5+f,y+8+f},
        }

      } else {

        // Draw up-pointing triangle
        g = [][]int {
          []int {x-4,y+10+f},
          []int {x+16,y+10+f},
          []int {x+6+f ,y+0},
        }
      }

      return fmt.Sprintf(
        "<polygon points=\"%d,%d %d,%d %d,%d\" fill=\"#%x\"/>",
        g[0][0],
        g[0][1],
        g[1][0],
        g[1][1],
        g[2][0],
        g[2][1],
        rgb,
      )
    })
  }
  return
}

func (px *pixel_array) FunkySquares(frames ...int) (svg string, error error) {

  for _, frame := range normalizeFramesArray(frames, px.GetSize()) {

    writeGroup(px, frame, func(rgb []uint8, x,y int) string {

      f := rand.Intn(Funkiness)

      x = x*10
      y = y*10

      g := [][]int {
        []int {x+f/2,y-f},
        []int {x+10,y-f},
        []int {x+10-f,y+10},
        []int {x-f,y+10},
      }

      return fmt.Sprintf(
        "<polygon points=\"%d,%d %d,%d %d,%d %d,%d\" fill=\"#%x\"/>",
        g[0][0],
        g[0][1],
        g[1][0],
        g[1][1],
        g[2][0],
        g[2][1],
        g[3][0],
        g[2][1],
        rgb,
      )
    })
  }
  return
}

func (px *pixel_array) Hexagons(frames ...int) (svg string, error error) {

  for _, frame := range normalizeFramesArray(frames, px.GetSize()) {

    writeGroup(px, frame, func(rgb []uint8, x,y int) string {

      z := 0

      if x % 2 == 0 {
        z = 5
      } else {
        z = 0
      }

      x = x*10
      y = y*10

      g := [][]int {
        []int {x+0,y+5+z},
        []int {x+5,y+0+z},
        []int {x+10,y+0+z},
        []int {x+15,y+5+z},
        []int {x+10,y+10+z},
        []int {x+5,y+10+z},
      }

      return fmt.Sprintf(
        "<polygon points=\"%d,%d %d,%d %d,%d %d,%d %d,%d %d,%d\" fill=\"#%x\"/>",
        g[0][0],
        g[0][1],
        g[1][0],
        g[1][1],
        g[2][0],
        g[2][1],
        g[3][0],
        g[3][1],
        g[4][0],
        g[4][1],
        g[5][0],
        g[5][1],
        rgb,
      )
    })
  }
  return
}

func (px *pixel_array) SingleChannel(channelName, color string, opacity float64, scale uint8, offset int, negative bool, frames ...int) {

  channel      := 0
  color_offset := uint8(0)

  switch channelName {
  case "red" :
    channel = 0
  case "green" :
    channel = 1
  case "blue" :
    channel = 2
  }

  if negative {
    color_offset = 0
  } else {
    color_offset = 255
  }


  for _, frame := range normalizeFramesArray(frames, px.GetSize()) {

    writeGroup(px, frame, func(rgb []uint8, x,y int) string {
      return fmt.Sprintf(
        "<circle r=\"%d\" cy=\"%d\" cx=\"%d\" opacity=\"%f\" fill=\"%s\"/>", 
        (color_offset-rgb[channel])/scale, 
        y*10+offset, 
        x*10+offset, 
        opacity, 
        color,
      )
    })
  }
  return
}

func (pxa *pixel_array) Save(dest string) {

  file, err := os.Create(dest)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  contents := pxa.svg_content.start + pxa.svg_content.g + pxa.svg_content.end

  w:= bufio.NewWriter(file)
  w.WriteString(contents)
  w.Flush()
}

func writeContainer(w,h uint) svg_content {
  return svg_content {
    start: fmt.Sprintf(
      "<svg viewbox=\"0 0 %d %d\" xmlns=\"http://www.w3.org/2000/svg\">", 
      w*10, 
      h*10,
    ),
    end: "</svg>",
  }
}

func writeGroup(pxa *pixel_array, frameIndex int, render_method func([]uint8, int, int) string) {

 pxa.svg_content.g += fmt.Sprintf("<g id=\"f%d\">", frameIndex)

  i := 0

  // Iterate over rows
  for row := 0; row < pxa.h; row++ {

    // Iterate over columns
    for col := 0; col < pxa.w; col++ {

      pxa.svg_content.g += render_method(
        []uint8{
          pxa.pixel_data[frameIndex][i], 
          pxa.pixel_data[frameIndex][i+1], 
          pxa.pixel_data[frameIndex][i+2],
        },
        col,
        row,
      )

      i = i+3
    }
  }

  pxa.svg_content.g += "</g>"

  return
}

func shrink_image(wand *imagick.MagickWand, maxSize int) (w,h uint) {

  w,h = get_dimensions(wand)

  shrinkBy := 1

  if w >= h {
    shrinkBy = int(w)/maxSize
  } else {
    shrinkBy = int(h)/maxSize
  }

  wand.AdaptiveResizeImage(
    uint(int(w)/shrinkBy), 
    uint(int(h)/shrinkBy),
  )

  wand.AdaptiveSharpenImage(0,AdaptiveSharpenVal)

  w,h = get_dimensions(wand)

  return
}

func get_dimensions(wand *imagick.MagickWand) (w,h uint) {
  h = wand.GetImageHeight()
  w = wand.GetImageWidth()
  return
}