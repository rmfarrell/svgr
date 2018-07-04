package svgr_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rmfarrell/svgr"
)

var input svgr.Input = svgr.Input{
	Path: "./testdata/doge.png",
	Resolution: &svgr.Resolution{
		Width:  10,
		Height: 10,
	},
	ID: "",
}

func TestNewMosaic(t *testing.T) {
	t.Skip()
	fixtures := []struct {
		input *svgr.Input
		err   error
	}{
		// okay
		{
			&svgr.Input{
				Path: "./testdata/doge.png",
				Resolution: &svgr.Resolution{
					Width:  10,
					Height: 10,
				},
				ID: "",
			},
			nil,
		},
	}

	for _, fix := range fixtures {
		m, err := svgr.NewMosaic(fix.input)

		if fix.err != nil && err.Error() != fix.err.Error() {
			t.Errorf("Expected %v \n Received %v", fix.err, err)
		}
		if err != nil {
			continue
		}
		// check that type if Mosaic
		if (reflect.TypeOf(m) != reflect.TypeOf(&svgr.Mosaic{})) {
			t.Errorf("Expected *svgr.Mosaic \n Received %v", m)
		}
		// check that it's not empty
		if (&svgr.Mosaic{}) == m {
			t.Error("Mosaic should be non-empty")
		}
	}
}

func TestTriangles(t *testing.T) {
	m, err := svgr.NewMosaic(&input)
	if err != nil {
		t.Error(err)
	}
	svg := m.Triangles(3)
	if svg == "" {
		t.Error("svg should not be blank")
	}
	// TODO: test number of polygons
}

func TestSquares(t *testing.T) {
	m, err := svgr.NewMosaic(&input)
	if err != nil {
		t.Error(err)
	}
	svg := m.Squares(0)
	if svg == "" {
		t.Error("svg should not be blank")
	}
}
func TestDots(t *testing.T) {
	m, err := svgr.NewMosaic(&input)
	if err != nil {
		t.Error(err)
	}
	svg := m.Dots(0.8)
	if svg == "" {
		t.Error("svg should not be blank")
	}
}
func TestHexagaons(t *testing.T) {
	m, err := svgr.NewMosaic(&input)
	if err != nil {
		t.Error(err)
	}
	svg := m.Hexagons()
	fmt.Println(svg)
	if svg == "" {
		t.Error("svg should not be blank")
	}
}
