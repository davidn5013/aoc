package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type coord struct {
	p    []bool
	size int
}

func newCoord(size int) coord {
	var c coord
	c.p = make([]bool, size*size)
	c.size = size
	return c
}

func (c coord) pos(x, y int) int {
	return (y * c.size) + x
}

func (c coord) posbounce(pos int) bool {
	if pos < len(c.p) {
		return true
	}
	log.Fatalf("Out of posbounce x=%d\n", pos)
	return false
}

func (c coord) setpos(pos int) {
	if c.posbounce(pos) {
		c.p[pos] = true
	}
}

func (c coord) unsetpos(pos int) {
	if c.posbounce(pos) {
		c.p[pos] = false
	}
}

func (c coord) setposRange(xs, xe int) {
	for x := xs; x <= xe; x++ {
		c.setpos(x)
	}
}
func (c coord) unsetposRange(xs, xe int) {
	for x := xs; x <= xe; x++ {
		c.unsetpos(x)
	}
}

func (c coord) getpos(pos int) int {
	if c.posbounce(pos) && c.p[pos] {
		return 1
	}
	return 0
}

func (c coord) tooglepos(pos int) {
	if c.posbounce(pos) {
		switch c.getpos(pos) == 1 {
		case true:
			c.unsetpos(pos)
		case false:
			c.setpos(pos)
		default:
			log.Fatalf("Error in tooglepos on %d", pos)
		}
	}
}

func (c coord) posToggleRange(xs, xe int) {
	for x := xs; x <= xe; x++ {
		c.tooglepos(x)
	}
}

func (c coord) bounce(x, y int) bool {
	if c.pos(x, y) < len(c.p) {
		return true
	}
	log.Fatalf("Out of bounce x=%d,y=%d\n", x, y)
	return false
}

func (c coord) set(x, y int) {
	if c.bounce(x, y) {
		c.p[c.pos(x, y)] = true
	}
}

func (c coord) setRange(xs, ys, xe, ye int) {
	for x := xs; x <= xe; x++ {
		for y := ys; y <= ye; y++ {
			c.set(x, y)
		}
	}
}

func (c coord) unset(x, y int) {
	if c.bounce(x, y) {
		c.p[c.pos(x, y)] = false
	}
}

func (c coord) unsetRange(xs, ys, xe, ye int) {
	for x := xs; x <= xe; x++ {
		for y := ys; y <= ye; y++ {
			c.unset(x, y)
		}
	}
}

func (c coord) get(x, y int) int {
	if c.bounce(x, y) && c.p[c.pos(x, y)] {
		return 1
	}
	return 0
}

func (c coord) toggle(x, y int) {
	c.p[c.pos(x, y)] = !c.p[c.pos(x, y)]
}

func (c coord) toggleRange(xs, ys, xe, ye int) {
	for x := xs; x <= xe; x++ {
		for y := ys; y <= ye; y++ {
			c.toggle(x, y)
		}
	}
}

func (c coord) cnt() int {
	cnt := 0
	for i := 0; i < c.size*c.size; i++ {
		if c.p[i] == true {
			cnt++
		}
	}
	return cnt
}

func (c coord) print() {
	for i := 0; i < c.size-1; i++ {
		for j := 0; j < c.size-1; j++ {
			fmt.Printf("%x", c.get(i, j))
		}
		fmt.Println()
	}
}

func (c coord) createImage() {
	width := c.size
	height := c.size

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	red := color.RGBA{200, 0, 0, 0xff}
	green := color.RGBA{0, 200, 0, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch c.get(x, y) {
			case 1:
				img.Set(x, y, red)
			case 0:
				img.Set(x, y, green)
			default:
				// Use zero value.
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
