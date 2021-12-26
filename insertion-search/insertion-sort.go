package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertionSort(number []int) []int {
	length := len(number)
	for i := 1; i < length; i++ {
		j := i
		for j > 0 {
			if number[j-1] > number[j]{
				number[j], number[j-1] = number[j-1], number[j]
			}
			j = j - 1
		}
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
	insertionSort(number)
	fmt.Println("after sort", number)
}