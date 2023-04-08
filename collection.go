package support

// callback
type callback func(any, any) any

// Operator
//
// @Description: Operator interface
type Operator[T map[any]any | []string] interface {
	// Map
	//
	// @Description: map data
	// @param i
	// @return Operator
	Map(callback) Operator[T]

	// To
	//
	// @Description: to data
	// @param i
	// @return Operator
	To() T
}

// collection
// @Description:
type collection[T map[any]any | []string] struct {
	data T
}

// Operator[T map[any]any | []string] is implements Operator interface?
var _ Operator[[]string] = (*collection[[]string])(nil)

// NewCollection NewCollection[T map[any]any | []string]
//
// @Description:
// @param i data
// @return Operator[T]
func NewCollection[T map[any]any | []string](i T) Operator[T] {
	return &collection[T]{data: i}
}

// Map
//
// @Description:
// @receiver c
// @param i
// @return Operator
func (c *collection[T]) Map(fn callback) Operator[T] {
	return &collection[T]{data: c.data}
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
