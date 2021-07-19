package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var size int64
var fileNum int

func main() {
	var err error
	var root string = "C:/Users/pavan"
	start := time.Now()
	err = FilePathWalkDir(root)
	spent := time.Since(start)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The size of the directory is %v MB\nAnd Time spent is %v ms\nNumber of Files :%v", ((size / 1024) / 1024), spent.Milliseconds(), fileNum)

}
func FilePathWalkDir(root string) error {
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileNum++
			size += info.Size()
		}
		return nil
	})
	return err
}
