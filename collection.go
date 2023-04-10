package support

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
	Reduce(fn func(carry any, item T) any, initialize any) any

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
	return &s[T]{data: MapSF[T](c.data, fn)}
}

// MapSF MapSF[T any]
//
// @Description: map splice function
// @param data
// @param fn
// @return []T
func MapSF[T any](data []T, fn func(int, T) T) []T {
	//  build slice
	ts := make([]T, 0, len(data))
	for k, v := range data {
		ts = append(ts, fn(k, v))
	}
	return ts
}

// Filter Map
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c *s[T]) Filter(fn func(k int, v T) (r bool)) Splicer[T] {
	return &s[T]{data: FilterSF[T](c.data, fn)}
}

// FilterSF [T any]
//
// @Description:
// @param data
// @param fn
// @return []T
func FilterSF[T any](data []T, fn func(int, T) bool) []T {
	//  build slice
	ts := make([]T, 0, len(data))
	for k, v := range data {
		if fn(k, v) {
			ts = append(ts, v)
		}
	}
	return ts
}

// Reduce Map
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c *s[T]) Reduce(fn func(carry any, item T) any, initialize any) any {
	return NewReduceHandler(fn, initialize)(c.data)
}

// NewReduceHandler NewReduceHandler[T any, C any]
//
// @Description:
// @param callback
// @param initialize
// @return func([]T) C
func NewReduceHandler[T any](callback func(carry any, item T) any, initialize any) func([]T) any {
	//
	return func(ts []T) any {
		//
		for _, v := range ts {
			initialize = callback(initialize, v)
		}
		//
		return initialize
	}

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
