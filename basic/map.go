package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"id":   "1",
		"age":  "23",
		"name": "zyh",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3== nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traveling Map...")
	for k, v := range m {
		fmt.Printf("key=%s, value=%s\n", k, v)
	}

	fmt.Println("Get values")
	if name, ok := m["id"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("Key is not exists")
	}

	delete(m, "name")
	if name, ok := m["name"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("Key is not exist")
	}
}
