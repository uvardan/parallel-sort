package bitonicsort

import (
	"math/rand"
	"sync"
	"time"
)

const (
	ASC  bool = true
	DESC bool = false
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Sort sorts an array in place using the parallel bitonic algorithm.
//
// It takes an array as an input.
// It returns the input array sorted.
func Sort(arr []int) []int {
	bitonic_sort(arr, ASC)
	return arr
}

func bitonic_sort(arr []int, orderby bool) {
	if len(arr) < 2 {
		return
	}

	middle := len(arr) / 2
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		bitonic_sort(arr[:middle], ASC)
	}()

	go func() {
		defer wg.Done()
		bitonic_sort(arr[middle:], DESC)
	}()
	wg.Wait()
	bitonic_merge(arr, orderby)
}

func bitonic_compare(arr []int, orderby bool) {
	middle := len(arr) / 2
	for i := 0; i < middle; i++ {
		if (arr[i] > arr[i+middle]) == orderby {
			arr[i], arr[i+middle] = arr[i+middle], arr[i]
		}
	}
}

func bitonic_merge(arr []int, orderby bool) {
	bitonic_compare(arr, orderby)
	middle := len(arr) / 2
	if middle > 1 {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			bitonic_merge(arr[:middle], orderby)
		}()
		go func() {
			defer wg.Done()
			bitonic_merge(arr[middle:], orderby)
		}()
		wg.Wait()

	}
}
