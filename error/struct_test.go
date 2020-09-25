package error

import (
	"errors"
	"fmt"
	"testing"
)

func Test_stringValueError_is(t *testing.T) {
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
				err: stringValueError{
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
				err: fmt.Errorf("wrap: %w", stringValueError{
					msg: "error 1",
				}),
			},
			want: false,
		},
		{
			name: "diff because of use structValueError",
			fields: fields{
				msg: "error 1",
			},
			args: args{
				err: structValueError{
					code:  code{code: 1},
					value: value{str: "internal server error"},
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
				err: stringValueError{
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
				err: &stringValueError{
					msg: "error 1",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := stringValueError{
				msg: tt.fields.msg,
			}
			if got := e.is(tt.args.err); got != tt.want {
				t.Logf("e = %T, tt.args.err = %T", e, tt.args.err)
				t.Errorf("stringValueError.is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringValueError_as(t *testing.T) {
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
				err: stringValueError{
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
				err: stringValueError{
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
				err: fmt.Errorf("wrap: %w", stringValueError{
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
				err: fmt.Errorf("wrap: %w", fmt.Errorf("wrap: %w", stringValueError{
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
				err: fmt.Errorf("wrap: %w", stringValueError{
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
				err: fmt.Errorf("wrap: %w", fmt.Errorf("wrap: %w", stringValueError{
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
				err: &stringValueError{
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
				err: structValueError{
					code:  code{code: 1},
					value: value{str: "internal server error"},
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
			e := stringValueError{
				msg: tt.fields.msg,
			}
			if got := e.as(tt.args.err); got != tt.want {
				t.Logf("e = %T, tt.args.err = %T", e, tt.args.err)
				t.Errorf("stringValueError.as() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_interfaceValueError_is(t *testing.T) {
	type fields struct {
		err error
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
			name: "absolutely false because of error interface is not comparable",
			fields: fields{
				err: errors.New("hoge"),
			},
			args: args{
				err: errors.New("hoge"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := interfaceValueError{
				err: tt.fields.err,
			}
			if got := e.is(tt.args.err); got != tt.want {
				t.Errorf("interfaceValueError.is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structValueError_is(t *testing.T) {
	type fields struct {
		code  code
		value value
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
			name: "same because of strcut is comparable",
			fields: fields{
				code:  code{code: 1},
				value: value{str: "internal server error"},
			},
			args: args{
				err: structValueError{
					code:  code{code: 1},
					value: value{str: "internal server error"},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := structValueError{
				code:  tt.fields.code,
				value: tt.fields.value,
			}
			if got := e.is(tt.args.err); got != tt.want {
				t.Errorf("structValueError.is() = %v, want %v", got, tt.want)
			}
		})
	}
}
