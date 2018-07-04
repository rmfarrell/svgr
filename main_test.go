package svgr_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rmfarrell/svgr"
)

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

func TestMain(t *testing.T) {
	m, err := svgr.NewMosaic(&svgr.Input{
		Path: "./testdata/doge.png",
		Resolution: &svgr.Resolution{
			Width:  10,
			Height: 10,
		},
		ID: "",
	})
	if err != nil {
		t.Error(err)
	}
	svg := m.Triangles()
	fmt.Println(svg)
	if svg == "" {
		t.Error("svg should not be blank")
	}
}
