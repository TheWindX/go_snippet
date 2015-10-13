package main

import (
	"fmt"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(is <-chan int) <-chan int {
	os := make(chan int)
	go func() {
		for i := range is {
			os <- i * i
		}
		close(os)
	}()
	return os
}

func main() {
	for i := range sq(gen(3, 4, 5, 6)) {
		fmt.Println(i)
	}
}
