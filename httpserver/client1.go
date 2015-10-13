package main

import (
	"io"
	"net/http"
	"os"
)

func f1(fname string) {
	os.Mkdir("mydown", os.ModePerm)
	out, _ := os.Create("mydown/" + fname)
	defer out.Close()

	resp, _ := http.Get("http://127.0.0.1:8888/files/" + fname)
	defer resp.Body.Close()

	_, err := io.Copy(out, resp.Body)
	if err != nil {

	}
}

func main() {
	f1("1.7z")
}
