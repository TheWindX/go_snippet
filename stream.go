// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	//"math/rand"
	//"runtime"
	//"time"
)

//整数流
func int_stream() chan int {
	s := make(chan int)
	go func() {
		for i := 0; ; i++ {
			s <- i
		}
	}()
	return s
}

func main() {
	ch1 := int_stream()
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
}
