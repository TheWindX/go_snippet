package utils

import (
	"fmt"
	"time"
)

func PrintNow() {
	now := time.Now()
	fmt.Printf("start: %02d:%02d:%02d. %03d\n",
		now.Hour(), now.Minute(), now.Second(), (now.Nanosecond()%1e9)/1e6)
}
