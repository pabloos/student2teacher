package main

import (
	"flag"
)

const (
	defaultDir = "home/"
)

func main() {
	path := flag.String("path", defaultDir, "path to save files")

	flag.Parse()

	pathCorrected := correctPath(*path)

	server := newServer(pathCorrected)

	server.serve()
}
