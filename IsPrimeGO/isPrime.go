package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	flag bool
	wg   = sync.WaitGroup{}
)

// Provide args from 3 to 5 to check performance increase

func main() {
	// var numLarge uint64 = 9999999992999999999
	var maxNum int = 100000 // Max count : 1000000
	primes := getAllPrime(maxNum)
	args := os.Args
	routines, _ := strconv.Atoi(args[1])

	// Sequential Execution
	start := time.Now()
	var result bool
	for _, num := range primes {
		n, _ := strconv.Atoi(num)
		result = isPrime(uint64(n))
	}
	// result = isPrime(numLarge)
	spentSeq := time.Since(start)
	// This is the most optimised algorithm for checking prime sequentially
	fmt.Printf("For Sequential Execution time taken is %v and result is %v\n", spentSeq, result)

	// Concurrent Execution
	start = time.Now()
	for _, num := range primes {
		n, _ := strconv.Atoi(num)
		isPrimeGo(uint64(n), routines)
		flag = false
	}
	// isPrimeGo(numLarge, routines)
	spentConc := time.Since(start)
	// And Go kills is with concurrent execution
	fmt.Printf("For Concurrent Execution time taken is %v and result is %v\n", spentConc, !flag)
}
func getAllPrime(maxNum int) []string {

	dat, err := ioutil.ReadFile("primes50.txt")
	if err != nil {
		fmt.Println(err)
	}

	var primes []string = strings.Fields(string(dat))
	return primes[:maxNum]

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

func isPrime(n uint64) bool {
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
	var i uint64
	for i = 6; i <= sqrtn; i = i + 6 {
		if n%(i-1) == 0 || n%(i+1) == 0 {
			return false
		}
	}
	return true

}
