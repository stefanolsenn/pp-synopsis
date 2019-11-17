package main

import "testing"

var (
	testData   []int64
	c          chan []int64
	goRoutines int
)

func init() {
	for i := 0; i < 1e6; i++ {
		testData = append(testData, int64(i))
	}
	goRoutines = 2
	c = make(chan []int64, goRoutines)
}

func BenchmarkFindEven(b *testing.B) {
	for n := 0; n < b.N; n++ {
		go FindEven(c, testData, goRoutines)
		<-c
	}
}

func BenchmarkFindEvenSeq(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FindEvenSeq(testData)
	}
}
