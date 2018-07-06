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
				x: x + 3,
				y: y + 0 + offset,
			},
			&point{
				x: x + 10,
				y: y + 0 + offset,
			},
			&point{
				x: x + 13,
				y: y + 5 + offset,
			},
			&point{
				x: x + 10,
				y: y + 10 + offset,
			},
			&point{
				x: x + 3,
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
