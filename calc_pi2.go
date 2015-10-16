package main

import (
	"fmt"
	"runtime"
	//"math"
	"math/big"
)

const prec = 2000
const piece = 600
const pieceLeng = 10000

var resList chan *big.Float
var sumRes chan *big.Float

func bigInt(num int) *big.Float {
	return new(big.Float).SetPrec(prec).SetInt64(int64(num))
}

func bigFloat(num float64) *big.Float {
	return new(big.Float).SetPrec(prec).SetFloat64(num)
}

func pieceWork(idx int) {
	go func() {
		cur := 2 + idx*2*pieceLeng
		posi := true
		sum := bigInt(0)
		for i := 0; i < pieceLeng; i++ {
			t := bigInt(4)
			t1 := bigInt(cur * (cur + 1) * (cur + 2))
			t.Quo(t, t1)
			//fmt.Printf("%.15f\n", t)
			if posi {
				sum.Add(sum, t)
			} else {
				sum.Sub(sum, t)
			}
			posi = !posi
			cur += 2
		}
		resList <- sum
	}()
}

func sumAll() {
	sum := bigInt(3)
	for i := 0; i < piece; i++ {
		sum.Add(sum, <-resList)
	}
	sumRes <- sum
}

func main() {
	runtime.GOMAXPROCS(15)
	resList = make(chan *big.Float, piece)
	sumRes = make(chan *big.Float)

	for i := 0; i < piece; i++ {
		pieceWork(i)
	}
	go sumAll()
	fmt.Printf("PI = %.50f\n", <-sumRes)
}
