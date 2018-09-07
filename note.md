## Inner variable type
* bool, string
* (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
* byte, rune
* float32, float64, complex64, complex128

## type cast
```
func tranfr() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))

	fmt.Println(c)
}
```

## const
* `const a,b = 3, 4`
* iota 
```
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
```

## if and switch
* `if...else...`
```
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
```

* `switch...case...`
```
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
```

## for
```
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}
```

## function
```
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

```

## Pointer (work with the value transfer)

swap
```
func swap(a, b *int) {
	*a, *b = *b, *a
}

func swap2(a, b int) (int, int) {
	return b, a
}
```

## Array
```
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	var grid [4][5]int
```

## Slice
```
func printArray(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

s1 := arr[2:6]
s2 := arr[:]
```
###  slice extension

```
s1 := arr[2:6]
s2 := s1[3:5]
```


```
s1 := []int{2, 4, 6, 8}
printSlice(s1)

s2 := make([]int, 16)
s3 := make([]int, 10, 32)

s2 = append(s2[:3], s2[4:]...)

```

## Map
```
	m := map[string]string{
		"id":   "1",
		"age":  "23",
		"name": "zyh",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3== nil
```

## OO(Object Oriented)
* support `package`, not support `Inheritance and polymorphism`, it implement by `interface`
* no `class`, just `struct`


### value receiver and pointer receiver
* use pointer receiver if changing the value
* use pointer receiver if the struct is big
* consistency: if have the pointer receiver, perfer to use pointer receiver


## Packaging
* name using `CamelCase`
* if first word is Upper, it is `public`, like `Print`
* if first world is Lower, it is `private`, like `print`

## Package extension
#### How to extend the system type or types which defined by others
* define the alias
```
type myTreeNode struct {
	node *tree.TreeNode
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}
```

* using the combine
```
package queue

type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
```

## gopm
* `go get github.com/gpmgo/gopm`


## duck typing
* need more than one interface
* need the python or C++ flexibility
* need the java type check

## interface
* implement interface ,just need to implement method in the interface.
* interface variable has the self-pointer
* interface valiable use the value transfer, not need the interface pointer
* pointer receiver implement just use the pointer; value receiver implement also can use the value or pointer 
* `interface{}` respect the all types
```
type Retiever struct {
	UserAgent string
	Timeout   time.Duration
}

func (r Retiever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}

	return string(result)
}
```

```
type Retiever struct {
	Contents string
}

func (r Retiever) Get(url string) string {
	return r.Contents
}
```

### Type assertion -- r.(type)
```
realRetiever := r.(*real.Retiever)
fmt.Println(realRetiever.Timeout)

func inspect(r Retiever) {
	fmt.Printf("(%T, %v)\n", r, r)
	switch v := r.(type) {
	case mock.Retiever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retiever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}
```


## interface combine
```
type Retiever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetieverPoster interface {
	Retiever
	Poster
	//Connect(host string)
}
```

implement method Get() and Post()
```
package mock

type M struct {
	Contents string
}

func (r *M) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}
func (r *M) Get(url string) string {
	return r.Contents
}
```

## System import interface
* `Stringer`
* `Reader`
* `Writer`
```
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}


aa := `asd"dsd
qwe
asdasf

23`

printFileContents(strings.NewReader(aa))

```

## Functional programming
* param, variable, return are also can be the function
* High order function
* Function -> Closure

fibonacci()
```
func finonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

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
	f := finonacci()

	printFileContents(f)
```


```
func (node *TreeNode) Traverse() {
	// if node == nil {
	// 	return
	// }
	// node.Left.Traverse()
	// node.Print()
	// node.Right.Traverse()
	node.TraverseFunc(func(tree *TreeNode) {
		tree.Print()
	})
	fmt.Println("------------------")
}

func (node *TreeNode) TraverseFunc(f func(*TreeNode)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
```

## Resource and Error Manage
* defer
	* Open/Close
	* Lock/Unlock
	* PrintHeader/PrintFooter

```
func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Finonacci()
	for i := 0; i < 30; i++ {
		fmt.Fprintln(writer, f())
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
```

## panic and recover
```
func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred: ", err)
		} else {
			panic(fmt.Sprintf("I don't know what to do: %v", r))
		}
	}()

	b := 0
	a := 5 / b
	fmt.Println(a)
	panic(errors.New("gun!"))
}
```

## goroutine and channel
No output about this program, because go use the 10 goroutine to execute,
but the main() function cannot wait the result , so execute the main() function, then exit,
So, no output from the console

* IO/select
* channel
* wait for lock
* runtime.Gosched()
* function call
```aidl
package main

import "fmt"

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	for i := 0; i < 10; i++ {
		go Add(i,i)
	}
}
```

* `channel` is the communicate model(message, not sharing memory) in the process,
not out of the process, if need the communication from the process to another
process, pls use the distributed system with `http` or `socket`, etc.
```aidl
var chanName chan ElementType

var ch chan int

var m map[string] chan bool

ch <- value  // write to channel ch

value := <- ch  // read from the channel ch

```

### Lock
* `sync.Mutex` and `sync.RWMutex`
* `sync.Once`

```aidl
package main

import (
	"sync"
	"fmt"
)

var a string

var once sync.Once

func setup() {
	a = "yzhao"
	fmt.Println("once:", a)
}



func doPrint(ch chan int) {
	ch <- 1
	once.Do(setup)
	println(a)
}

func twoPrint() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go doPrint(chs[i])
	}

	for _, ch := range(chs) {
		<- ch
	}
}

func main() {
	twoPrint()
}
```

Result:
```aidl
once: yzhao
yzhao
yzhao
yzhao
yzhao
yzhao
yzhao
yzhao
yzhao
yzhao
```