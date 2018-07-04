package svgr

// Draw the image as a field of interlocking hexagons
/*
func (px *pixelArray) Hexagons(frames ...int) (svg string, error error) {

  for _, frame := range normalizeFramesArray(frames, px.GetSize()) {

    writeGroup(px, frame, func(rgb []uint8, x,y int) string {

      z := 0

      if x % 2 == 0 {
        z = 5
      } else {
        z = 0
      }

      x = x*10
      y = y*10

      g := [][]int {
        []int {x+0,y+5+z},
        []int {x+5,y+0+z},
        []int {x+10,y+0+z},
        []int {x+15,y+5+z},
        []int {x+10,y+10+z},
        []int {x+5,y+10+z},
      }

      return fmt.Sprintf(
        "<polygon points=\"%d,%d %d,%d %d,%d %d,%d %d,%d %d,%d\" fill=\"#%x\"/>",
        g[0][0],
        g[0][1],
        g[1][0],
        g[1][1],
        g[2][0],
        g[2][1],
        g[3][0],
        g[3][1],
        g[4][0],
        g[4][1],
        g[5][0],
        g[5][1],
        rgb,
      )
    })
  }
  return
}
*/
