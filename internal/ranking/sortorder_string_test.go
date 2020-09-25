package ranking

import "testing"

func Test_sortOrder_String(t *testing.T) {
	tests := []struct {
		name string
		i    sortOrder
		want string
	}{
		{
			name: "unknown",
			i:    0,
			want: "unknown",
		},
		{
			name: "ascending",
			i:    1,
			want: "ascending",
		},
		{
			name: "descending",
			i:    2,
			want: "descending",
		},
		{
			name: "negative",
			i:    -1,
			want: "sortOrder(-1)",
		},
		{
			name: "undefined",
			i:    42,
			want: "sortOrder(42)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.String(); got != tt.want {
				t.Errorf("sortOrder.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
