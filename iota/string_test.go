package iota

import "testing"

func Test_bloodType_String(t *testing.T) {
	tests := []struct {
		name string
		b    bloodType
		want string
	}{
		{
			b:    A,
			want: "A",
		},
		{
			b:    B,
			want: "B",
		},
		{
			b:    O,
			want: "O",
		},
		{
			b:    AB,
			want: "AB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.String(); got != tt.want {
				t.Errorf("bloodType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
