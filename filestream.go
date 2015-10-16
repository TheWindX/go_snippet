// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"math/rand"
	//"runtime"
	//"time"
)

func file_stream(fpath string) chan string {
	s := make(chan string)
	go func() {
		file, err := os.Open(fpath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			s <- scanner.Text() //这里读取一行
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()
	return s
}

func main() {
	ch := file_stream("fstream.data")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
