package main

import (
	"fmt"
	"time"
)

func mySleep(secs float64) <-chan time.Time {
	ret := make(chan time.Time)
	go func() {
		time.Sleep(time.Duration(secs * float64(time.Second)))
		ret <- time.Now()
	}()
	return ret
}

type iTimeManager interface {
	setInterval(time float32, handle func())
	setTimeOut(time float32, handle func())
	run()
}

type timeManager struct {
	chans []interface{}
	funcs []func()
}

func (this *timeManager) setInterval(time1 float64, handle func()) {
	t := time.Tick(time.Duration(time1 * float64(time.Second)))
	this.chans = append(this.chans, t)
	this.funcs = append(this.funcs, handle)
}

func (this *timeManager) setTimeOut(time1 float64, handle func()) {
	t := mySleep(time1)
	this.chans = append(this.chans, t)
	this.funcs = append(this.funcs, handle)
}

func (this *timeManager) run() {
	for {
		for i := 0; i < len(this.chans); i++ {
			select {
			case <-(this.chans[i].(<-chan time.Time)):
				this.funcs[i]()
			default:
			}
			//safe type cast
			/*
				_, ok := this.chans[i].(<-chan time.Time)

				if ok {
					select {
					case <-(this.chans[i].(<-chan time.Time)):
						this.funcs[i]()
					default:
					}
				} else {
					select {
					case <-(this.chans[i].(chan time.Time)):
						this.funcs[i]()
					default:
					}
				} */
		}
	}
}

func main() {
	t := &timeManager{}
	t.setTimeOut(3, func() {
		fmt.Println("every 3")
	})
	t.setInterval(1.5, func() {
		fmt.Println("every 1.5")
	})
	t.run()
}
