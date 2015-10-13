package main

import (
	"./ratelimit"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

var f, _ = os.Create("data.copy")

func report() {
	for {
		time.Sleep(time.Duration(float64(1) * float64(time.Second)))
		fi, err := f.Stat()
		if err != nil {
			// Could not obtain stat, handle error
		}

		fmt.Printf("The file is %d bytes long", fi.Size())

		fmt.Println(time.Now())
	}

}

func ftrans() {
	// Source holding 1MB
	//var f, _ = os.Create("data.copy")
	defer f.Close()

	w := bufio.NewWriter(f)

	src := bytes.NewReader(make([]byte, 1024*1024))
	runtime.GOMAXPROCS(1)
	go report()

	bucket := ratelimit.NewBucketWithRate(100*1024, 100*1024)
	start := time.Now()
	io.Copy(w, ratelimit.Reader(src, bucket))
	fmt.Printf("Copied %d bytes in %s\n", src.Len(), time.Since(start))
	// Destination

	// Bucket adding 100KB every second, holding max 100KB

	// Copy source to destination, but wrap our reader with rate limited one
	//io.Copy(dst, ratelimit.Reader(src, bucket))

	//fmt.Printf("Copied %d bytes in %s\n", dst.Len(), time.Since(start))
}

func helloHandler(rw http.ResponseWriter, req *http.Request) {
	url := req.Method + " " + req.URL.Path
	if req.URL.RawQuery != "" {
		url += "?" + req.URL.RawQuery
	}
	log.Println(url)
	io.WriteString(rw, "hello world")
}

func fileserver() {
	fs := http.FileServer(http.Dir("/dir"))
	http.Handle("/static", http.StripPrefix("/dir", fs))
	err := http.ListenAndServe(":8080", fs)
	if err != nil {
		log.Fatal("...")
	}
}

func main() {
	fileserver()
}
