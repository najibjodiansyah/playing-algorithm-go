package main

import (
	"fmt"
	"strings"
)

func reverseArr(sentence string) string {
	word := strings.Split(sentence, "")
	result := ""
	for i := len(sentence)-1; i >= 0; i-- {
		result+=word[i]
	}
	return result
}

func reverseInt(arr []int) []int {
	result := []int{}
	for i := len(arr)-1; i >= 0; i-- {
		result = append(result, arr[i])
	}
	return result
}

func main() {
	fmt.Println(reverseArr("Najibuddin Jodiansyah"))
	fmt.Println(reverseInt([]int{1,2,3,4,5}))
}