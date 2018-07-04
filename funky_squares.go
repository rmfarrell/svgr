package svgr

// Draw the image as a field of rectangles with random angles thrown in.
/*
func (px *pixelArray) FunkySquares(frames ...int) (svg string, error error) {

  for _, frame := range normalizeFramesArray(frames, px.GetSize()) {

    writeGroup(px, frame, func(rgb []uint8, x,y int) string {

      f := rand.Intn(Funkiness)

      x = x*10
      y = y*10

      g := [][]int {
        []int {x+f/2,y-f},
        []int {x+10,y-f},
        []int {x+10-f,y+10},
        []int {x-f,y+10},
      }

      return fmt.Sprintf(
        "<polygon points=\"%d,%d %d,%d %d,%d %d,%d\" fill=\"#%x\"/>",
        g[0][0],
        g[0][1],
        g[1][0],
        g[1][1],
        g[2][0],
        g[2][1],
        g[3][0],
        g[2][1],
        rgb,
      )
    })
  }
  return
}
*/
