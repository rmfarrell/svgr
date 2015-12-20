package main

import (
  // "fmt"
  // "reflect"
  _ "image"
  // "log"
  "os"
  // err "errors"
  "./write_svg"
  "github.com/gographics/imagick/imagick"
)

const (
  dest  = "./dest/"
  src   = "./src/"
)

func main() {

  // Open the image
  reader, err := os.Open("src/indian.jpg")
  if err != nil {
    panic(err.Error())
  }
  defer reader.Close()

  // decoded_img, _, _ := image.Decode(reader)

  imagick.Initialize()
  defer imagick.Terminate()

  wand := imagick.NewMagickWand()

  wand.ReadImageFile(reader)

  w,h := get_dimensions(wand)

  wand.AdaptiveResizeImage(w/20, h/20)

  w,h = get_dimensions(wand)

  wand.AdaptiveSharpenImage(0,16)

  // Send image data to write_svg
  pixel_data, err := wand.ExportImagePixels(0,0,w,h,"RGB", imagick.PIXEL_CHAR)
  if err != nil {
    panic(err.Error())
  }

  write_svg.Write(pixel_data.([]uint8),int(w),int(h))
}

func get_dimensions(wand *imagick.MagickWand) (w,h uint) {
  h = wand.GetImageHeight()
  w = wand.GetImageWidth()
  return
}

func save_image() {

  // write the image
  toimg, err := os.Create("dessafdt/nesw.jpg")
  if err != nil {
    panic(err.Error())
  }
  defer toimg.Close()
}