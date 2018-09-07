package main

import (
	"fmt"
)

func updateSlice(a []int) {
	a[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])

	s1 := arr[2:6]
	//s2 := arr[:]
	s3 := s1[3:5] //slice extension

	//updateSlice(s1)
	//updateSlice(s2)

	fmt.Println(arr, s1, s3)

	fmt.Println("======================")
	s4 := append(s3, 10)
	s5 := append(s4, 11)
	s6 := append(s5, 12)
	fmt.Println(s4, s5, s6)

	// s5, s6 no longer view arr
	fmt.Println("arr = ", arr)

}
