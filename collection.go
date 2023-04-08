package support

// callback
type callback[T any] func(T, T) T

// Operator
//
// @Description: Operator interface
type Operator[T any] interface {
	// Map
	//
	// @Description: map data
	// @param i
	// @return Operator
	Map(callback[T]) Operator[T]

	// To
	//
	// @Description: to data
	// @param i
	// @return Operator
	To() T
}

// collection
// @Description:
type collection[T any] struct {
	data T
}

// Operator[T map[any]any | []string] is implements Operator interface?
var _ Operator[[]string] = (*collection[[]string])(nil)

// NewCollection NewCollection[T map[any]any | []string]
//
// @Description:
// @param i data
// @return Operator[T]
func NewCollection[T any](i T) Operator[T] {
	return &collection[T]{data: i}
}

// Map
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c *collection[T]) Map(fn callback[T]) Operator[T] {
	return NewCollection(c.data)
}

// To
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c *collection[T]) To() T {
	return c.data
}
