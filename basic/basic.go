package main

import "fmt"
import "math/cmplx"
import "math"

var (
	aa = 3
	ss = "kkk"
	zz = true
)

func variableZeroValue() {
	var a, b int = 1, 3
	d := true
	d = false
	var s string = "zyh"
	var c = "hello"
	fmt.Println(a, s, b, c, d)
}

func eular() {
	// c := 3 + 4i
	// fmt.Println(cmplx.Abs(c))

	d := cmplx.Pow(math.E, 1i*math.Pi) + 1
	fmt.Println(d)
}

func tranfr() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))

	fmt.Println(c)
}

func consts() {
	const filename string = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))

	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		java
		golang
		python
	)
	fmt.Println(cpp, java, golang, python)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	fmt.Println(aa, ss, zz)
	eular()
	tranfr()
	consts()
	enums()


	fmt.Println("------------------------------------------")
	s := "xx"
	title := fmt.Sprintf("%s/helloworld", s)
	//title = fmt.Sprintln(s)
	fmt.Println(title)
}
