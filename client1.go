package main

import (
	"./ratelimit"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

var exit chan bool
var fpath chan string

func report() (exit chan bool, fpath chan string) {
	exit = make(chan bool)
	fpath = make(chan string)
	go func() {
		files := []string{}
		for {
			select {
			case <-exit:
				break
			case f := <-fpath:
				files = append(files, f)
			default:
				time.Sleep(time.Duration(float64(time.Second) * float64(2)))
				for _, file := range files {
					fi, _ := os.Stat(file)
					fmt.Println(file + " 已下载字节:" + strconv.FormatInt(fi.Size(), 10))
				}
				fmt.Println("\n\n")
			}
		}
	}()
	return
}

func fileLoader(url string, fname string) chan bool {
	done := make(chan bool)
	go func() {
		os.Mkdir("mydown", os.ModePerm)
		out, _ := os.Create("mydown/" + fname)
		defer out.Close()
		fpath <- "mydown/" + fname
		resp, _ := http.Get(url)
		defer resp.Body.Close()
		bucket := ratelimit.NewBucketWithRate(100*1024, 100*1024)
		io.Copy(out, ratelimit.Reader(resp.Body, bucket))
		done <- true
	}()
	return done
}

func main() {
	runtime.GOMAXPROCS(10)
	exit, fpath = report()

	d1 := fileLoader("http://www.sina.com.cn", "sina.html")
	d2 := fileLoader("http://www.sohu.com", "sohu.html")
	d3 := fileLoader("http://www.163.com", "163.html")

	/* d2 := fileLoader("2.7z")
	d3 := fileLoader("3.7z")
	d4 := fileLoader("4.7z")
	<-d1
	<-d2
	<-d3
	<-d4 */
	<-d1
	<-d2
	<-d3
	exit <- true
}
