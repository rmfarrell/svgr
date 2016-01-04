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
  reader := readFile("./src/lemmy_guitar.gif")

  // Store each file in memory so we get access to each frame of the animated gif
  imgFiles, imgPath := separateAnimatedGif(reader)

  // Delete the directory once we have stored the image files in memory
  if err := os.RemoveAll(imgPath); err != nil {
    panic(err.Error())
  }

  lwf := svgr.NewSvgr(imgFiles, 80, "lemmy_guitar")

  // for x:=0; x < len(imgFiles); x++ {
  //   lwf.SingleChannel("red", "#f03c3c", .6, 50, 0, false, x)
  //   lwf.SingleChannel("blue", "#3c9cf0", .6, 50, -8, false, x)
  //   lwf.SingleChannel("green", "#63f03c", .4, 50, 6, false, x)
  // }

  lwf.FunkyTriangles()

  lwf.Save(dest + lwf.GetName() + ".svg")
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
    fmt.Sprintf("./%s/image_%%03d.gif", dir),
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