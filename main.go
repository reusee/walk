package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	paths := os.Args[1:]
	if len(paths) == 0 {
		paths = append(paths, ".")
	}
	for _, path := range paths {
		walk(path)
	}
}

func walk(path string) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	fmt.Println(path)
	defer f.Close()
	for {
		infos, err := f.Readdir(512)
		for _, info := range infos {
			if info.IsDir() {
				walk(filepath.Join(path, info.Name()))
			} else {
				fmt.Println(filepath.Join(path, info.Name()))
			}
		}
		if err != nil {
			break
		}
	}
}
