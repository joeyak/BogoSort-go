package main

// inspired by https://gist.github.com/zorchenhimer/6d82758c54f16a02074e

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var size, max, cores int
	flag.IntVar(&size, "size", 8, "size of slice to sort")
	flag.IntVar(&max, "max", 1000000, "max iterations to sort")
	flag.IntVar(&cores, "cores", 1, "number of goroutines to use")
	flag.Parse()

	if max == 0 {
		max = math.MaxInt32
	}

	rand.Seed(time.Now().UnixNano())

	data := generateRandomSlice(size)
	fmt.Println(data)

	start := time.Now()
	data = doSort(max, cores, data)

	fmt.Println(time.Since(start))
	fmt.Println(data)
}

func doSort(max, cores int, data []int) []int {
	// Don't you love struct channels. I love them 💜!
	// 😢 I didn't get to use the struct channels because
	// returning early solved the problem
	ch := make(chan []int)
	var wg sync.WaitGroup

	for i := 0; i < cores; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < max; c++ {
				newData := bogoSort(data)
				if checkSort(newData) {
					ch <- newData
					break
				}
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		ch <- nil
	}()
	return <-ch
}

func generateRandomSlice(count int) []int {
	var data []int
	for i := 0; i < count; i++ {
		data = append(data, rand.Intn(100))
	}
	return data
}

func checkSort(data []int) bool {
	a := math.MinInt32
	for _, b := range data {
		if a == math.MinInt32 {
			a = b
		} else if b >= a {
			a = b
		} else {
			return false
		}
	}
	return true
}

func bogoSort(data []int) []int {

	var addedIndex []int
	var sortedData []int
	for len(data) != len(sortedData) {

		idx := rand.Intn(len(data))
		var alreadyAdded bool
		for i := range addedIndex {
			if idx == addedIndex[i] {
				alreadyAdded = true
				break
			}
		}

		if alreadyAdded {
			continue
		}
		addedIndex = append(addedIndex, idx)
		sortedData = append(sortedData, data[idx])
	}
	return sortedData
}
