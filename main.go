package main

import (
	"flag"
	"fmt"
	"net/http"
)

var path string

//

func main() {
	files := http.FileServer(http.Dir("public"))
	path := flag.String("path", "/home/", "path to save files")

	http.Handle("/", files)
	http.HandleFunc("/upload", upload)

	flag.Parse()

	fmt.Println(*path)
	http.ListenAndServe(":80", nil)
}
