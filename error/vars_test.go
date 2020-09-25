package error

import (
	"errors"
	"testing"
)

func TestCompare(t *testing.T) {
	type args struct {
		err    error
		target error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "same",
			args: args{
				err:    ErrInternalServer,
				target: ErrInternalServer,
			},
			want: true,
		},
		{
			name: "same",
			args: args{
				err:    ErrNotFound,
				target: ErrNotFound,
			},
			want: true,
		},
		{
			name: "diff",
			args: args{
				err:    ErrInternalServer,
				target: ErrNotFound,
			},
			want: false,
		},
		{
			name: "diff",
			args: args{
				err:    ErrNotFound,
				target: ErrInternalServer,
			},
			want: false,
		},
		{
			name: "diff",
			args: args{
				err:    ErrInternalServer,
				target: errors.New("internal server error"),
			},
			want: false,
		},
		{
			name: "diff",
			args: args{
				err:    ErrNotFound,
				target: errors.New("not found error"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compare(tt.args.err, tt.args.target); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
