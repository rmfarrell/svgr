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
  //Apply adaptive sharpening to shrunk images
  AdaptiveSharpenVal  float64   = 16 
  //Amount of randomness to apply to "Funky" pixelation methods
  Funkiness           int       = 6 
)

// The actual body of the svg output
type svgContent struct {
  start, g, end string
}

type pixelArray struct {
  svgContent
  // An array of array of exported pixels.
  // Each parent array is a frame in an animation (a single frame is) pixelData[0]
  // The child arrays represent the rgb value of every pixel of the image
  pixelData   [][]uint8
  // Height and width of the image(s)
  w,h         int
  name        string
}

/*
* Constructor
*/

/* 
* Shrinks the image to its maxSize, and samples each pixels in the image array
* @param imageFiles {[][]byte} expects an array of image blobs
* @param maxSize {int} is the maximum size length of longest size
* @param name {string}
* @return pixelArray {struct}
*/
func NewSvgr(imageFiles [][]byte, maxSize int, name string) pixelArray {

  imagick.Initialize()
  defer imagick.Terminate()

  var (
    w,h         uint
    pixelData   [][]uint8
  )

  for _, i := range imageFiles {

    wand := imagick.NewMagickWand()

    if err := wand.ReadImageBlob(i); err != nil {
      panic(err.Error())
    }

    // maxSize is the longest size
    w,h = shrinkImage(wand, maxSize)

    px, err := wand.ExportImagePixels(0,0,w,h,"RGB", imagick.PIXEL_CHAR)
    if err != nil {
      panic(err.Error())
    }
    pixelData = append(pixelData, px.([]uint8))
  }

  return pixelArray {
    svgContent:  writeContainer(w,h),
    pixelData:   pixelData,
    w:            int(w),
    h:            int(h),
    name:         name,
  }
}

/*
* Public Methods
*/

// Get the number of frames in the pixelArray
func (px pixelArray) GetSize() int {

  return len(px.pixelData)
}

func (px pixelArray) GetName() string {
  return px.name
}

func (px *pixelArray) SetName(name string) {
  px.name = name
  return
}

// Reset the content of the output svg, except for the opening <svg> tags 
func (px *pixelArray) Reset() {
  px.svgContent.g = ""
}

// The following funcs output svg's drawn in various styles

