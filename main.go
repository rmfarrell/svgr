package main

import (
  "fmt"
  // "reflect"
  "os"
  "os/exec"
  "github.com/gographics/imagick/imagick"
  "github.com/satori/go.uuid"
  "io/ioutil"
  svgr "./write_svg"
)

const (
  dest  string  = "./dest/"
  src   string  = "./src/"
)

func main() {

  // Open the image
  reader := readFile("./src/animated.gif")

  // Store each file in memory so we get access to each frame of the animated gif
  imgFiles, imgPath := separateAnimatedGif(reader)

  // Delete the directory once we have stored the image files in memory
  if err := os.RemoveAll(imgPath); err != nil {
    panic(err.Error())
  }

  svg := svgr.NewSvgr(imgFiles, 60, "my_svg")

  svg.Pixels()
  svg.Save(dest + svg.GetName() + "pixels.svg")

  /*

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
  */
}



/*
* Utities
*/

func videoToAnimatedGif(video *os.File) *os.File {
  return video
}

func separateAnimatedGif(animated *os.File) (imageFiles [][]byte, dir string) {

  // Generate a UUID and make a directory with corresponding name
  dir = fmt.Sprintf("./_%s", uuid.NewV4())
  if err := os.Mkdir(dir, 0777); err != nil {
    panic(err.Error())
  }

  // Separate and coalesce each frame of the animation into the new folder
  cmd := exec.Command(
    "convert", 
    "-coalesce", 
    animated.Name(), 
    fmt.Sprintf("./%s/image_%%d.gif", dir),
  )
  cmd.Run()

  // Save a reference to each file in the directory
  files, _ := ioutil.ReadDir(dir)
  for _, f := range files {
    rf, _ := ioutil.ReadFile(dir + "/" + f.Name())
    imageFiles = append(imageFiles, rf)
  }

  return
}

func readFile(file string) *os.File {
  reader, err := os.Open(file)
  if err != nil {
    panic(err.Error())
  }
  defer reader.Close()
  return reader
}

func _createImageSeqence(file *os.File) {

  wand := imagick.NewMagickWand()

  wand.ReadImageFile(file)

  fmt.Println(wand)
}