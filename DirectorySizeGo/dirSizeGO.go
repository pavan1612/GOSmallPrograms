package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

var sizeDir int64
var chSize = make(chan int64, 100)
var fileNum int

var wg = sync.WaitGroup{}

func main() {
	start := time.Now()
	wg.Add(1)
	go FileSize("C:/Users/pavan")
	go func() {
		for i := range chSize {
			sizeDir += i
		}
	}()

	wg.Wait()
	spent := time.Since(start)
	fmt.Printf("The size of the directory is %v MB\nAnd Time spent is %v ms\nNumber of Files :%v", ((sizeDir / 1024) / 1024), spent.Milliseconds(), fileNum)
}

func FileSize(root string) {
	var size int64
	files, err := ioutil.ReadDir(root)
	if err != nil {
		// log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			wg.Add(1)
			go FileSize(root + "/" + f.Name())
		}
		fileNum++
		size += f.Size()
	}
	chSize <- size
	wg.Done()
}