/*
* Draw a <g> from the image as a field of Squares
* @param frames {...int} a comma-separated list of frames from the array to render
* a blank value loops through every image in the pixelArray.pixelData array.
*/
func (px *pixelArray) Pixels(frames ...int) {
  // TODO: error/recover from pushing frames that exceed size of pixelArray

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

// Draw the image as a field of circles
func (px *pixelArray) Dots(frames ...int) {

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

// Draw the image as a field of interlocking triangles
func (px *pixelArray) Triangles(frames ...int) (svg string, error error) {

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

// Draw the image as a field of interlocking triangles but with random angles thrown in.
// Warning: funky
func (px *pixelArray) FunkyTriangles(frames ...int) (svg string, error error) {

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

// Draw the image as a field of rectangles with random angles thrown in.
func (px *pixelArray) FunkySquares(frames ...int) (svg string, error error) {

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

// Draw the image as a field of interlocking hexagons
func (px *pixelArray) Hexagons(frames ...int) (svg string, error error) {

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

/*
* Draw the image as a dot matrix representing a single color channel (ie r,g, or b)
* Combine multiple channels in the by calling this in for loop with the iterator pointing at the same frame
* @param channelName {string} acceptable values "red", "green", "blue"
* @param color {string} hex value of the output color (eg "#63f03c")
* @param opacity {float64} opacity of the output layer. Useful for drawing a multiple channels
* @param scale {uint8} increase or decrease the size of the dot.
* @param offset {int} shift position of the dot matrix to replicate a screen print effect
* @param negative {boolean} set to true to get the negative values from the input. A
*   false: represents the positive value of the r,g, or b channel from the input. 
*   true:  represents the negative value.
* @param frames {...int} selectively apply this treatment to a comma separated list of indeces of images. See above.
*/
func (px *pixelArray) SingleChannel(channelName, color string, opacity float64, scale uint8, offset int, negative bool, frames ...int) {

  channel      := 0
  colorOffset := uint8(0)

  switch channelName {
  case "red" :
    channel = 0
  case "green" :
    channel = 1
  case "blue" :
    channel = 2
  }

  if negative {
    colorOffset = 0
  } else {
    colorOffset = 255
  }

  for _, frame := range normalizeFramesArray(frames, px.GetSize()) {

    writeGroup(px, frame, func(rgb []uint8, x,y int) string {
      return fmt.Sprintf(
        "<circle r=\"%d\" cy=\"%d\" cx=\"%d\" opacity=\"%f\" fill=\"%s\"/>", 
        (colorOffset-rgb[channel])/scale, 
        y*10+offset, 
        x*10+offset, 
        opacity, 
        color,
      )
    })
  }
  return
}

// Save the svg output as a .svg to the destination specified in @param dest
func (pxa *pixelArray) Save(dest string) {

  fmt.Printf("Saving %s.svg...", pxa.GetName(),)

  file, err := os.Create(dest + ".svg")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  contents := pxa.svgContent.start + pxa.svgContent.g + pxa.svgContent.end

  w:= bufio.NewWriter(file)
  w.WriteString(contents)
  w.Flush()

  fmt.Println("success!")
}


/*
* Private Methods
*/

// populates the opening <svg> tags of the svg output.
func writeContainer(w,h uint) svgContent {
  return svgContent {
    start: fmt.Sprintf(
      "<svg viewbox=\"0 0 %d %d\" xmlns=\"http://www.w3.org/2000/svg\">", 
      w*10, 
      h*10,
    ),
    end: "</svg>",
  }
}

// Filter which defaults to the length of the image array if no values are passed.
// If no frames are passed, return an array of all frames. Otherwise, do nothing.
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

/* 
* Write a <g> in the output svg, based on a per-pixel drawing method from its callee.
* @param pxa {*pixelArray}
* @param frameIndex {int} points at the image in the pixelData array
* @param renderMethod {func} calls a function which outputs an svg drawing method 
*   on a each pixel using its rgb value
*/
func writeGroup(pxa *pixelArray, frameIndex int, renderMethod func([]uint8, int, int) string) {

  fmt.Printf("writing <g> %d...", frameIndex + 1)
  
  pxa.svgContent.g += fmt.Sprintf("<g id=\"f%d\">", frameIndex)

  i := 0

  // Iterate over rows
  for row := 0; row < pxa.h; row++ {

    // Iterate over columns
    for col := 0; col < pxa.w; col++ {

      // Iterate through each pixel of the image (input) and call the renderMethod on it.
      pxa.svgContent.g += renderMethod(
        []uint8{
          pxa.pixelData[frameIndex][i], 
          pxa.pixelData[frameIndex][i+1], 
          pxa.pixelData[frameIndex][i+2],
        },
        col,
        row,
      )

      i = i+3
    }
  }

  pxa.svgContent.g += "</g>"

  fmt.Println("success!")

  return
}

// Shrink an image so that its longest dimension is no longer than maxSize
func shrinkImage(wand *imagick.MagickWand, maxSize int) (w,h uint) {

  w,h = getDimensions(wand)

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

  // Sharpen the image to bring back some of the color lost in the shrinking
  wand.AdaptiveSharpenImage(0,AdaptiveSharpenVal)

  w,h = getDimensions(wand)

  return
}

// Returns an the width and height of magick wand
func getDimensions(wand *imagick.MagickWand) (w,h uint) {
  h = wand.GetImageHeight()
  w = wand.GetImageWidth()
  return
}