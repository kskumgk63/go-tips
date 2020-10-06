package ch

import (
	"reflect"
	"testing"
	"time"
)

func Test_getFastestDuration(t *testing.T) {
	type args struct {
		duration [5]time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "100ms is the fastest",
			args: args{
				duration: [5]time.Duration{
					1 * time.Hour,
					1 * time.Minute,
					10 * time.Second,
					1 * time.Second,
					100 * time.Millisecond,
				},
			},
			want: 100 * time.Millisecond,
		},
		{
			name: "10ms is the fastest",
			args: args{
				duration: [5]time.Duration{
					1 * time.Hour,
					10 * time.Millisecond,
					1 * time.Minute,
					10 * time.Second,
					1 * time.Second,
				},
			},
			want: 10 * time.Millisecond,
		},
		{
			name: "lack of args",
			args: args{
				duration: [5]time.Duration{
					1 * time.Second,
					100 * time.Millisecond,
				},
			},
			want: time.Duration(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFastestDuration(tt.args.duration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFastestDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
