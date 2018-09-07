package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func operate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a / b
		q, _ := div(a, b)
		return q, nil
	default:
		//panic("unsupport operation: " + op)
		return 0, fmt.Errorf(
			"unsupported operation: %s", op,
		)
	}
}

// 13/3 => 4....1
func div(a, b int) (int, int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()

	fmt.Printf("Calling %s with %d, %d\n",
		opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	result := int(math.Pow(float64(a), float64(b)))
	return result
}

func sum(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}

	return sum
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	if result, err := operate(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(div(13, 3))

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 3))

	fmt.Println(sum(7, 2, 3, 4, 5))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	c, d := 5, 6
	c, d = swap2(c, d)
	fmt.Println(c, d)

	fmt.Println("---------------------------------")

	aa := []int{10, 20, 30}
	var i int
	var ptr [3]*int

	for i = 0; i < 3; i++ {
		fmt.Printf("aa[%d] = %d\n", i, aa[i])
		ptr[i] = &aa[i]
		fmt.Printf("--aa[%d] = %d\n", i, *ptr[i])
	}
}
