package support

// base handler
type (
	// map handler
	mapHandler[T any] func(b Slicer[T]) []T

	// filter handler
	filterHandler[T any] func(b Slicer[T]) []T

	// reduce handler
	reduceHandler[T any] func(b Slicer[T]) any
)

// Slicer Slicer[Tan]
//
// @Description: Slicer type interface
type Slicer[T any] interface {

	// Map
	//
	// @Description: map data
	// @param i
	// @return Operator
	Map(handler mapHandler[T]) Slicer[T]

	// Filter
	//
	// @Description:
	// @param func(int, T) T
	// @return Slicer[T]
	Filter(handler filterHandler[T]) Slicer[T]

	// Reduce
	//
	// @Description:
	// @param func(k int, v T) (r bool)
	// @return Slicer[T]
	Reduce(handler reduceHandler[T]) any

	// Get list
	//
	// @Description:
	// @return []T
	Get() []T
}

// s
// @Description: Slicer type
type s[T any] struct {
	data []T
}

// Operator[[]string] is implements Operator interface ?
//
//	build slice
var _ Slicer[[]any] = (*s[[]any])(nil)

// NewS NewS[T map[any]any | []string]
//
// @Description: new Slicer
// @param i data
// @return Operator[T]
func NewS[T any](i []T) Slicer[T] {
	return &s[T]{data: i}
}

// Map
//
// @Description:  (fn func(b Slicer[T]) []T) used MapH generate handle
// @receiver c
// @param i
// @return Operator
func (c *s[T]) Map(fn mapHandler[T]) Slicer[T] {
	return NewS(fn(c))
}

// MapH  bb[T any, C any]
//
// @Description: reduce handle of reduceP
// @param fn
// @param initialize
// @return func(b Slicer[T]) any
func MapH[T any](fn func(int, T) T) func(b Slicer[T]) []T {
	return func(b Slicer[T]) []T {
		r := &mapP[T, int]{fn: fn}
		return r.mapPF()(b.Get())
	}
}

// reduceP Reduce parameter build
// @Description:
type mapP[T any, C any] struct {
	fn func(int, T) T
}

// mapPF reducePF ReduceSF[T any, C any]
//
// @Description: Reduce parameter function
// @param callback
// @param initialize
// @return func([]T) C
func (s *mapP[T, C]) mapPF() func([]T) []T {
	return func(ts []T) []T {
		return MapSF(s.fn)(ts)
	}
}

// MapSF MapSF[T any]
//
// @Description: Map Slicer function
// @param data
// @param fn
// @return []T
func MapSF[T any](fn func(int, T) T) func(d []T) []T {
	return func(d []T) []T {
		//  build slice
		ts := make([]T, 0, len(d))
		for k, v := range d {
			ts = append(ts, fn(k, v))
		}
		return ts
	}

}

// Filter Map
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c *s[T]) Filter(fn filterHandler[T]) Slicer[T] {
	return NewS(fn(c))
}

// FilterH MapH  bb[T any, C any]
//
// @Description: reduce handle of reduceP
// @param fn
// @param initialize
// @return func(b Slicer[T]) any
func FilterH[T any](fn func(int, T) bool) func(b Slicer[T]) []T {
	return func(b Slicer[T]) []T {
		r := &filterP[T, int]{fn: fn}
		return r.filterPF()(b.Get())
	}
}

// reduceP Reduce parameter build
// @Description:
type filterP[T any, C any] struct {
	fn func(int, T) bool
}

// mapPF reducePF ReduceSF[T any, C any]
//
// @Description: Reduce parameter function
// @param callback
// @param initialize
// @return func([]T) C
func (s *filterP[T, C]) filterPF() func([]T) []T {
	return func(ts []T) []T {
		return FilterSF(s.fn)(ts)
	}
}

// FilterSF [T any]
//
// @Description: filter slice function
// @param data
// @param fn
// @return []T
func FilterSF[T any](fn func(int, T) bool) func(d []T) []T {
	//  build slice
	return func(d []T) []T {
		//  build slice
		ts := make([]T, 0, len(d))
		for k, v := range d {
			if fn(k, v) {
				ts = append(ts, v)
			}
		}
		return ts
	}

}

// Reduce Map
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c *s[T]) Reduce(fn reduceHandler[T]) any {
	return fn(c)
}

// ReduceH bb[T any, C any]
//
// @Description: reduce handle of reduceP
// @param fn
// @param initialize
// @return func(b Slicer[T]) any
func ReduceH[T any, C any](fn func(carry C, item T) C, initialize C) func(b Slicer[T]) any {
	return func(b Slicer[T]) any {
		r := reduceP[T, C]{fn: fn, initialize: initialize}
		return r.reducePF()(b.Get())
	}
}

// reduceP Reduce parameter build
// @Description:
type reduceP[T any, C any] struct {
	fn         func(carry C, item T) C
	initialize C
}

// reducePF ReduceSF[T any, C any]
//
// @Description: Reduce parameter function
// @param callback
// @param initialize
// @return func([]T) C
func (s *reduceP[T, C]) reducePF() func([]T) C {
	return ReduceSF[T, C](s.fn, s.initialize)
}

// ReduceSF ReduceSF[T any, C any]
//
// @Description:
// @param callback
// @param initialize
// @return func([]T) C
func ReduceSF[T any, C any](callback func(carry C, item T) C, initialize C) func([]T) C {
	return func(ts []T) C {
		for _, v := range ts {
			initialize = callback(initialize, v)
		}
		return initialize
	}
}

// Get To
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c *s[T]) Get() []T {
	return c.data
}
