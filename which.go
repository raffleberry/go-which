package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func foundAll(nFiles, nFound int) bool {
	return nFiles == nFound
}

func print(fullPath string, file string, nFiles int) {
	if nFiles == 1 {
		fmt.Println(fullPath)
		return
	}
	fmt.Printf("%s\t%s\n", file, fullPath)
}

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("Need 1 or more arguments")
		os.Exit(1)
	}

	found := 0
	files := args[1:]
	nFiles := len(files)
	path := os.Getenv("PATH")
	dirs := filepath.SplitList(path)

	for _, file := range files {
		for _, dir := range dirs {
			fullPath := filepath.Join(dir, file)
			fileInfo, err := os.Stat(fullPath)
			if err == nil {
				mode := fileInfo.Mode()
				if mode.IsRegular() {
					if mode&0111 != 0 {
						found++
						done := foundAll(nFiles, found)
						print(fullPath, file, nFiles)
						if done {
							os.Exit(0)
						}
						break
					}
				}
			}
		}
	}

	os.Exit(1)

}
