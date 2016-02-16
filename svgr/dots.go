package svgr

import "fmt"

// Draw the image as a field of circles
func (px *pixelArray) Dots(frames ...int) {

  frames = normalizeFramesArray(frames, px.GetSize())

  for _, frame := range frames {
    writeGroup(px, frame, func(rgb []uint8, x,y int) string {
      return fmt.Sprintf(
        "<circle r=\"5px\" cy=\"%d\" cx=\"%d\" fill=\"#%x\"/>",
        y*10,
        x*10, 
        rgb,
      )
    })
  }
  return
}