package main

import (
	"fmt"
	"strconv"
)

func printArray(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)

	fmt.Println(grid)

	for i, v := range arr3 {
		fmt.Println(strconv.Itoa(i) + ": " + strconv.Itoa(v))
	}

	for _, v := range arr3 {
		fmt.Println(v)
	}

	printArray(arr1[:])
	printArray(arr3[:])

	fmt.Println(arr1, arr3)
	//printArray(arr2)
}
