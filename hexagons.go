package svgr

import (
	"fmt"
)

type hexagon [6]*point

func (m *Mosaic) Hexagons() string {

	return m.render(func() string {

		var hex hexagon
		offset := 0
		x := m.current.X * shapeSize
		y := m.current.Y * shapeSize

		if m.current.X%2 == 0 {
			offset = 5
		} else {
			offset = 0
		}

		hex = hexagon{
			&point{
				x: x,
				y: y + 5 + offset,
			},
			&point{
				x: x + 5,
				y: y + 0 + offset,
			},
			&point{
				x: x + 10,
				y: y + 0 + offset,
			},
			&point{
				x: x + 15,
				y: y + 5 + offset,
			},
			&point{
				x: x + 10,
				y: y + 10 + offset,
			},
			&point{
				x: x + 5,
				y: y + 10 + offset,
			},
		}

		return fmt.Sprintf(
			"<polygon points=\"%d,%d %d,%d %d,%d %d,%d %d,%d %d,%d\" fill=\"%s\"/>",
			hex[0].x,
			hex[0].y,
			hex[1].x,
			hex[1].y,
			hex[2].x,
			hex[2].y,
			hex[3].x,
			hex[3].y,
			hex[4].x,
			hex[4].y,
			hex[5].x,
			hex[5].y,
			m.colorAtCurrent(),
		)
	})
}

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
