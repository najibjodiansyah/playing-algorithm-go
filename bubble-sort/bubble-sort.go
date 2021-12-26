package main

import "fmt"

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr)-i-1; j++ {
			if arr[j]>arr[j+1]{
				arr[j], arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
	return arr
}

func main() {
	fmt.Println(bubbleSort([]int{8,0 ,-49 ,-67 ,-29 ,-18 ,-21 ,-23 ,14 ,-47}))
}