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
	// var head T = "HELLO"
	xs := linkedlist.New("H", "E", "L")
	fmt.Println(xs)

	// xs.Map(func(v T) U { return strings.ToLower(v.(string)) })

	// fmt.Println(xs)
	// fmt.Println(xs.Size())
	// fmt.Println(
	// 	xs.Map(func(v T) U {
	// 		return strings.ToLower(v.(string))
	// 	}).Filter(func(v T) bool {
	// 		return v.(string) != "l"
	// 	}))

	// people := NewLinkedList(
	// 	Person{"Badi", 10},
	// 	Person{"Bidi", 11},
	// 	Person{"Budi", 12},
	// 	Person{"Bedi", 13},
	// 	Person{"Bodi", 14},
	// )

	// fmt.Println(people.Filter(func(v T) bool {
	// 	person := v.(Person)
	// 	return person.Age > 13
	// }))

	// fmt.Println(people.Reduce(">", func(acc U, v T) U {
	// 	person := v.(Person)
	// 	return acc.(string) + " " + person.Name
	// }))

}

func main() {
	tryLinkedList()
}
