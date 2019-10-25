package main

import (
	"fmt"
	"runtime"
	"time"
)

func DoPart(sem chan int, num int) {
	fmt.Printf("%d. Computing...\n", num)
	time.Sleep(2 * time.Second)
	sem <- num
}

func DoTask(c chan int, result int) {
	NCPU := runtime.NumCPU() // Antal kerner i cpu
	sem := make(chan int, NCPU)

	for i := 0; i < NCPU; i++ {
		go DoPart(sem, i)
	}

	for i := 0; i < NCPU; i++ {
		x := <-sem
		result += x
		fmt.Printf("Tilføjede %d til resultat\n", x)
	}
	c <- result
}

func main() {
	c := make(chan int)
	result := 0
	go DoTask(c, result)
	result = <-c
	go DoTask(c, result)
	result = <-c
	fmt.Printf("result: %d", result)
}
