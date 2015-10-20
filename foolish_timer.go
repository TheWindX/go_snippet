package main

import (
	//"fmt"
	"./utils"
	"runtime"
	"time"
)

//定时器
func newTimer(t float64) chan bool {
	start := time.Now()
	r := make(chan bool)

	go func() {
		for {
			dis := time.Now().Sub(start).Seconds()
			if dis >= t {
				r <- true
				//runtime.Gosched() //相当于sleep(0)
				close(r)
				break
			}
		}
	}()

	return r
}

//定时器tick
func newTick(t float64) chan bool {
	start := time.Now()
	r := make(chan bool)

	go func() {
		for {
			newStart := time.Now()
			dis := newStart.Sub(start).Seconds()
			if dis >= t {
				r <- true
				start = newStart.Add(time.Duration(-float64(dis) - 1))
			}
		}
	}()

	return r
}

func main() {
	//runtime.GOMAXPROCS(0)
	utils.PrintNow()

	t := newTimer(2.0)
	_ = <-t

	t = newTick(0.5)
	for {
		_ = <-t
		utils.PrintNow()
	}

	utils.PrintNow()
}
