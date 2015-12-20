package write_svg

import (
  "fmt"
  "bufio"
  "os"
)

type pixel_array struct {
  pixel_data []uint8
  w          int
  h          int
}

func NewPixelData(pixel_data []uint8, w,h int) pixel_array {
  return pixel_array {
    pixel_data: pixel_data,
    w: w,
    h: h,
  }
}

func draw_rectangle(x,y uint8, color string) string {

  // <rect id="svg_2" height="112" width="84" y="17" x="1" stroke-width="5" fill="rgb(0,150,150)"/>
  // <rect id="svg_3" height="112" width="102" y="18" x="84" stroke-width="5" fill="#1500ff"/>

  return ""
}

func (px pixel_array) Write(dest string) (svg string, error error) {

  svg = fmt.Sprintf("<svg width=\"%d\" height=\"%d\"  xmlns=\"http://www.w3.org/2000/svg\"><g>", px.w*10, px.h*10)

  i := 0

  // Iterate over rows
  for row := 0; row < px.h; row++ {

    // Iterate over columns
    for col := 0; col < px.w; col++ {

      r:= px.pixel_data[i]
      g:= px.pixel_data[i+1]
      b:= px.pixel_data[i+2]

      i = i+3
      
      svg += fmt.Sprintf("<rect height=\"10\" width=\"10\" y=\"%d\" x=\"%d\" fill=\"rgb(%d,%d,%d)\"/>", row*10, col*10, r, g, b)    

    }
  }

  svg += "</g></svg>"

  write_file(svg, dest)

  return
}

func write_file(contents string, dest string) {

  file, err := os.Create(dest)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  w:= bufio.NewWriter(file)
  w.WriteString(contents)
  w.Flush()

}

func gather_rgb_values(pixels []uint8) [][]int {

  rgb_array := [][]int{}

  for p := 0; p < len(pixels); p=p+3 {

    // pa := make([]int, 3)

    pa := []int{
      int(pixels[p]),
      int(pixels[p+1]),
      int(pixels[p+2]),
    }

    rgb_array = append(rgb_array, pa) 
  }

  return rgb_array
}

func gather_green_values() {

}

func gather_red_values() {

}

func gather_blue_values() {

}