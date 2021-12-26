package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(number []int) []int {
	if len(number)<2 {
		return number
	}
	left , right := 0, len(number)-1
	center := rand.Int() % len(number)

	number[center], number[right] = number[right], number[center]

	for i, _ := range number {
		if number[i]<number[right]{
			number[left],number[i] = number[i],number[left]
			left++
		}
	}
	number[left], number[right] = number[right],number[left]

	quickSort(number[:left])
	quickSort(number[left+1:])

	return  number
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
	quickSort(number)
	fmt.Println("after sort", number)
}