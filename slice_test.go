package support

import (
	"reflect"
	"testing"
)

// Test_collection_To
//
// @Description: test to method
// @param t
func Test_s_Get(t *testing.T) {

	type fields struct {
		data []any
	}

	tests := []struct {
		name   string
		fields fields
		want   []any
	}{
		{
			name:   "test_to_1",
			fields: fields{data: []any{"a", "b", "c"}},
			want:   []any{"a", "b", "c"},
		},

		{
			name:   "test_to_1",
			fields: fields{data: []any{"a1", "b2", "c1"}},
			want:   []any{"a1", "b2", "c1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewS(tt.fields.data)
			if got := c.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}

}

// Test_s_Map
//
// @Description:
// @param t
func Test_s_Map(t *testing.T) {

	type fields struct {
		data []string
	}

	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name:   "test_to_1",
			fields: fields{data: []string{"a", "b", "c"}},
			want:   []string{"a", "b", "c"},
		},

		{
			name:   "test_to_1",
			fields: fields{data: []string{"a1", "b2", "c1"}},
			want:   []string{"a1", "b2", "c1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := NewS(tt.fields.data).Map(SMapF(func(k int, v string) string {
				return v
			}))

			if got := c.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}

}

// Test_s_Filter
//
// @Description:
// @param t
func Test_s_Filter(t *testing.T) {
	type fields struct {
		data []string
	}

	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name:   "test_to_1",
			fields: fields{data: []string{"a", "b", "c"}},
			want:   []string{"a", "c"},
		},

		{
			name:   "test_to_1",
			fields: fields{data: []string{"a1", "b2", "c1"}},
			want:   []string{"b2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := NewS(tt.fields.data).Filter(SFilterF(func(k int, v string) bool {
				for _, i2 := range []string{"a", "c", "b2"} {
					if v == i2 {
						return true
					}
				}
				return false
			}))

			if got := c.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}

}

// Test_s_Reduce
//
// @Description:
// @param t
func Test_s_Reduce(t *testing.T) {
	type fields struct {
		data []string
	}
	type args struct {
		callback   func(carry string, item string) string
		initialize string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "test_2",

			fields: fields{
				data: []string{"a", "b", "c", "d"},
			},

			args: args{
				callback: func(carry string, item string) string {
					return carry + item
				},
				initialize: "",
			},

			want: "abcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewS(tt.fields.data).Reduce(SReduceF(tt.args.callback, tt.args.initialize)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test_s_Reduce
//
// @Description:
// @param t
func Test_s_Reduce_when_initializeIsSlice(t *testing.T) {
	type fields struct {
		data []string
	}
	type args struct {
		callback   func(carry []string, item string) []string
		initialize []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "test_2",

			fields: fields{
				data: []string{"a", "b", "c", "d"},
			},

			args: args{
				callback: func(carry []string, item string) []string {

					if item == "b" {
						return carry
					}
					return append(carry, item)

				},
				initialize: []string{},
			},

			want: []string{"a", "c", "d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewS(tt.fields.data).Reduce(SReduceF(tt.args.callback, tt.args.initialize)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestS_Filter
//
// @Description:
// @param t
// func TestS_Filter(t *testing.T) {
//
// 	fmt.Println(NewS([]string{"a", "b", "c"}).Map(SMapF(func(key int, item string) string {
// 		return item + "hello1"
// 	})).Filter(SFilterF(func(i int, t string) bool {
// 		return t != "ahello1"
// 	})).Reduce(SReduceF(func(carry string, item string) string {
// 		return carry + "=" + item
// 	}, "")))
//
// }

// TestMapH
//
// @Description:  test SMapF
// @param t
func TestSMapF(t *testing.T) {
	type args struct {
		data []string
		fn   func(int, string) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test map sf 1 ",
			args: args{
				data: []string{"a", "b", "c"},
				fn: func(i int, s2 string) string {
					return s2 + "-hello"
				},
			},

			want: []string{"a-hello", "b-hello", "c-hello"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SMapF(tt.args.fn)(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapSF() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestSFilterF
//
// @Description: test SFilterF sf
// @param t
func TestSFilterF(t *testing.T) {
	type args struct {
		data []string
		fn   func(int, string) bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test map sf 1 ",
			args: args{
				data: []string{"a", "b", "c"},
				fn: func(i int, s2 string) bool {
					return i != 1
				},
			},

			want: []string{"a", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SFilterF(tt.args.fn)(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterSF() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestReduceSF
//
// @Description: test SReduceF sf
// @param t
func TestSReduceF(t *testing.T) {

	type args struct {
		data       []string
		fn         func([]string, string) []string
		initialize []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test map sf 1 ",
			args: args{
				data: []string{"a", "b", "c"},
				fn: func(i []string, s2 string) []string {
					if s2 == "b" {
						return i
					}
					return append(i, s2)
				},
				initialize: []string{},
			},

			want: []string{"a", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SReduceF(tt.args.fn, tt.args.initialize)(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reduceSF() = %v, want %v", got, tt.want)
			}
		})
	}
}
