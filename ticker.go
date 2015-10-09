package main

import (
	"fmt"
	"time"
)

func main() {
	b := time.Tick(1e9)
	e := time.Tick(115e9)
	for {
		select {
		case <-b:
			fmt.Println("1e9")
		case <-e:
			fmt.Println("end")
			return
		default:
			fmt.Println("...")
			time.Sleep(5e8)
		}
	}
}
