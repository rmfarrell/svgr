package svgr

import "fmt"

// Draw the image as a field of interlocking triangles
func (px *pixelArray) Triangles(frames ...int) (svg string, error error) {

  for _, frame := range normalizeFramesArray(frames, px.GetSize()) {

    writeGroup(px, frame, func(rgb []uint8, x,y int) string {

      x = x*10
      y = y*10
      g := [][]int{}

      if x/10 % 2 == 0 {
        
        // Draw down-pointing triangle
        g = [][]int {
          []int {x-4,y+0},
          []int {x+16,y+0},
          []int {x+6,y+10},
        }

      } else {

        // Draw up-pointing triangle
        g = [][]int {
          []int {x-4,y+10},
          []int {x+16,y+10},
          []int {x+6,y+0},
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