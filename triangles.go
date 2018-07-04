package svgr

import (
	"fmt"
	"math/rand"
	"time"
)

func (m *Mosaic) Triangles(rnd int) string {
	src := rand.New(rand.NewSource(time.Now().UnixNano()))

	return m.render(func() string {

		x := m.current.X * shapeSize
		y := m.current.Y * shapeSize
		triangle := [3]*point{}

		if m.current.X%2 == 0 {
			// Draw down-pointing triangle
			triangle = [3]*point{
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
			// Draw up-pointing triangle
			triangle = [3]*point{
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

		if rnd > 0 {
			for _, pt := range triangle {
				pt.randomize(*src, rnd)
			}
		}

		return fmt.Sprintf(
			"<polygon points=\"%d,%d %d,%d %d,%d\" fill=\"%s\"/>",
			triangle[0].x,
			triangle[0].y,
			triangle[1].x,
			triangle[1].y,
			triangle[2].x,
			triangle[2].y,
			m.colorAtCurrent(),
		)
	})
}
