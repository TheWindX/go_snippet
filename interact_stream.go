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

func cmd_stream() chan string {
	cmd := make(chan string)
	go func() {
		for {
			in := ""
			fmt.Scanf("%s", &in)
			cmd <- in
		}
	}()
	return cmd
}

func main() {
	flines := file_stream("fstream.data")
	cmds := cmd_stream()
	for {
		cmd := <-cmds
		if cmd == "quit" {
			break
		} else if cmd == "n" {
			l := <-flines
			fmt.Printf(l)
		}
	}
}
