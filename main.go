package main

// inspired by https://gist.github.com/zorchenhimer/6d82758c54f16a02074e

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var size, max int
	flag.IntVar(&size, "size", 8, "size of slice to sort")
	flag.IntVar(&max, "max", 1000000, "max iterations to sort")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	data := generateRandomSlice(size)
	fmt.Println(data)

	start := time.Now()
	data = doSort(max, data)

	fmt.Println(time.Since(start))
	fmt.Println(data)
}

func doSort(max int, data []int) []int {
	var count int

	for !checkSort(data) {
		count++
		if count > max {
			// According to the original author, just give up
			// instead of calling sort.Ints()...it could be so simple
			fmt.Println("\nRED ALERT RED ALERT BAIL BAIL")
			break
		}
		fmt.Printf("%07d\r", count)

		data = bogoSort(data)
	}

	return data
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
