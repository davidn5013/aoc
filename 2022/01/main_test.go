package main

import "testing"

var (
	inpStr = []string{
		"1000",
		"2000",
		"3000",
		" ",
		"4000",
		" ",
		"5000",
		"6000",
		" ",
		"7000",
		"8000",
		"9000",
		" ",
		"10000",
	}
)

func TestElfMosteFood(t *testing.T) {
	ans := elfMosteFood(inpStr)
	if ans != 24000 {
		t.Errorf("elfMosteFood(input) = %d ; want 24000\n", ans)
	}
}

func TestCaloriePerElfSort(t *testing.T) {
	ans := caloriePerElfSort(inpStr)[0]
	if ans != 24000 {
		t.Errorf("caloriePerElfSort(input) = %d ; want 24000\n", ans)
	}
}
