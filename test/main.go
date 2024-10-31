package main

import (
	"fmt"

	"github.com/davidn5013/aoc/data-structures/slice"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b := []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	z1, err := slice.ZipArr(a, b)
	if err != nil {
		panic(err)
	}
	for _, v := range z1 {
		fmt.Printf("%d%d", v[0], v[1])
	}
	fmt.Println()

	z2 := slice.NewZipTuple()
	err = z2.ZipTuple(a, b)
	if err != nil {
		panic(err)
	}
	for _, v := range z2 {
		fmt.Printf("%d%d", v.A, v.B)
	}
	fmt.Println()

	z3 := make([][2]int, len(a))
	err = slice.ZipGen(a, b, &z3)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", z3)
	fmt.Println()

	d := []float64{0.1, 0.2, 0.3}
	e := []float64{0.1, 0.2, 0.3}
	z4 := [][2]float64{}
	err = slice.ZipGen(d, e, &z4)
	if err != nil {
		panic(err)
	}
	for _, v := range z4 {
		fmt.Printf("%v %v", v[0], v[1])
	}
	fmt.Println()
}
