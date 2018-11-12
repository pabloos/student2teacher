package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func getUploadHandler(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			file, handler, err := r.FormFile("file")
			if err != nil {
				log.Print(err)
			}

			log.Printf("storing %s file", handler.Filename)

			defer file.Close()

			buf := bytes.NewBuffer(nil)
			io.Copy(buf, file)

			createDir(path)

			err = ioutil.WriteFile(fmt.Sprintf("%s%s", path, handler.Filename), buf.Bytes(), 0755)
			if err != nil {
				log.Print(err)
			}

			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}
}
