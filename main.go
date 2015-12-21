package main

import (
  // "fmt"
  // "reflect"
  "os"
  svgr "./write_svg"
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

  svgr := svgr.NewSvgr(reader)

  svgr.Hexagons("./dest/test.svg")
}