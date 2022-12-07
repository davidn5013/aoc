package main

import "testing"

func TestSol1(t *testing.T) {
	ans := sol1("input_test.txt", 3)
	if ans != "CMZ" {
		t.Errorf("Sol1_1 =%s ;want MCD", ans)
	}
}

func BenchmarkSol1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sol1("input_test.txt", 3)
	}
}

func TestSol2(t *testing.T) {
	ans := sol2("input_test.txt", 3)
	if ans != "MCD" {
		t.Errorf("Sol1_1 =%s ;want MCD", ans)
	}
}

func BenchmarkSol2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sol2("input_test.txt", 3)
	}
}
