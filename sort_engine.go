package main

import (
	"math/rand"
	"sort"
)

// Sort - Sort the array passed with bogo-sort mudda
func Sort(arr []int) []int {

	arr = shuffle(arr)

	arr = work(arr)

	return arr
}

func work(arr []int) []int {

	for {
		if sort.IntsAreSorted(arr) {
			return arr
		}

		arr = shuffle(arr)

		incrementCounter(ctx)
	}
}

func shuffle(arr []int) []int {

	mu.Lock()
	defer mu.Unlock()

	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
