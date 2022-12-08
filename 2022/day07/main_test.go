package main

import "testing"

func TestSol1(t *testing.T) {
	ans := sol1("input_test.txt")
	if ans != 95437 {
		t.Errorf("Sol1_1 =%d ;want 95437", ans)
	}
}

func BenchmarkSol1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sol1("input_test.txt")
	}
}

func TestSol2(t *testing.T) {
	ans := sol2("input_test.txt")
	if ans != 24933642 {
		t.Errorf("Sol1_1 =%d ;want 24933642", ans)
	}
}

func BenchmarkSol2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sol2("input_test.txt")
	}
}
