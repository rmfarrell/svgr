package svgr

import "fmt"

/*
* Draw a <g> from the image as a field of Squares
* @param frames {...int} a comma-separated list of frames from the array to render
* a blank value loops through every image in the pixelArray.pixelData array.
*/
func (px *pixelArray) Blocks(frames ...int) {
  // TODO: error/recover from pushing frames that exceed size of pixelArray

  frames = normalizeFramesArray(frames, px.GetSize())

  for _, frame := range frames {
    writeGroup(px, frame, func(rgb []uint8, x,y int) string {
      return fmt.Sprintf(
        "<rect height=\"10px\" width=\"10px\" y=\"%d\" x=\"%d\" fill=\"#%x\"/>",
        y*10, 
        x*10, 
        rgb,
      )
    })
  }
  return
}