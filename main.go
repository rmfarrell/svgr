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
  reader, err := os.Open("src/img.jpg")
  if err != nil {
    panic(err.Error())
  }
  defer reader.Close()

  svgr := svgr.NewSvgr(reader)

  name := "steve_harvey"
  dest := "./dest"

  svgr.FunkyTriangles(name + "_funky_triangles", dest)
  svgr.Triangles(name + "_triangles", dest)
  svgr.Dots(name + "_dots", dest)
  svgr.Pixels(name + "_pixels", dest)
  svgr.FunkySquares(name + "_funky_squares", dest)
  svgr.Hexagons(name + "_hexagons", dest)
}