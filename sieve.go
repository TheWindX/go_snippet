package main

import (
	"fmt"
)

func intStream(c chan int) {
	for i := 2; ; i++ {
		c <- i
	}
}

func filteStream(os chan int, is chan int, filteFunc func(int) bool) {
	for {
		i := <-is
		if filteFunc(i) {
			os <- i
		}
	}
}

func sieveStream(os chan int, is chan int) {
	head := <-is
	os <- head
	nextSieve := make(chan int)
	go sieveStream(os, nextSieve)
	go filteStream(nextSieve, is, func(c int) bool {
		if c%head == 0 {
			return false
		}
		return true
	})
}

func main() {
	from2 := make(chan int)
	sieves := make(chan int)
	go intStream(from2)
	go sieveStream(sieves, from2)

	for {
		i := <-sieves
		fmt.Println(i)
		if i > 1000 {
			return
		}
	}
}
