package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	files, _ := filepath.Glob(".html")
	for _, fname := range files {
		fileReplace(fname, "hoge@example.com", "fuga@example.com")
	}
}

func fileReplace(fname, src, dst string) {
	bytes, _ := ioutil.ReadFile(fname)
	lines := strings.Replace(string(bytes), src, dst, -1)
	result := []byte(lines)
	ioutil.WriteFile(fname, []byte(result), 0666)
	fmt.Println("ok:", fname)
}
