package ch

import (
	"reflect"
	"testing"
)

func Test_pipelineMultiply(t *testing.T) {
	var (
		doneCh = func() doneCh {
			return make(doneCh)
		}
	)
	type args struct {
		intStream   intStream
		multipliers [3]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				intStream:   newIntStream(doneCh(), 1, 2, 3),
				multipliers: [3]int{1, 2, 3},
			},
			want: []int{6, 12, 18},
		},
		{
			args: args{
				intStream:   newIntStream(doneCh(), 2, 2, 2),
				multipliers: [3]int{2, 2, 2},
			},
			want: []int{16, 16, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pipelineMultiply(tt.args.intStream, tt.args.multipliers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pipelineMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
