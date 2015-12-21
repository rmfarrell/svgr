package main

import (
  // "fmt"
  // "reflect"
  "os"
  svgr "./write_svg"
)

const (
  dest  string  = "./dest/"
  src   string  = "./src/"
)

func main() {

  // Open the image
  reader, err := os.Open("src/img.jpg")
  if err != nil {
    panic(err.Error())
  }
  defer reader.Close()

  svgr := svgr.NewSvgr(reader, 30, "steve_harvey")

  svgr.MultiChannel()
  svgr.Save(dest + svgr.GetName() + "_multichannel.svg")
  svgr.Reset()

  svgr.FunkyTriangles()
  svgr.Save(dest + svgr.GetName() + "_funky_triangles.svg")
  svgr.Reset()

  svgr.Triangles()
  svgr.Save(dest + svgr.GetName() + "_triangles.svg")
  svgr.Reset()

  svgr.Dots()
  svgr.Save(dest + svgr.GetName() + "_dots.svg")
  svgr.Reset()

  svgr.Pixels()
  svgr.Save(dest + svgr.GetName() + "_pixels.svg")
  svgr.Reset()

  svgr.FunkySquares()
  svgr.Save(dest + svgr.GetName() + "_funky_squares.svg")
  svgr.Reset()

  svgr.Hexagons()
  svgr.Save(dest + svgr.GetName() + "_hexagons.svg")
  svgr.Reset()



  // svgr.Save(dest + "/" + svgr.name + ".svg")
  // svgr.FunkyTriangles(name + "_funky_triangles", dest)
  // svgr.Triangles(name + "_triangles", dest)
  // svgr.Dots(name + "_dots", dest)
  // svgr.Pixels(name + "_pixels", dest)
  // svgr.FunkySquares(name + "_funky_squares", dest)
  // svgr.Hexagons(name + "_hexagons", dest)
}