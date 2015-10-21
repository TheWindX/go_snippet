// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	//"math/rand"
	"runtime"
	"time"
)

func f1() {
	for {
		fmt.Println("loop in f1  ")
		runtime.Gosched()
		time.Sleep(1)
	}
}

func f2() {
	for {
		fmt.Println("loop in f2  ")
		runtime.Gosched()
		time.Sleep(1)
	}
}

func main() {
	//runtime.GOMAXPROCS(1)
	go f1()
	go f2()
	for i := 0; i < 10; i++ {
		fmt.Println("loop in main  ")
		runtime.Gosched()
		time.Sleep(1)
	}
}
