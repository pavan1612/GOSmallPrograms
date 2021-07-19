package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	flag bool
	wg   = sync.WaitGroup{}
)

// Provide args from 3 to 5 to check performance increase

func main() {
	var num uint64 = 9999999992999999999
	args := os.Args
	routines, _ := strconv.Atoi(args[1])

	// Sequential Execution
	start := time.Now()
	result := isPrime(num, routines)
	spentSeq := time.Since(start)
	// This is the most optimised algorithm for checking prime sequentially
	fmt.Printf("For Sequential Execution time taken is %v and result is %v\n", spentSeq, result)

	// Concurrent Execution
	start = time.Now()
	isPrimeGo(num, routines)
	spentConc := time.Since(start)
	// And Go kills is with sequential execution
	fmt.Printf("For Concurrent Execution time taken is %v and result is %v\n", spentConc, !flag)
}
func isPrimeGo(n uint64, rt int) {
	if n < 2 {
		flag = true
		return
	}
	if n == 2 || n == 3 {
		return
	}
	if n%2 == 0 || n%3 == 0 {
		flag = true
		return
	}
	var sqrtn uint64 = uint64(math.Sqrt(float64(n))) + 1
	var sqrtnPart uint64 = sqrtn / uint64(rt)
	var start uint64 = 0

	for i := 0; i < rt; i++ {
		end := start + sqrtnPart
		wg.Add(1)
		go checkPartModulusGo(n, start, end)
		start = end
	}
	wg.Wait()

}

func checkPartModulusGo(n, start, end uint64) {
	for i := start + 6; i <= end; i = i + 6 {
		if flag {
			wg.Done()
			return
		}
		if n%(i-1) == 0 || n%(i+1) == 0 {
			flag = true
		}
	}
	wg.Done()
}

func isPrime(n uint64, rt int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	var sqrtn uint64 = uint64(math.Sqrt(float64(n))) + 1
	var sqrtnPart uint64 = sqrtn / uint64(rt)
	var start uint64 = 0
	for i := 0; i < rt; i++ {
		end := start + sqrtnPart
		if !checkPartModulus(n, start, end) {
			return false
		}
		start = end
	}
	return true

}

func checkPartModulus(n, start, end uint64) bool {
	for i := start + 6; i <= end; i = i + 6 {
		if n%(i-1) == 0 || n%(i+1) == 0 {
			return false
		}
	}
	return true
}