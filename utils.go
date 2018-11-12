package main

import (
	"fmt"
	"log"
	"os"
)

const (
	slash           = "/"
	filePermissions = 0755
)

func endsWithSlash(path string) bool {
	lastChar := path[len(path)-1:]
	if lastChar == slash {
		return true
	}

	return false
}

func correctPath(path string) string {
	if !endsWithSlash(path) {
		return fmt.Sprintf("%s%s", path, slash)
	}

	return path
}

func createDir(path string) {
	if !dirExists(path) {
		err := os.Mkdir(path, filePermissions)

		if err != nil {
			log.Print(err)
		}
	}
}

func dirExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}
