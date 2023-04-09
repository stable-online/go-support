package support

import (
	"github.com/stable-online/support/internal"
)

// Splicer Splicer[Tan]
//
// @Description: splice type interface
type Splicer[T any] interface {
	// Map
	//
	// @Description: map data
	// @param i
	// @return Operator
	Map(func(int, T) T) Splicer[T]

	// Filter
	//
	// @Description:
	// @param func(int, T) T
	// @return Splicer[T]
	Filter(func(int, T) bool) Splicer[T]

	// Reduce
	//
	// @Description:
	// @param func(k int, v T) (r bool)
	// @return Splicer[T]
	Reduce(callback func(carry any, item T) (res any), initialize any, data []T) any

	// To Splicer
	//
	// @Description:
	// @return []T
	To() []T
}

// s
// @Description: splice type
type s[T any] struct {
	data []T
}

// Operator[[]string] is implements Operator interface ?
//
//	build slice
var _ Splicer[[]any] = (*s[[]any])(nil)

// NewS NewS[T map[any]any | []string]
//
// @Description: new splice
// @param i data
// @return Operator[T]
func NewS[T any](i []T) Splicer[T] {
	return &s[T]{data: i}
}

// Map
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c *s[T]) Map(fn func(int, T) T) Splicer[T] {
	return &s[T]{data: internal.MapS[T](c.data, fn)}
}

// Filter Map
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c *s[T]) Filter(fn func(k int, v T) (r bool)) Splicer[T] {
	return &s[T]{data: internal.FilterS[T](c.data, fn)}
}

// Reduce Map
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c *s[T]) Reduce(callback func(carry any, item T) (res any), initialize any, data []T) any {
	return internal.ReduceS(callback, initialize)(data)
}

// To
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c *s[T]) To() []T {
	return c.data
}
