package main

import (
	"bufio"
	"errors"
	"fmt"
	"golearn/functional/fib"
	"os"
)

func tryDefer() {
	//Like stack, 3 2 1
	// defer fmt.Println(1)
	// fmt.Println(11)
	// defer fmt.Println(2)
	// fmt.Println(3)
	// //panic("err occurred")
	// fmt.Println(4)

	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("Print too words")
		}
	}
}

func writeFile(filename string) {
	//file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	err = errors.New("this is a custom error")
	if err != nil {
		//panic(err)
		// fmt.Println("Error: ", err.Error())
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Finonacci()
	for i := 0; i < 30; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	writeFile("abc.txt")
}
