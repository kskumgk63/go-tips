package generatestr

import "testing"

func Test_hash(t *testing.T) {
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
			want: "c44dbfd454ddd80dbad79b99a03b7fa04762cb89a4f8a87f2b5c8931aa34b96d",
		},
		{
			name: `given: "password"`,
			args: args{
				args: []string{"password"},
			},
			want: "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8",
		},
		{
			name: `given: "password", "mail@example.com"`,
			args: args{
				args: []string{"password", "mail@example.com"},
			},
			want: "673da4c959b1796a1e8a1e9d32dd1b358d45b2523141486c19fac2397d144e35",
		},
		{
			name: `given: "1"`,
			args: args{
				args: []string{"1"},
			},
			want: "6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash(tt.args.args...); got != tt.want {
				t.Errorf("generateHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
