package main

import (
	"fmt"
	"time"
)

func TimeSince(start time.Time) int {
	var dur = time.Now().Sub(start)
	return int(dur.Seconds() * 1000)
}

func NewTimer(endFlag chan int, t int) {
	start := time.Now()
	timeAdd := 0
	fmt.Printf("计时开始...0\n")
	for {
		if TimeSince(start) > timeAdd {
			timeAdd = TimeSince(start)
			if TimeSince(start)%500 == 0 {
				fmt.Printf("计时中...%d\n", TimeSince(start))
			}
		}
		if TimeSince(start) > t {
			endFlag <- TimeSince(start)
			break
		}
		//yield return null;
	}
}

func afterTimer(endFlag chan int) {
	val := <-endFlag
	fmt.Printf("时间到达:%d\n", val)
}

func main() {
	var endFlag = make(chan int)

	go NewTimer(endFlag, 3000)
	go afterTimer(endFlag)

	for {

	}
}
