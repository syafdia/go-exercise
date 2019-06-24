package pipeline

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkWithPipelineModule(b *testing.B) {
	outC := New(func(inC chan interface{}) {
		defer close(inC)
		for i := 0; i < b.N; i++ {
			inC <- i
		}
	}).
		Pipe(func(in interface{}) (interface{}, error) {
			return multiplyTwo(in.(int)), nil
		}).
		Pipe(func(in interface{}) (interface{}, error) {
			return square(in.(int)), nil
		}).
		Pipe(func(in interface{}) (interface{}, error) {
			return addQuoute(in.(int)), nil
		}).
		Pipe(func(in interface{}) (interface{}, error) {
			return addFoo(in.(string)), nil
		}).
		Merge()

	for range outC {
		// Do nothing, just for  drain out channel
	}
}

func BenchmarkWithoutPipelineModule(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addFoo(addQuoute(square(multiplyTwo(i))))
	}
}

func multiplyTwo(v int) int {
	time.Sleep(100 * time.Millisecond)
	return v * 2
}

func square(v int) int {
	time.Sleep(200 * time.Millisecond)
	return v * v
}

func addQuoute(v int) string {
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("'%d'", v)
}

func addFoo(v string) string {
	time.Sleep(200 * time.Millisecond)
	return fmt.Sprintf("%s - Foo", v)
}
