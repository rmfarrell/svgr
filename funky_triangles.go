package svgr

// Draw the image as a field of interlocking triangles but with random angles thrown in.
/*
func (px *pixelArray) FunkyTriangles(frames ...int) (svg string, error error) {

  for _, frame := range normalizeFramesArray(frames, px.GetSize()) {

    writeGroup(px, frame, func(rgb []uint8, x,y int) string {

      x = x*10
      y = y*10
      g := [][]int{}
      f := rand.Intn(Funkiness)

      if x/10 % 2 == 0 {

        // Draw down-pointing triangle
        g = [][]int {
          []int {x-f,y+f},
          []int {x+16,y},
          []int {x+5+f,y+8+f},
        }

      } else {

        // Draw up-pointing triangle
        g = [][]int {
          []int {x-4,y+10+f},
          []int {x+16,y+10+f},
          []int {x+6+f ,y+0},
        }
      }

      return fmt.Sprintf(
        "<polygon points=\"%d,%d %d,%d %d,%d\" fill=\"#%x\"/>",
        g[0][0],
        g[0][1],
        g[1][0],
        g[1][1],
        g[2][0],
        g[2][1],
        rgb,
      )
    })
  }
  return
}
*/
