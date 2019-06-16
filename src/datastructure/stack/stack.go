package stack

import "github.com/syafdia/go-exercise/src/types"

type Stack interface {
	Pop() types.T
	Push(v types.T)
}
