package svgr

import (
	"fmt"
	"math/rand"
	"time"
)

type square [4]*point

// Render squares
// NOTE: <polygon> is more slightly terser than <rect> ¯\_(ツ)_/¯
func (m *Mosaic) Squares(rnd int) string {
	src := rand.New(rand.NewSource(time.Now().UnixNano()))

	return m.render(func() string {

		x := m.current.X * shapeSize
		y := m.current.Y * shapeSize

		sq := square{
			&point{x, y},
			&point{x + shapeSize, y},
			&point{x + shapeSize, y + shapeSize},
			&point{x, y + shapeSize},
		}

		if rnd > 0 {
			for _, pt := range sq {
				pt.randomize(*src, rnd)
			}
		}

		return fmt.Sprintf(
			"<polygon points=\"%d,%d %d,%d %d,%d %d,%d\" fill=\"%s\"/>",
			sq[0].x,
			sq[0].y,
			sq[1].x,
			sq[1].y,
			sq[2].x,
			sq[2].y,
			sq[3].x,
			sq[3].y,
			m.colorAtCurrent(),
		)
	})
}
