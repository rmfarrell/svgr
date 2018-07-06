package svgr

import (
	"fmt"
)

func (m *Mosaic) Dots(size float64) string {

	return m.render(func() string {

		return fmt.Sprintf(
			"<circle r=\"%.2fpx\" cy=\"%d\" cx=\"%d\" fill=\"%s\"/>",
			size*0.5*float64(shapeSize),
			m.current.X*shapeSize+5,
			m.current.Y*shapeSize+5,
			m.colorAtCurrent(),
		)
	})
}
