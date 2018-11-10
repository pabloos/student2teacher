package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//funcion para manejar el archivo subido
func upload(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Panic(err)
	}

	defer file.Close()

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, file)

	err = os.Mkdir(path, 0755)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(path+handler.Filename, buf.Bytes(), 0755)
	if err != nil {
		log.Panic(err)
	}
}
