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

	// To Splicer
	//
	// @Description:
	// @return []T
	To() []T
}

// s
// @Description: splice type
type s[K any, T any] struct {
	data []T
}

// Operator[[]string] is implements Operator interface ?
//
//	build slice
var _ Splicer[[]any] = (*s[int, []any])(nil)

// NewS NewS[T map[any]any | []string]
//
// @Description: new splice
// @param i data
// @return Operator[T]
func NewS[T any](i []T) Splicer[T] {
	return &s[int, T]{data: i}
}

// Map
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c *s[K, T]) Map(fn func(int, T) T) Splicer[T] {
	return NewS(internal.MapS[T](c.data, fn))
}

// Filter Map
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c *s[K, T]) Filter(fn func(int, T) bool) Splicer[T] {
	return NewS(internal.FilterS[T](c.data, fn))
}

// To
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c *s[K, T]) To() []T {
	return c.data
}
