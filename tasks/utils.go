package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"
)

// GoFiles returns all files in the directory from the eventPath.
func GoFiles(eventPath string) string {
	dir := DirFromPath(eventPath)
	list, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var files bytes.Buffer
	for _, file := range list {
		if strings.Contains(file.Name(), ".go") {
			files.WriteString(dir + file.Name() + " ")
		}
	}
	return files.String()
}

// DirFromPath removes the file at the end and returns the path.
func DirFromPath(path string) string {
	slice := strings.Split(path, "/")
	dir := "/" + strings.Join(slice[1:len(slice)-1], "/") + "/"
	return dir
}
