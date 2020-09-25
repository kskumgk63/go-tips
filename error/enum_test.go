package error

import (
	"errors"
	"fmt"
	"testing"

	"golang.org/x/xerrors"
)

func TestCode_is(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		code Code
		args args
		want bool
	}{
		{
			name: "same",
			code: Zero,
			args: args{
				err: Zero,
			},
			want: true,
		},
		{
			name: "same",
			code: One,
			args: args{
				err: One,
			},
			want: true,
		},
		{
			name: "diff: wrap",
			code: One,
			args: args{
				err: fmt.Errorf("wrap: %w", One),
			},
			want: false,
		},
		{
			name: "diff: wrap",
			code: One,
			args: args{
				err: xerrors.Errorf("wrap: %w", One),
			},
			want: false,
		},
		{
			name: "diff",
			code: Zero,
			args: args{
				err: One,
			},
			want: false,
		},
		{
			name: "diff",
			code: One,
			args: args{
				err: Zero,
			},
			want: false,
		},
		{
			name: "diff",
			code: One,
			args: args{
				err: errors.New("internal server error"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.code.is(tt.args.err); got != tt.want {
				t.Errorf("Code.is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCode_as(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		code Code
		args args
		want bool
	}{
		{
			name: "as",
			code: Zero,
			args: args{
				err: Zero,
			},
			want: true,
		},
		{
			name: "as",
			code: Zero,
			args: args{
				err: One,
			},
			want: true,
		},
		{
			name: "as: wrap",
			code: Zero,
			args: args{
				err: fmt.Errorf("wrap: %w", Zero),
			},
			want: true,
		},
		{
			name: "as: wrap",
			code: Zero,
			args: args{
				err: fmt.Errorf("wrap: %w", One),
			},
			want: true,
		},
		{
			name: "as: wrap*2",
			code: Zero,
			args: args{
				err: fmt.Errorf("wrap2: %w", fmt.Errorf("wrap: %w", One)),
			},
			want: true,
		},
		{
			name: "diff",
			code: Zero,
			args: args{
				err: errors.New("internal server error"),
			},
			want: false,
		},
		{
			name: "diff",
			code: Zero,
			args: args{
				err: errors.New("Zero"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.code.as(tt.args.err); got != tt.want {
				t.Errorf("Code.as() = %v, want %v", got, tt.want)
			}
		})
	}
}
