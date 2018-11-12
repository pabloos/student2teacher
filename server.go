package main

import (
	"log"
	"net/http"
)

const (
	protocol = ":http"
	htmldir  = "public"
)

type server struct {
	dirPath string
	router  *http.ServeMux
}

func newServer(dirPath string) *server {
	s := &server{
		dirPath: dirPath,
		router:  http.NewServeMux(),
	}

	defer s.routes()

	return s
}

func (s *server) routes() {
	log.Print("Setting routes")
	files := http.FileServer(http.Dir(htmldir))

	s.router.Handle("/", files)
	s.router.HandleFunc("/upload", getUploadHandler(s.dirPath))
}

func (s *server) serve() {
	log.Print("Starting the server")
	log.Fatal(http.ListenAndServe(protocol, s.router))
}
