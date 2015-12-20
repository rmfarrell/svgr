package write_svg

import (
  "fmt"
  "bufio"
  "os"
)

func init() {

}

func draw_rectangle(x,y uint8, color string) string {

  // <rect id="svg_2" height="112" width="84" y="17" x="1" stroke-width="5" fill="rgb(0,150,150)"/>
  // <rect id="svg_3" height="112" width="102" y="18" x="84" stroke-width="5" fill="#1500ff"/>

  return ""
}

func Write(pixel_data []uint8, width, height int) (svg string, error error) {
  svg = fmt.Sprintf("<svg width=\"%d\" height=\"%d\"  xmlns=\"http://www.w3.org/2000/svg\"><g>", width*10, height*10)

  fmt.Println(width, height)

  i := 0

  // Iterate over rows
  for row := 0; row < height; row++ {

    // Iterate over columns
    for col := 0; col < width; col++ {

      r:= pixel_data[i]
      g:= pixel_data[i+1]
      b:= pixel_data[i+2]

      i = i+3
      
      svg += fmt.Sprintf("<rect height=\"10\" width=\"10\" y=\"%d\" x=\"%d\" fill=\"rgb(%d,%d,%d)\"/>", row*10, col*10, r, g, b)    

    }
  }

  svg += "</g></svg>"

  write_file(svg)

  return
}

func write_file(contents string) {

  file, err := os.Create("mine.svg")
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