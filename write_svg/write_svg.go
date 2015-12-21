package write_svg

import (
  "fmt"
  // "reflect"
  "bufio"
  _ "image"
  "os"
  "github.com/gographics/imagick/imagick"
  "math/rand"
)

const (
  MaxSize             int       = 80
  AdaptiveSharpenVal  float64   = 16
  Funkiness           int       = 6
)

type pixel_array struct {
  pixel_data []uint8
  w          int
  h          int
}

// Constructor

func NewSvgr(img *os.File) pixel_array {

  imagick.Initialize()
  defer imagick.Terminate()

  wand := imagick.NewMagickWand()

  wand.ReadImageFile(img)

  w,h := shrink_image(wand)

  pixel_data, err := wand.ExportImagePixels(0,0,w,h,"RGB", imagick.PIXEL_CHAR)
  if err != nil {
    panic(err.Error())
  }

  return pixel_array {
    pixel_data: pixel_data.([]uint8),
    w: int(w),
    h: int(h),
  }
}

// Class Methods

func (px pixel_array) Pixels(name, dest string) (svg string, error error) {

  svg, error = write(px, name, dest, func(rgb []uint8, x,y int) string {
    return fmt.Sprintf("<rect height=\"10\" width=\"10\" y=\"%d\" x=\"%d\" fill=\"#%x\"/>", y*10, x*10, rgb)
  })
  return
}

func (px pixel_array) Dots(name, dest string) (svg string, error error) {

  svg, error = write(px, name, dest, func(rgb []uint8, x,y int) string {
    return fmt.Sprintf("<circle r=\"5\" cy=\"%d\" cx=\"%d\" fill=\"#%x\"/>", y*10, x*10, rgb)
  })
  return
}

func (px pixel_array) PolyGonSquare(name, dest string) (svg string, error error) {

  svg, error = write(px, name, dest, func(rgb []uint8, x,y int) string {

    x = x*10
    y = y*10

    g := [][]int {
      []int {x+0,y+0},
      []int {x+10,y+0},
      []int {x+10,y+10},
      []int {x+0,y+10},
    }

    return fmt.Sprintf(
      "<polygon points=\"%d,%d %d,%d %d,%d %d,%d\" fill=\"#%x\"/>",
      g[0][0],
      g[0][1],
      g[1][0],
      g[1][1],
      g[2][0],
      g[2][1],
      g[3][0],
      g[2][1],
      rgb,
    )
  })
  return
}

func (px pixel_array) FunkyTriangles(name, dest string) (svg string, error error) {

  svg, error = write(px, name, dest, func(rgb []uint8, x,y int) string {

    x = x*10
    y = y*10
    g := [][]int{}
    f := rand.Intn(Funkiness)

    if x/10 % 2 == 0 {
      
      // Draw down-pointing triangle
      g = [][]int {
        []int {x-f,y-f/2},
        []int {x+16,y},
        []int {x+f,y+10},
      }

    } else {

      // Draw up-pointing triangle
      g = [][]int {
        []int {x-4,y+10},
        []int {x+16,y+10},
        []int {x+6+f ,y+0},
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
  return
}

func (px pixel_array) Triangles(name, dest string) (svg string, error error) {

  svg, error = write(px, name, dest, func(rgb []uint8, x,y int) string {

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
  return
}

func (px pixel_array) FunkySquares(name, dest string) (svg string, error error) {

  svg, error = write(px, name, dest, func(rgb []uint8, x,y int) string {

    f := rand.Intn(Funkiness)

    x = x*10
    y = y*10

    g := [][]int {
      []int {x+f/2,y-f},
      []int {x+10,y-f},
      []int {x+10-f,y+10},
      []int {x-f,y+10},
    }

    return fmt.Sprintf(
      "<polygon points=\"%d,%d %d,%d %d,%d %d,%d\" fill=\"#%x\"/>",
      g[0][0],
      g[0][1],
      g[1][0],
      g[1][1],
      g[2][0],
      g[2][1],
      g[3][0],
      g[2][1],
      rgb,
    )
  })
  return
}


func (px pixel_array) Hexagons(name, dest string) (svg string, error error) {

  svg, error = write(px, name, dest, func(rgb []uint8, x,y int) string {

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
  return
}

func write(pxa pixel_array, name, dest string, pixel_method func([]uint8, int, int) string) (svg string, error error) {

  svg = fmt.Sprintf("<svg width=\"%d\" height=\"%d\"  xmlns=\"http://www.w3.org/2000/svg\"><g>", pxa.w*10, pxa.h*10)

  i := 0

  // Iterate over rows
  for row := 0; row < pxa.h; row++ {

    // Iterate over columns
    for col := 0; col < pxa.w; col++ {

      svg += pixel_method(
        []uint8{
          pxa.pixel_data[i], 
          pxa.pixel_data[i+1], 
          pxa.pixel_data[i+2],
        },
        col,
        row,
      )

      i = i+3
    }
  }

  svg += "</g></svg>"

  write_file(svg, name, dest)

  return
}

func shrink_image(wand *imagick.MagickWand) (w,h uint) {

  w,h = get_dimensions(wand)

  shrinkBy := 1

  if w >= h {
    shrinkBy = int(w)/MaxSize
  } else {
    shrinkBy = int(h)/MaxSize
  }

  wand.AdaptiveResizeImage(
    uint(int(w)/shrinkBy), 
    uint(int(h)/shrinkBy),
  )

  wand.AdaptiveSharpenImage(0,AdaptiveSharpenVal)

  w,h = get_dimensions(wand)

  return
}

func get_dimensions(wand *imagick.MagickWand) (w,h uint) {
  h = wand.GetImageHeight()
  w = wand.GetImageWidth()
  return
}

func write_file(contents, name, dest string) {

  file, err := os.Create(dest + "/" + name + ".svg")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  w:= bufio.NewWriter(file)
  w.WriteString(contents)
  w.Flush()
}

// func gather_rgb_values(pixels []uint8) [][]int {

//   rgb_array := [][]int{}

//   for p := 0; p < len(pixels); p=p+3 {\

//     pa := []int{
//       int(pixels[p]),
//       int(pixels[p+1]),
//       int(pixels[p+2]),
//     }

//     rgb_array = append(rgb_array, pa) 
//   }

//   return rgb_array
// }

// func gather_green_values() {

// }

// func gather_red_values() {

// }

// func gather_blue_values() {

// }