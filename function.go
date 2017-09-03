package main

import "fmt"


func main(){
	a,b,c,d,e := 1,2,3,4,5
	A(a,b,c,d,e)
	fmt.Println(a,b,c,d,e)
}

func A(a ...int){ //slice
	a[0] = 2
	a[1] = 4
	a[2] = 6
	a[3] = 8
	a[4] = 10
	fmt.Println(a)

}


/*
func main(){
	s1 := []int{1,2,3,4,5}
	A(s1)
	fmt.Println(s1)
}

func A(a []int){
	a[0] = 2
	a[1] = 4
	a[2] = 6
	a[3] = 8
	a[4] = 10
	fmt.Println(a)
}
*/

/*
func main(){
	a := func (){
		fmt.Println("Good Job!...")
	}
	a()

	f := close(10)
	fmt.Println(f(1))
	fmt.Println(f(2))

	//a := 1
	//A(&a)
	//fmt.Println(a)

}

func close(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func (y int) int{
		fmt.Printf("%p\n", &x)
		return x + y
	}
	//fmt.Println("Good Job!..")
	//*a = 2
	//fmt.Println(*a)
}
*/