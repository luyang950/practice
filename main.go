package main

import (
	"fmt"
	"practice/mysort"
)

func main() {
	var input = []int{3, 5, 2, 1, 6, 8, 7, 9, 4, 10, 19, 18, 16, 15, 13, 14 ,12, 11, 17}

	var res = mysort.MergeSort(input)

	fmt.Println(`res-------->`, res)
}
