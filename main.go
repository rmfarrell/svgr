package main

import (
  "fmt"
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

  fmt.Println(svgr)
}



// func save_image() {

//   // write the image
//   toimg, err := os.Create("dessafdt/nesw.jpg")
//   if err != nil {
//     panic(err.Error())
//   }
//   defer toimg.Close()
// }