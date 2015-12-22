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
  reader, err := os.Open("src/baby.png")
  if err != nil {
    panic(err.Error())
  }
  defer reader.Close()

  svgr := svgr.NewSvgr(reader, 60, "steve_harvey")
  svgr.SingleChannel(
    "blue",
    "#3c9cf0", //color value
    .7,        //opacity
    35,        //scale
    -8,        //offset
    false,     //negative
  )
  svgr.SingleChannel("red", "#f03c3c", .6, 50, 0, false)
  svgr.SingleChannel("green", "#63f03c", .4, 50, 6, false)
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
}