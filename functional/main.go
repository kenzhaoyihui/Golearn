package main

import (
	"bufio"
	"fmt"
	"golearn/functional/fib"
	"io"
	"strings"
)

// func finonacci() intGen {
// 	a, b := 0, 1
// 	return func() int {
// 		a, b = b, a+b
// 		return a
// 	}
// }

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()

	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	// TODO: incorect if the p is too small
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	//func() int does not implement io.Reader (missing Read method)
	//f := fib.Finonacci()

	var f intGen = fib.Finonacci()

	printFileContents(f)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }
}
