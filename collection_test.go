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
			c := NewCollection(tt.fields.data)
			if got := c.To(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("To() = %v, want %v", got, tt.want)
			}
		})
	}

}
