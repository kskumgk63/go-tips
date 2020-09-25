package error

import (
	"errors"
	"fmt"
	"testing"
)

func Test_err1_is(t *testing.T) {
	type fields struct {
		msg string
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "same",
			fields: fields{
				msg: "error 1",
			},
			args: args{
				err: err1{
					msg: "error 1",
				},
			},
			want: true,
		},
		{
			name: "diff because of wrap",
			fields: fields{
				msg: "error 1",
			},
			args: args{
				err: fmt.Errorf("wrap: %w", err1{
					msg: "error 1",
				}),
			},
			want: false,
		},
		{
			name: "diff because of use err2",
			fields: fields{
				msg: "error 1",
			},
			args: args{
				err: err2{
					msg: "error 1",
				},
			},
			want: false,
		},
		{
			name: "diff because of value",
			fields: fields{
				msg: "error 1",
			},
			args: args{
				err: err1{
					msg: "",
				},
			},
			want: false,
		},
		{
			name: "diff because of pointer",
			fields: fields{
				msg: "error 1",
			},
			args: args{
				err: &err1{
					msg: "error 1",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := err1{
				msg: tt.fields.msg,
			}
			if got := e.is(tt.args.err); got != tt.want {
				t.Logf("e = %T, tt.args.err = %T", e, tt.args.err)
				t.Errorf("err1.is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_err1_as(t *testing.T) {
	type fields struct {
		msg string
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "as",
			fields: fields{
				msg: "error 1",
			},
			args: args{
				err: err1{
					msg: "error 1",
				},
			},
			want: true,
		},
		{
			name: "as: different value",
			fields: fields{
				msg: "error 1",
			},
			args: args{
				err: err1{
					msg: "",
				},
			},
			want: true,
		},
		{
			name: "as: wrap",
			fields: fields{
				msg: "error 1",
			},
			args: args{
				err: fmt.Errorf("wrap: %w", err1{
					msg: "error 1",
				}),
			},
			want: true,
		},
		{
			name: "as: wrap*2",
			fields: fields{
				msg: "error 1",
			},
			args: args{
				err: fmt.Errorf("wrap: %w", fmt.Errorf("wrap: %w", err1{
					msg: "error 1",
				})),
			},
			want: true,
		},
		{
			name: "as: wrap, different value",
			fields: fields{
				msg: "error 1",
			},
			args: args{
				err: fmt.Errorf("wrap: %w", err1{
					msg: "",
				}),
			},
			want: true,
		},
		{
			name: "as: wrap*2, different value",
			fields: fields{
				msg: "error 1",
			},
			args: args{
				err: fmt.Errorf("wrap: %w", fmt.Errorf("wrap: %w", err1{
					msg: "",
				})),
			},
			want: true,
		},
		{
			name: "not as because of use pointer",
			fields: fields{
				msg: "error 1",
			},
			args: args{
				err: &err1{
					msg: "error 1",
				},
			},
			want: false,
		},
		{
			name: "not as because of different type",
			fields: fields{
				msg: "error 1",
			},
			args: args{
				err: err2{
					msg: "error 1",
				},
			},
			want: false,
		},
		{
			name: "not as because of different type",
			fields: fields{
				msg: "error 1",
			},
			args: args{
				err: errors.New("internal server error"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := err1{
				msg: tt.fields.msg,
			}
			if got := e.as(tt.args.err); got != tt.want {
				t.Logf("e = %T, tt.args.err = %T", e, tt.args.err)
				t.Errorf("err1.as() = %v, want %v", got, tt.want)
			}
		})
	}
}
