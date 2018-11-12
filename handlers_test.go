package main

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_getUploadHandler(t *testing.T) {
	type args struct {
		path string
	}

	tests := []struct {
		name     string
		args     args
		filename string
		want     http.HandlerFunc
	}{
		{
			name: "basic case",
			args: args{
				path: "dir",
			},
			filename: "main.go",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := newfileUploadRequest("/upload", "file", tt.filename)
			if err != nil {
				t.Error(err)
			}

			response := httptest.NewRecorder()

			path := correctPath(tt.args.path) //path maybe needs to be corrected

			uploadHandler := getUploadHandler(path)

			http.HandlerFunc(uploadHandler).ServeHTTP(response, req)

			if status := response.Code; status != http.StatusOK {
				t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
			}

			if !fileExists(path) {
				t.Error("file wasn't stored")
			}

			os.RemoveAll(path) //backwards the
		})
	}
}

//based on https://gist.github.com/mattes/d13e273314c3b3ade33f
func fileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

//based on https://gist.github.com/mattetti/5914158/f4d1393d83ebedc682a3c8e7bdc6b49670083b84
func newfileUploadRequest(uri, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	defer file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())

	if err != nil {
		return nil, err
	}

	part.Write(fileContents)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())
	return request, nil
}
