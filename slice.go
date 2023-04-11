package support

// base handler
type (
	// SMapFunction slice map function
	SMapFunction[T any] func(b []T) []T

	// SFilterFunction slice filter function
	SFilterFunction[T any] func(b []T) []T

	// SReduceFunction slice reduce function
	SReduceFunction[T any] func(b []T) any
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
	Map(handler SMapFunction[T]) Slicer[T]

	// Filter
	//
	// @Description:
	// @param func(int, T) T
	// @return Slicer[T]
	Filter(handler SFilterFunction[T]) Slicer[T]

	// Reduce
	//
	// @Description:
	// @param func(k int, v T) (r bool)
	// @return Slicer[T]
	Reduce(handler SReduceFunction[T]) any

	// Get list
	//
	// @Description:
	// @return []T
	Get() []T
}

// s Slicer type implement
type s[T any] []T

// Slicer[[]string] is implements Slicer interface ?
//
//	build slice
var _ Slicer[string] = (s[string])(nil)

// NewS NewS[T map[any]any | []string]
//
// @Description: new Slicer
// @param i data
// @return Operator[T]
func NewS[T any](i []T) Slicer[T] {
	return s[T](i)
}

// Map
//
// @Description:  (fn func(b Slicer[T]) []T) used SMapF generate handle
// @receiver c
// @param i
// @return Operator
func (c s[T]) Map(fn SMapFunction[T]) Slicer[T] {
	return NewS(fn(c))
}

// SMapF  bb[T any, C any]
//
// @Description: slice map handle
// @param fn
// @param initialize
// @return func(b Slicer[T]) any
func SMapF[T any](fn func(int, T) T) SMapFunction[T] {
	return func(i []T) []T {
		r := &SmapP[T]{fn: fn}
		return r.mapPF()(i)
	}
}

// SmapP SReduceP Reduce parameter build
// @Description:
type SmapP[T any] struct {
	fn func(int, T) T
}

// mapPF [T any, C any]
//
// @Description: mapPF parameter function
// @param callback
// @param initialize
// @return func([]T) C
func (s *SmapP[T]) mapPF() func([]T) []T {
	return func(d []T) []T {
		//  build slice
		ts := make([]T, 0, len(d))
		for k, v := range d {
			ts = append(ts, s.fn(k, v))
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
func (c s[T]) Filter(fn SFilterFunction[T]) Slicer[T] {
	return NewS(fn(c))
}

// SFilterF [T any, C any]
//
// @Description: slice filter handle
// @param fn
// @param initialize
// @return func(b Slicer[T]) any
func SFilterF[T any](fn func(int, T) bool) SFilterFunction[T] {
	return func(b []T) []T {
		r := &SFilterP[T]{fn: fn}
		return r.filterPF()(b)
	}
}

// SFilterP SReduceP  parameter build
// @Description:
type SFilterP[T any] struct {
	fn func(int, T) bool
}

// filterPF [T any, C any]
//
// @Description: filterPF method of SFilterP
// @param callback
// @param initialize
// @return func([]T) C
func (s *SFilterP[T]) filterPF() func([]T) []T {
	return func(d []T) []T {
		//  build slice
		ts := make([]T, 0, len(d))
		for k, v := range d {
			if s.fn(k, v) {
				ts = append(ts, v)
			}
		}
		return ts
	}
}

// Reduce
//
// @Description: Reduce Of s
// @receiver c
// @param i
// @return Operator
func (c s[T]) Reduce(fn SReduceFunction[T]) any {
	return fn(c)
}

// SReduceF
//
// @Description: slice reduce handle
// @param fn
// @param initialize
// @return func(b Slicer[T]) any
func SReduceF[T any, C any](fn func(carry C, item T) C, initialize C) SReduceFunction[T] {
	return func(b []T) any {
		r := SReduceP[T, C]{fn: fn, initialize: initialize}
		return r.reducePF()(b)
	}
}

// SReduceP  parameter build
// @Description:
type SReduceP[T any, C any] struct {
	fn         func(carry C, item T) C
	initialize C
}

// reducePF reducePF[T any, C any]
//
// @Description: Reduce parameter function
// @param callback
// @param initialize
// @return func([]T) C
func (s *SReduceP[T, C]) reducePF() func([]T) C {
	return func(ts []T) C {
		for _, v := range ts {
			s.initialize = s.fn(s.initialize, v)
		}
		return s.initialize
	}
}

// Get To
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c s[T]) Get() []T {
	return c
}
