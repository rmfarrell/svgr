package svgr

import (
	"fmt"
	"math/rand"
	"time"
)

type triangle [3]*point

func (m *Mosaic) Triangles(rnd int) string {
	src := rand.New(rand.NewSource(time.Now().UnixNano()))

	return m.render(func() string {

		var tri triangle
		x := m.current.X * shapeSize
		y := m.current.Y * shapeSize

		if m.current.X%2 == 0 {
			// Draw down-pointing triangle
			tri = [3]*point{
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
			tri = [3]*point{
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
			for _, pt := range tri {
				pt.randomize(*src, rnd)
			}
		}

		return fmt.Sprintf(
			"<polygon points=\"%d,%d %d,%d %d,%d\" fill=\"%s\"/>",
			tri[0].x,
			tri[0].y,
			tri[1].x,
			tri[1].y,
			tri[2].x,
			tri[2].y,
			m.colorAtCurrent(),
		)
	})
}
