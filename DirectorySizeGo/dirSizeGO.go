package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

var (
	sizeDir       int64
	fileNum       int
	wg            = sync.WaitGroup{}
	chSize        = make(chan int64)
	chFileNum     = make(chan int)
	donechSize    = make(chan struct{})
	donechFileNum = make(chan struct{})
)

func main() {
	start := time.Now()
	wg.Add(1)
	go FileSize("C:/Users")
	go func() {
		for {
			select {
			case i := <-chSize:
				sizeDir += i
			case <-donechSize:
				close(chSize)
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case i := <-chFileNum:
				fileNum += i
			case <-donechFileNum:
				close(chFileNum)
				return
			}
		}
	}()

	wg.Wait()
	donechFileNum <- struct{}{}
	donechSize <- struct{}{}
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
		} else {
			chFileNum <- 1
		}
		size += f.Size()
	}
	chSize <- size
	wg.Done()
}
