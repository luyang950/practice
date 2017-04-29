package main

import (
	"fmt"
	"math/rand"
	"practice/mysort"
	"time"
)

const sortSliLength = 50

var sortTestSli []int

func main() {
	sortTestSli = genRandSli(sortSliLength)
	fmt.Println("Input slice:", sortTestSli)

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

func genRandSli(num int) []int {
	var resSli []int
	var avoidDupMap = make(map[int]bool)

	for i := 0; i < num; i++ {
		var resNum = rand.Intn(num)
		if _, ok := avoidDupMap[resNum]; !ok {
			avoidDupMap[resNum] = true
			resSli = append(resSli, resNum)
		} else {
			i--
			continue
		}
	}

	return resSli
}
