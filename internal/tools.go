package internal

// ReduceS ReduceS[T any]
//
// @Description:
// @param data
// @param fn
// @return []T
func ReduceS[T any, C any](callback func(carry C, item T) C, initialize C) func([]T) C {
	//
	return func(ts []T) C {
		//
		for _, v := range ts {
			initialize = callback(initialize, v)
		}
		//
		return initialize
	}

}

// MapS MapS[T any]
//
// @Description:
// @param data
// @param fn
// @return []T
func MapS[T any](data []T, fn func(int, T) T) []T {
	//  build slice
	ts := make([]T, 0, len(data))
	for k, v := range data {
		ts = append(ts, fn(k, v))
	}
	return ts
}

// FilterS [T any]
//
// @Description:
// @param data
// @param fn
// @return []T
func FilterS[T any](data []T, fn func(int, T) bool) []T {
	//  build slice
	ts := make([]T, 0, len(data))
	for k, v := range data {
		if fn(k, v) {
			ts = append(ts, v)
		}
	}
	return ts
}
