package example

import (
	"fmt"
	"time"
)

func MultiplyTwo(v int) int {
	time.Sleep(2 * time.Second)
	return v * 2
}

func Square(v int) int {
	time.Sleep(2 * time.Second)
	return v * v
}

func AddQuoute(v int) string {
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("'%d'", v)
}

func AddFoo(v string) string {
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("Foo %v", v)
}
