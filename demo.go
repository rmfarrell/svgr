package main

import (
  "fmt"
  // "reflect"
  "os"
  "os/exec"
  "github.com/satori/go.uuid"
  "io/ioutil"
  svgr "./svgr"
)

const (
  dest  string  = "./dest/"
  src   string  = "./src/"
)

var file string = "./src/lemmy_guitar.gif"

func main() {

  // Open the file
  reader, err := os.Open(file)
  if err != nil {
    panic(err.Error())
  }
  defer reader.Close()

  // Store each file in memory so we get access to each frame of the animated gif
  imgFiles := separateAnimatedGif(reader)

  // Create our 
  lwf := svgr.NewSvgr(imgFiles, 20, "lemmy_guitar")

  /*  

  Example of a 3-channel single frame ouput

  for x:=0; x < len(imgFiles); x++ {
    lwf.SingleChannel("red", "#f03c3c", .6, 50, 0, false, x)
    lwf.SingleChannel("blue", "#3c9cf0", .6, 50, -8, false, x)
    lwf.SingleChannel("green", "#63f03c", .4, 50, 6, false, x)
  }

  */

  lwf.FunkyTriangles()

  lwf.Save(dest + lwf.GetName() + ".svg")
}

/*
* Utities
*/

// TODO
func videoToAnimatedGif(video *os.File) *os.File {
  return video
}

// Separate each image in an animated gif and resave in a unique folder
// Create a read of each file in the directory
// Return an array of blobs of each image and the directory
func separateAnimatedGif(animated *os.File) (imageFiles [][]byte) {

  // Generate a UUID and make a directory with corresponding name
  dir := fmt.Sprintf("./_%s", uuid.NewV4())
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

  // Clean up the temprorary directory once each image is stored in imageFiles blob
  if err := os.RemoveAll(dir); err != nil {
    panic(err.Error())
  }

  return
}