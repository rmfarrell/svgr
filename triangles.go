package svgr

import (
	"fmt"
)

func (m *Mosaic) Triangles() string {

	return m.render(func() string {

		x := m.current.X * shapeSize
		y := m.current.Y * shapeSize
		poly := [3]*point{}

		if m.current.X%2 == 0 {
			// Draw down-pointing triangle
			poly = [3]*point{
				&point{
					x - 4,
					y,
				},
				&point{
					x + 16,
					y,
				},
				&point{
					x + 6,
					y + 10,
				},
			}
		} else {
			poly = [3]*point{
				&point{
					x - 4,
					y + 10,
				},
				&point{
					x + 16,
					y + 10,
				},
				&point{
					x + 6,
					y,
				},
			}
		}

		return fmt.Sprintf(
			"<polygon points=\"%d,%d %d,%d %d,%d\" fill=\"%s\"/>",
			poly[0].x,
			poly[0].y,
			poly[1].x,
			poly[1].y,
			poly[2].x,
			poly[2].y,
			m.colorAtCurrent(),
		)
	})
}
