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

func cmd_stream() (quit chan bool, next chan bool) {
	quit = make(chan bool)
	next = make(chan bool)
	go func() {
		for {
			in := ""
			fmt.Scanf("%s", &in)
			if in == "q" {
				quit <- true
			} else if in == "n" {
				next <- true
			}
		}
	}()
	return
}

func main() {
	flines := file_stream("fstream.data")
	is_quit, is_next := cmd_stream()
	for {
		select { //多数异步选择
		case _ = <-is_quit: //退出
			return
		case _ = <-is_next: //下一句
			l := <-flines
			fmt.Printf(l)
		}
	}
}
