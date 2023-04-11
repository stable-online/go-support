package support

// base handler
type (
	// MMapFunction [K comparable, T any] map[K]T
	MMapFunction[K comparable, T any] func(b map[K]T) map[K]T

	// MFilterFunction map filter function
	MFilterFunction[K comparable, T any] func(b map[K]T) map[K]T

	// MReduceFunction map reduce function
	MReduceFunction[K comparable, T any] func(b map[K]T) any
)

// Mapper Mapper[Tan]
//
// @Description: Mapper type interface
type Mapper[K comparable, T any] interface {

	// Map
	//
	// @Description: map data
	// @param i
	// @return Operator
	Map(handler MMapFunction[K, T]) Mapper[K, T]

	// Filter
	//
	// @Description:
	// @param func(int, T) T
	// @return Mapper[T]
	Filter(handler MFilterFunction[K, T]) Mapper[K, T]

	// Reduce
	//
	// @Description:
	// @param func(k int, v T) (r bool)
	// @return Mapper[T]
	Reduce(handler MReduceFunction[K, T]) any

	// Get list
	//
	// @Description:
	// @return []T
	Get() map[K]T
}

// s Mapper type implement
type m[K comparable, T any] map[K]T

// Mapper[[]string] is implements Mapper interface ?
//
//	build mapper
var _ Mapper[string, string] = (m[string, string])(nil)

// NewM NewM[T map[any]any | []string]
//
// @Description: new Mapper
// @param i data
// @return Operator[T]
func NewM[K comparable, T any](i map[K]T) Mapper[K, T] {
	return m[K, T](i)
}

// Map
//
// @Description:  (fn func(b Mapper[T]) []T) used MMapF generate handle
// @receiver c
// @param i
// @return Operator
func (c m[K, T]) Map(fn MMapFunction[K, T]) Mapper[K, T] {
	return NewM(fn(c))
}

// MMapF  bb[T any, C any]
//
// @Description: map map handle
// @param fn
// @param initialize
// @return func(b Mapper[T]) any
func MMapF[K comparable, T any](fn func(key K, val T) (res T)) MMapFunction[K, T] {
	return func(i map[K]T) map[K]T {
		r := &MMapP[K, T]{fn: fn}
		return r.mapPF()(i)
	}
}

// MMapP MReduceP Reduce parameter build
// @Description:
type MMapP[K comparable, T any] struct {
	fn func(K, T) T
}

// mapPF [T any, C any]
//
// @Description: mapPF parameter function
// @param callback
// @param initialize
// @return func([]T) C
func (s *MMapP[K, T]) mapPF() func(map[K]T) map[K]T {
	return func(d map[K]T) map[K]T {
		//  build map
		ts := make(map[K]T, len(d))
		for k, v := range d {
			ts[k] = s.fn(k, v)
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
func (c m[K, T]) Filter(fn MFilterFunction[K, T]) Mapper[K, T] {
	return NewM(fn(c))
}

// MFilterF [T any, C any]
//
// @Description: map filter handle
// @param fn
// @param initialize
// @return func(b Mapper[T]) any
func MFilterF[K comparable, T any](fn func(K, T) bool) MFilterFunction[K, T] {
	return func(b map[K]T) map[K]T {
		r := &MFilterP[K, T]{fn: fn}
		return r.filterPF()(b)
	}
}

// MFilterP MReduceP  parameter build
// @Description:
type MFilterP[K comparable, T any] struct {
	fn func(K, T) bool
}

// filterPF [T any, C any]
//
// @Description: filterPF method of MFilterP
// @param callback
// @param initialize
// @return func([]T) C
func (s *MFilterP[K, T]) filterPF() func(d map[K]T) map[K]T {
	return func(d map[K]T) map[K]T {

		//  build map
		ts := make(map[K]T, len(d))

		for k, v := range d {
			if s.fn(k, v) {
				ts[k] = v
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
func (c m[K, T]) Reduce(fn MReduceFunction[K, T]) any {
	return fn(c)
}

// MReduceF
//
// @Description: map reduce handle
// @param fn
// @param initialize
// @return func(b Mapper[T]) any
func MReduceF[T any, K comparable, V any](fn func(carry V, key K, item T) V, initialize V) MReduceFunction[K, T] {
	return func(b map[K]T) any {
		r := MReduceP[K, T, V]{fn: fn, initialize: initialize}
		return r.reducePF()(b)
	}
}

// MReduceP  parameter build
// @Description:
type MReduceP[K comparable, T any, C any] struct {
	fn         func(carry C, key K, item T) C
	initialize C
}

// reducePF reducePF[T any, C any]
//
// @Description: Reduce parameter function
// @param callback
// @param initialize
// @return func([]T) C
func (s *MReduceP[K, T, C]) reducePF() func(map[K]T) C {
	return func(ts map[K]T) C {
		for k, v := range ts {
			s.initialize = s.fn(s.initialize, k, v)
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
func (c m[K, T]) Get() map[K]T {
	return c
}
