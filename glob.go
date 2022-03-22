package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	files, _ := filepath.Glob("*")
	for i, name := range files {
		fmt.Println(i, "=", name)
	}
}
