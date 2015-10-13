package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("Dir"))))
	http.HandleFunc("/hello", myhandler)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("Error listening: ", err)
	}
}

func myhandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello!")
}
