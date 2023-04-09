package support

type Mapper[K int | string, V any] map[K]V

// Operator
//
// @Description: Operator interface
type Operator[T any] interface {
	// Map
	//
	// @Description: map data
	// @param i
	// @return Operator
	Map(func(int, T) T) Splice[T]
}

// Splice Splice[Tan]
//
// @Description: splice type interface
type Splice[T any] interface {
	Operator[T]
	To() []T
}

// s
// @Description: splice type
type s[T any] struct {
	data []T
}

// Operator[[]string] is implements Operator interface ?
var _ Splice[[]any] = (*s[[]any])(nil)

// OfS OfS[T map[any]any | []string]
//
// @Description: of splice
// @param i data
// @return Operator[T]
func OfS[T any](i []T) Splice[T] {
	return &s[T]{data: i}
}

// Map
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c *s[T]) Map(fn func(int, T) T) Splice[T] {

	ts := make([]T, 0, len(c.data))
	for k, v := range c.data {
		ts = append(ts, fn(k, v))
	}

	// Maps[int64, T](func(i int, a string) string {
	// 	return a
	// })(c.data)

	return OfS(ts)
}

// func Maps[K any, T any](callback func(K, T) T) func([]T) []T {
//
// 	return func(xs []T) []T {
//
// 		result := make([]T, 0, len(xs))
//
// 		var k K
// 		var v T
// 		for k, v = range xs {
// 			result = append(result, callback(k, v))
// 		}
//
// 		return result
// 	}
// }

// To
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c *s[T]) To() []T {
	return c.data
}
