package generatestr

import "testing"

func Test_generateHash(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `given: "example", "test", "sample"`,
			args: args{
				args: []string{"example", "test", "sample"},
			},
			want: "01deaf2a837cf5e77fede720acb5b4448e36c401",
		},
		{
			name: `given: "password"`,
			args: args{
				args: []string{"password"},
			},
			want: "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8",
		},
		{
			name: `given: "password", "mail@example.com"`,
			args: args{
				args: []string{"password", "mail@example.com"},
			},
			want: "ecfbc6ba884c40dac7681553f3fbfba4f7a37fbb",
		},
		{
			name: `given: "1"`,
			args: args{
				args: []string{"1"},
			},
			want: "356a192b7913b04c54574d18c28d46e6395428ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateHash(tt.args.args...); got != tt.want {
				t.Errorf("generateHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
