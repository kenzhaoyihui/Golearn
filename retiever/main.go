package main

import (
	"bufio"
	"fmt"
	mock "golearn/retiever/mock"
	real "golearn/retiever/real"
	"io"
	"strings"
	"time"
)

type Retiever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retiever) string {
	return r.Get("http://www.baidu.com")
}

func post(poster Poster) {
	poster.Post("http://www.baidu.com",
		map[string]string{
			"name":   "yzhao",
			"cource": "golang",
		})
}

const url = "http://www.baidu.com"

type RetieverPoster interface {
	Retiever
	Poster
	//Connect(host string)
}

func session(s RetieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "AAA",
	})

	return s.Get(url)
}

func inspect(r Retiever) {
	fmt.Printf("> (%T, %v)\n", r, r)
	switch v := r.(type) {
	case *mock.M:
		fmt.Println("Contents: ", v.Contents)
	case *real.R:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var r Retiever
	retiever := mock.M{"my name is yzhao"}

	r = &retiever
	//fmt.Println(download(r))
	//fmt.Printf("(%T, %v)\n", r, r)
	inspect(r)
	r = &real.R{
		UserAgent: "Chrome",
		Timeout:   time.Minute,
	}
	inspect(r)
	//fmt.Printf("(%T, %v)\n", r, r)
	//fmt.Println(download(r))

	//Type assertion
	if realRetiever, ok := r.(*real.R); ok {
		fmt.Println(realRetiever.Timeout)
	} else {
		fmt.Println("not a *realRetiever")
	}

	fmt.Println("Try a session...")

	// var s Retiever
	// s = mock.Retiever{"hello yzhao"}
	fmt.Println(session(&retiever))

	fmt.Println("Try session with mock retiever, because this retiever implement the Get and Post method")
	s := mock.M{"hello yzhao"}
	fmt.Println(&s)
	fmt.Println(session(&s))

	aa := `asd"dsd
	qwe
	asdasf
	
	23`

	printFileContents(strings.NewReader(aa))
}
