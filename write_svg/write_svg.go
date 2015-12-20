package write_svg

import (
  "fmt"
  "bufio"
  _ "image"
  "os"
  "github.com/gographics/imagick/imagick"
)

type pixel_array struct {
  pixel_data []uint8
  w          int
  h          int
}

func NewSvgr(img *os.File) pixel_array {

  imagick.Initialize()
  defer imagick.Terminate()

  wand := imagick.NewMagickWand()

  wand.ReadImageFile(img)

  w,h := get_dimensions(wand)

  wand.AdaptiveResizeImage(w/20, h/20)

  w,h = get_dimensions(wand)

  wand.AdaptiveSharpenImage(0,16)

  pixel_data, err := wand.ExportImagePixels(0,0,w,h,"RGB", imagick.PIXEL_CHAR)
  if err != nil {
    panic(err.Error())
  }

  return pixel_array {
    pixel_data: pixel_data.([]uint8),
    w: int(w),
    h: int(h),
  }
}

func (px pixel_array) Write(dest string) (svg string, error error) {

  svg = fmt.Sprintf("<svg width=\"%d\" height=\"%d\"  xmlns=\"http://www.w3.org/2000/svg\"><g>", px.w*10, px.h*10)

  i := 0

  // Iterate over rows
  for row := 0; row < px.h; row++ {

    // Iterate over columns
    for col := 0; col < px.w; col++ {

      r:= px.pixel_data[i]
      g:= px.pixel_data[i+1]
      b:= px.pixel_data[i+2]

      i = i+3
      
      svg += fmt.Sprintf("<rect height=\"10\" width=\"10\" y=\"%d\" x=\"%d\" fill=\"rgb(%d,%d,%d)\"/>", row*10, col*10, r, g, b)    

    }
  }

  svg += "</g></svg>"

  write_file(svg, dest)

  return
}




// Private Methods

func get_dimensions(wand *imagick.MagickWand) (w,h uint) {
  h = wand.GetImageHeight()
  w = wand.GetImageWidth()
  return
}

func newPixelArray(pixel_data []uint8, w,h int) pixel_array {
  return pixel_array {
    pixel_data: pixel_data,
    w: w,
    h: h,
  }
}

func write_file(contents string, dest string) {

  file, err := os.Create(dest)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  w:= bufio.NewWriter(file)
  w.WriteString(contents)
  w.Flush()

}



func gather_rgb_values(pixels []uint8) [][]int {

  rgb_array := [][]int{}

  for p := 0; p < len(pixels); p=p+3 {

    // pa := make([]int, 3)

    pa := []int{
      int(pixels[p]),
      int(pixels[p+1]),
      int(pixels[p+2]),
    }

    rgb_array = append(rgb_array, pa) 
  }

  return rgb_array
}

func gather_green_values() {

}

func gather_red_values() {

}

func gather_blue_values() {

}