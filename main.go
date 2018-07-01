package svgr

import (
  "fmt"
  "bufio"
  _ "image"
  "os"
  "github.com/gographics/imagick/imagick"
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