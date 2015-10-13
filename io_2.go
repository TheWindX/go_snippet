package main

import (
	//"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"os"
)

func main() {
	file, err := os.Create("evil.png")
	if err != nil {
		panic("cannot create evil file")
	}
	resp, err := http.Get("https://www.baidu.com/img/bdlogo.png")
	if err != nil {
		panic("cannot get bdlogo file")
	}
	io.Copy(file, resp.Body)
}
