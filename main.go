package main

import (
	"fmt"
	"time"
)

var _data []int64

const (
	_goRoutines = 20  // Numbers of goroutines that should run at once
	_dataSize   = 1e6 // Size of the _data array
)

// Sub-method for finding even numbers
// Takes in a chan that acts as a Semaphore
// An array to find even numbers in
// and a integer to determine which sub-method this is.
func even(sem chan []int64, arr []int64, i int) {
	fmt.Println("Finding even numbers: ", i)
	time.Sleep(time.Second * 1)
	tmpArr := make([]int64, int64(len(arr)))
	index := 0
	for _, d := range arr {
		if d%2 == 0 && d != 0 {
			tmpArr[index] = d
			index++
		}
	}
	if index < len(tmpArr) {
		tmpArr = tmpArr[:index]
	}
	fmt.Println("Sending result into semaphore channel..: ", i)
	sem <- tmpArr
}

func FindEven(c chan []int64, data []int64, goRoutines int) {
	ngr := goRoutines              // Number of goroutines
	sem := make(chan []int64, ngr) // semaphore

	chunkSize := len(data) / ngr
	if chunkSize == 0 {
		chunkSize = 1
	}
	start := 0
	for i := 0; i < ngr; i++ {
		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}
		go even(sem, data[start:end], i)
		start += chunkSize
	}
	var result []int64
	for i := 0; i < ngr; i++ {
		result = append(result, <-sem...)
	}
	fmt.Println("Done...")
	c <- result
}

func FindEvenSeq(arr []int64) []int64 {
	tmpArr := make([]int64, int64(len(arr)))
	index := 0
	for _, d := range arr {
		if d%2 == 0 && d != 0 {
			tmpArr[index] = d
			index++
		}
	}
	if index < len(tmpArr) {
		tmpArr = tmpArr[:index]
	}
	return tmpArr
}

func main() {
	c := make(chan []int64)
	_data = generateData(_dataSize)
	FindEvenSeq(_data)
	go FindEven(c, _data, _goRoutines)
	<-c
	printResults()
}

func printResults() {
	fmt.Println("-------------------------")
	fmt.Println("Chunksize: ", (len(_data) / _goRoutines))
	fmt.Println("Numbers calculated: ", (len(_data)/_goRoutines)*_goRoutines)
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s time:  %v\n", what, time.Since(start))
	}
}

func generateData(max int64) []int64 {
	arr := make([]int64, max)
	for i := int64(0); i < int64(len(arr)); i++ {
		arr[i] = i
	}
	return arr
}
