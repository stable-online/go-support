package support

import (
	"reflect"
	"testing"
)

// Test_collection_To
//
// @Description: test to method
// @param t
func Test_collection_To(t *testing.T) {

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
			if got := c.To(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("To() = %v, want %v", got, tt.want)
			}
		})
	}

}

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

			c := NewS(tt.fields.data).Map(func(k int, v string) string {
				return v
			})

			if got := c.To(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("To() = %v, want %v", got, tt.want)
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

			c := NewS(tt.fields.data).Filter(func(k int, v string) bool {
				for _, i2 := range []string{"a", "c", "b2"} {
					if v == i2 {
						return true
					}
				}
				return false
			})

			if got := c.To(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("To() = %v, want %v", got, tt.want)
			}
		})
	}

}
