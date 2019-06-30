package main

import (
	"fmt"

	"github.com/syafdia/go-exercise/src/datastructure/linkedlist"
)

type Person struct {
	Name string
	Age  int
}

func tryLinkedList() {

}

func main() {
	xs := linkedlist.New("H", "E", "L")
	fmt.Println(xs)
}
