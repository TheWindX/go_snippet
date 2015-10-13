package main

import (
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var file *os.File
	var resp *http.Response
	var succ bool
	var err error

	fio := make(chan bool)
	go func() {
		file, err = os.Create("evil.png")
		if err != nil {
			fio <- false
			return
		}
		fio <- true
	}()

	succ = <-fio

	if succ {
		netio := make(chan bool)
		go func() {
			resp, err = http.Get("https://www.baidu.com/img/bdlogo.png")
			if err != nil {
				netio <- false
				return
			}
			netio <- true
		}()
		succ := <-netio
		if succ {
			io.Copy(file, resp.Body)
		} else {
			fmt.Println("cannot get file \"bdlogo.png\"")
		}
	} else {
		fmt.Println("cannot create \"evil\"")
	}
}
