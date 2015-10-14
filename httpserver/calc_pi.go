package main

import (
	"fmt"
	"runtime"
	//"math"
	"math/big"
)

const prec = 2000
const piece = 1000
const pieceLeng = 3000000

var resList chan float64

func bigNum(num float64) *big.Float {
	return new(big.Float).SetPrec(prec).SetFloat64(num)
}

func pieceWork(idx int) {
	go func() {
		_cur := 2 + idx*2*pieceLeng
		cur := float64(_cur)
		var sum float64 = 0
		posi := true
		for i := 0; i < pieceLeng; i++ {
			t := (float64(4) / (cur * (cur + float64(1)) * (cur + float64(2))))
			if posi {
				sum += t
			} else {
				sum -= t
			}
			posi = !posi
			//fmt.Println(t)
			cur += float64(2)
		}
		resList <- sum
	}()
}

var sumRes chan float64

func sumAll() {
	var sum float64 = 0
	for i := 0; i < piece; i++ {
		sum += <-resList
	}
	sumRes <- sum
}

func main() {
	runtime.GOMAXPROCS(4)
	resList = make(chan float64, piece)
	sumRes = make(chan float64)

	for i := 0; i < piece; i++ {
		pieceWork(i)
	}
	go sumAll()
	fmt.Printf("PI = %.50f\n", 3+<-sumRes)
}
