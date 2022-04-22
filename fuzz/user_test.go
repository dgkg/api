package fuzz

import "testing"

func TestFuzzUser(t *testing.T) {
	fuzzUser_old(100000)
}

func BenchmarkFuzzUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuzzUser(1000)
	}
}

func BenchmarkFuzzUserOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fuzzUser_old(1000)
	}
}
