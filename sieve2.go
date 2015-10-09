package main

import (
	"fmt"
)

func intStream(from int) chan int {
	s := make(chan int)
	go func() {
		for i := from; ; i++ {
			s <- i
		}
	}()
	return s
}

func filteStream(is chan int, pred func(int) bool) chan int {
	os := make(chan int)
	go func() {
		for {
			i := <-is
			if pred(i) {
				os <- i
			}
		}
	}()
	return os
}

func sieveStream(toSieve int, is chan int) chan int {
	os := make(chan int)
	go func() {
		os <- toSieve
		//fmt.Println("toSieve:", toSieve)
		fs := filteStream(is,
			func(n int) bool {
				return n%toSieve != 0
			})
		next := sieveStream(<-fs, fs)
		for {
			os <- <-next
		}

	}()
	return os
}

func main() {
	sieves := sieveStream(2, intStream(2))
	for {
		i := <-sieves
		if i > 100 {
			return
		}
		fmt.Println(i)
	}
}
