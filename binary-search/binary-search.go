package main

import "fmt"

func binarySearch(arr []int, target int) bool {
	low := 0
	high := len(arr)-1

	for low <= high {
		middle := (low+high)/2

		if arr[middle] < target {
			low = middle + 1
		}else {
			high = middle - 1
		}
	}
	if low == len(arr) || arr[low] != target {
		return false
	}
	return true
}

func main() {
	fmt.Println(binarySearch([]int{1,2,3,4,5,6,7,8,9,10,11,12},5))
}