package main

import (
	"fmt"
	"practice/mysort"
	"time"
)

var sortTestSli = []int{3, 5, 2, 1, 20, 6, 8, 7, 9, 4, 10, 19, 18, 16, 15, 13, 14, 12, 11, 17}

func main() {
	mergeSortTest()
	sleepSortTest()
}

func mergeSortTest() {
	var timeStart = time.Now()
	var res = mysort.MergeSort(sortTestSli)
	var timeEnd = time.Now()

	var timeElapsed = timeEnd.Sub(timeStart).Nanoseconds() / 1000000

	fmt.Println(`Merge Sort:`, res, "time:", timeElapsed, "ms")
}

func sleepSortTest() {
	var timeStart = time.Now()
	var res = mysort.SleepSort(sortTestSli)
	var timeEnd = time.Now()

	var timeElapsed = timeEnd.Sub(timeStart).Nanoseconds() / 1000000

	fmt.Println(`Sleep Sort:`, res, "time:", timeElapsed, "ms")
}
