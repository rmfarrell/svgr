package svgr

/*
* Draw the image as a dot matrix representing a single color channel (ie r,g, or b)
* Combine multiple channels in the by calling this in for loop with the iterator pointing at the same frame
* @param channelName {string} acceptable values "red", "green", "blue"
* @param color {string} hex value of the output color (eg "#63f03c")
* @param opacity {float64} opacity of the output layer. Useful for drawing a multiple channels
* @param scale {uint8} increase or decrease the size of the dot.
* @param offset {int} shift position of the dot matrix to replicate a screen print effect
* @param negative {boolean} set to true to get the negative values from the input. A
*   false: represents the positive value of the r,g, or b channel from the input.
*   true:  represents the negative value.
* @param frames {...int} selectively apply this treatment to a comma separated list of indeces of images. See above.
 */
/*
func (px *pixelArray) SingleChannel(channelName, color string, opacity float64, scale uint8, offset int, negative bool, frames ...int) {

  channel      := 0
  colorOffset := uint8(0)

  switch channelName {
  case "red" :
    channel = 0
  case "green" :
    channel = 1
  case "blue" :
    channel = 2
  }

  if negative {
    colorOffset = 0
  } else {
    colorOffset = 255
  }

  for _, frame := range normalizeFramesArray(frames, px.GetSize()) {

    writeGroup(px, frame, func(rgb []uint8, x,y int) string {
      return fmt.Sprintf(
        "<circle r=\"%d\" cy=\"%d\" cx=\"%d\" opacity=\"%f\" fill=\"%s\"/>",
        (colorOffset-rgb[channel])/scale,
        y*10+offset,
        x*10+offset,
        opacity,
        color,
      )
    })
  }
  return
}
*/
