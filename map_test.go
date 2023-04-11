package support

import (
	"reflect"
	"testing"
)

// Test_m_Get
//
// @Description:
// @param t
func Test_m_Get(t *testing.T) {

	type fields struct {
		data map[string]string
	}

	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			name:   "test_to_1",
			fields: fields{data: map[string]string{"a": "1", "b": "2", "c": "2"}},
			want:   map[string]string{"a": "1", "b": "2", "c": "2"},
		},

		{
			name:   "test_to_1",
			fields: fields{data: map[string]string{"a": "1", "b": "2", "c": "2"}},
			want:   map[string]string{"a": "1", "b": "2", "c": "2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewM(tt.fields.data)
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
func Test_m_Map(t *testing.T) {
	type fields struct {
		data map[string]string
	}

	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			name:   "test_to_1",
			fields: fields{data: map[string]string{"a": "1", "b": "2", "c": "2"}},
			want:   map[string]string{"a": "a1", "b": "b2", "c": "c2"},
		},

		{
			name:   "test_to_1",
			fields: fields{data: map[string]string{"a": "1", "b": "2", "c": "2"}},
			want:   map[string]string{"a": "a1", "b": "b2", "c": "c2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewM(tt.fields.data)
			if got := c.Map(MMapF(func(k string, t string) string {
				return k + t
			})).Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test_s_Filter
//
// @Description:
// @param t
func Test_m_Filter(t *testing.T) {
	type fields struct {
		data map[string]string
	}

	tests := []struct {
		name      string
		fields    fields
		want      map[string]string
		ignoreKey map[string]string
	}{
		{
			name:      "test_to_1",
			fields:    fields{data: map[string]string{"a": "1", "b": "2", "c": "2"}},
			ignoreKey: map[string]string{"b": ""},
			want:      map[string]string{"a": "1", "c": "2"},
		},

		{
			name:      "test_to_1",
			fields:    fields{data: map[string]string{"a": "1", "b": "2", "c": "2"}},
			ignoreKey: map[string]string{"c": ""},
			want:      map[string]string{"a": "1", "b": "2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewM(tt.fields.data)
			if got := c.Filter(MFilterF(func(k string, t string) bool {

				_, ok := tt.ignoreKey[k]

				return ok == false

			})).Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test_s_Filter
//
// @Description:
// @param t
func Test_m_Reduce(t *testing.T) {
	type fields struct {
		data map[string]string
	}

	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			name:   "test_to_1",
			fields: fields{data: map[string]string{"a": "1", "b": "2", "c": "2"}},
			want:   map[string]string{"a": "a1", "b": "b2", "c": "c2"},
		},

		{
			name:   "test_to_1",
			fields: fields{data: map[string]string{"a": "1", "b": "2", "c": "3"}},
			want:   map[string]string{"a": "a1", "b": "b2", "c": "c3"},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			c := NewM(tt.fields.data)

			if got := c.Reduce(MReduceF(func(carry map[string]string, k string, item string) map[string]string {

				carry[k] = k + item
				return carry

			}, map[string]string{})); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMMapF(t *testing.T) {
	type fields struct {
		data map[string]string
	}

	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			name:   "test_to_1",
			fields: fields{data: map[string]string{"a": "1", "b": "2", "c": "2"}},
			want:   map[string]string{"a": "a1", "b": "b2", "c": "c2"},
		},

		{
			name:   "test_to_1",
			fields: fields{data: map[string]string{"a": "1", "b": "2", "c": "2"}},
			want:   map[string]string{"a": "a1", "b": "b2", "c": "c2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MMapF(func(k string, t string) string {
				return k + t
			})(tt.fields.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMFilterF(t *testing.T) {
	type fields struct {
		data map[string]string
	}

	tests := []struct {
		name      string
		fields    fields
		want      map[string]string
		ignoreKey map[string]string
	}{
		{
			name:      "test_to_1",
			fields:    fields{data: map[string]string{"a": "1", "b": "2", "c": "2"}},
			ignoreKey: map[string]string{"b": ""},
			want:      map[string]string{"a": "1", "c": "2"},
		},

		{
			name:      "test_to_1",
			fields:    fields{data: map[string]string{"a": "1", "b": "2", "c": "2"}},
			ignoreKey: map[string]string{"c": ""},
			want:      map[string]string{"a": "1", "b": "2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MFilterF(func(k string, t string) bool {

				_, ok := tt.ignoreKey[k]
				return ok == false

			})(tt.fields.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMReduceF(t *testing.T) {

	type fields struct {
		data map[string]string
	}

	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			name:   "test_to_1",
			fields: fields{data: map[string]string{"a": "1", "b": "2", "c": "2"}},
			want:   map[string]string{"a": "a1", "b": "b2", "c": "c2"},
		},

		{
			name:   "test_to_1",
			fields: fields{data: map[string]string{"a": "1", "b": "2", "c": "3"}},
			want:   map[string]string{"a": "a1", "b": "b2", "c": "c3"},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if got := MReduceF(func(carry map[string]string, k string, item string) map[string]string {

				carry[k] = k + item
				return carry

			}, map[string]string{})(tt.fields.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
