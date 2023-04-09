package internal

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
