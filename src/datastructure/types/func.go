package types

type Func0 func(v T)

type Func1 func(v T) U

type Func2 func(v1 T, v2 U) V

type Mapper func(v T) U

type Filterrer func(v T) bool

type Reducer func(acc U, v T) U
