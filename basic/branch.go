package main

import (
	"fmt"
	"io/ioutil"
)

func eval(a int) string {
	g := ""
	switch {
	case a < 60:
		g = "F"
	case a < 80:
		g = "C"
	case a < 90:
		g = "B"
	case a <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Worng score: %d", a))
	}

	return g
}

func main() {
	const filename = "note.md"
	contents, err := ioutil.ReadFile(filename)
	// if contents, err := ioutil.ReadFile(filename); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("%s\n", contents)
	// }

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		eval(70),
		eval(50),
		eval(85),
		eval(100),
		eval(100))
}
