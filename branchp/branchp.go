package branchp

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	NUM_ELEMENTS   = 32767
	NUM_ITERATIONS = 10000
)

var (
	arr, sorted_arr []int
)

func initialize() {
	arr = make([]int, NUM_ELEMENTS)
	for i := 0; i < NUM_ELEMENTS; i++ {
		arr[i] = rand.Intn(1000)
	}
	sorted_arr = make([]int, NUM_ELEMENTS)
	copy(sorted_arr, arr)
	sort.Ints(sorted_arr)
}

func test(target []int) {
	start := time.Now()
	var dummy int
	for i := 0; i < NUM_ITERATIONS ; i++ {
		dummy+= doIt(target)
	}
	fmt.Printf("Dummy(%d) Took %d ns\n", dummy, time.Now().Sub(start).Nanoseconds())
}

func doIt(target []int) (result int) {
	for _, x := range target {
		if x < 500 {
			result++
		}
	}
	return result
}

func Exec() {
	initialize()
	test(sorted_arr)
	test(arr)
	test(sorted_arr)
	test(arr)
}
