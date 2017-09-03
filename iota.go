package main

import "fmt"

const(
	B float64 = 1 << (iota*10)
	KB
	MB
	GB
)

func main() {
	fmt.Println(GB)
}