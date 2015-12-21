package main

import (
  // "fmt"
  // "reflect"
  // "log"
  "os"
  // err "errors"
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

  svgr.Write('./dest/test.svg')
}