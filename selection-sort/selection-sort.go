package main

import (
	"fmt"
	"math/rand"
	"time"
)

func selectionSort(number []int) []int {
	n := len(number)
	for i := 0; i < n; i++ {
		min := i
		for j := i; j < n; j++ {
			if number[j]<number[min]{
				min = j
			}
		}
		number[i], number[min] = number[min],number[i]
	}
	return number
}

func generateNumber(size int) []int {
	slice := make([]int,size,size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i]=rand.Intn(99)-rand.Intn(99)
	}
	return slice
}

func main() {
	// fmt.Println(selectionSort([]int{66,55,43,23,524,46,234,132,754,856,192,3,9,4,6,19}))
	number := generateNumber(10)
	fmt.Println("before sort:", number)
	selectionSort(number)
	fmt.Println("after sort", number)
}