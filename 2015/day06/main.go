package main

import (
	_ "embed"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/davidn5013/aoc/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}
func createImageInt(grid [][]int) {
	width := len(grid)
	height := len(grid[0])

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	green := color.RGBA{0, 200, 0, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case grid[x][y] > 1:
				img.Set(x, y, color.RGBA{200, 0, 0, uint8(grid[x][y] * 20)})
			case grid[x][y] == 0:
				img.Set(x, y, green)
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func createImage(grid [][]bool) {
	width := len(grid)
	height := len(grid[0])

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	red := color.RGBA{200, 0, 0, 0xff}
	green := color.RGBA{0, 200, 0, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch grid[x][y] {
			case true:
				img.Set(x, y, red)
			case false:
				img.Set(x, y, green)
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
func _part1(input string) int {
	// 1000x1000 grid
	grid := make([][]bool, 1000)
	for i := range grid {
		grid[i] = make([]bool, 1000)
	}

	for _, line := range strings.Split(input, "\n") {
		switch {
		case strings.HasPrefix(line, "toggle"):
			var row1, col1, row2, col2 int
			fmt.Sscanf(line, "toggle %d,%d through %d,%d", &row1, &col1, &row2, &col2)
			for i := row1; i <= row2; i++ {
				for j := col1; j <= col2; j++ {
					grid[i][j] = !grid[i][j]
				}
			}
		case strings.HasPrefix(line, "turn on"):
			var row1, col1, row2, col2 int
			fmt.Sscanf(line, "turn on %d,%d through %d,%d", &row1, &col1, &row2, &col2)
			for i := row1; i <= row2; i++ {
				for j := col1; j <= col2; j++ {
					grid[i][j] = true
				}
			}
		case strings.HasPrefix(line, "turn off"):
			var row1, col1, row2, col2 int
			fmt.Sscanf(line, "turn off %d,%d through %d,%d", &row1, &col1, &row2, &col2)
			for i := row1; i <= row2; i++ {
				for j := col1; j <= col2; j++ {
					grid[i][j] = false
				}
			}
		default:
			panic("unhandled instruction")
		}
	}
	var count int
	for _, row := range grid {
		for _, b := range row {
			if b {
				count++
			}
		}
	}
	createImage(grid)
	return count
}

func part2(input string) int {
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, line := range strings.Split(input, "\n") {
		switch {
		case strings.HasPrefix(line, "toggle"):
			var row1, col1, row2, col2 int
			fmt.Sscanf(line, "toggle %d,%d through %d,%d", &row1, &col1, &row2, &col2)
			for i := row1; i <= row2; i++ {
				for j := col1; j <= col2; j++ {
					grid[i][j] += 2
				}
			}
		case strings.HasPrefix(line, "turn on"):
			var row1, col1, row2, col2 int
			fmt.Sscanf(line, "turn on %d,%d through %d,%d", &row1, &col1, &row2, &col2)
			for i := row1; i <= row2; i++ {
				for j := col1; j <= col2; j++ {
					grid[i][j]++
				}
			}
		case strings.HasPrefix(line, "turn off"):
			var row1, col1, row2, col2 int
			fmt.Sscanf(line, "turn off %d,%d through %d,%d", &row1, &col1, &row2, &col2)
			for i := row1; i <= row2; i++ {
				for j := col1; j <= col2; j++ {
					if grid[i][j] > 0 {
						grid[i][j]--
					}
				}
			}
		default:
			panic("unhandled instruction")
		}
	}
	var count int
	for _, row := range grid {
		for _, b := range row {
			count += b
		}
	}
	createImageInt(grid)
	return count
}

func part1(input string) int {
	c := newCoord(1000)
	/*569999
	  command to parse
	  turn on 489,959 through 759,964
	  turn off 820,516 through 871,914
	  toggle 756,965 through 812,992
	*/
	for _, line := range strings.Split(input, "\n") {
		var row1, col1, row2, col2 int
		switch {
		case line[:7] == "turn on":
			fmt.Sscanf(line, "turn on %d,%d through %d,%d", &row1, &col1, &row2, &col2)
			c.setRange(row1, col1, row2, col2)
		case line[:8] == "turn off":
			fmt.Sscanf(line, "turn off %d,%d through %d,%d", &row1, &col1, &row2, &col2)
			c.unsetRange(row1, col1, row2, col2)
		case line[:6] == "toggle":
			fmt.Sscanf(line, "toggle %d,%d through %d,%d", &row1, &col1, &row2, &col2)
			c.toggleRange(row1, col1, row2, col2)
		default:
			log.Fatalf("Wrong input format %s", line)
		}
	}
	c.createImage()
	return c.cnt()
}
