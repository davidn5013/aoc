package main

import (
	"fmt"
	"testing"
)

func ExampleMostSol1() {
	fmt.Println(mostSol1([]string{
		"A Y",
		"B X",
		"C Z",
	}))
	// Output: 15
}

func TestMostSol1(t *testing.T) {
	ans := mostSol1([]string{
		"A Y",
		"B X",
		"C Z",
	})
	if ans != 15 {
		t.Errorf("Sol1 =%d ;want 15", ans)
	}
}
