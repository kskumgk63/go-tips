package generatestr

import "testing"

func Test_generateRand(t *testing.T) {
	tests := []struct {
		name  string
		digit uint32
	}{
		// The probability of conflict is at least 1 in 25 million.
		{
			name:  "digit = 10",
			digit: 10,
		},
		// The probability of conflict is at least 1 in 25 million.
		{
			name:  "digit = 5",
			digit: 5,
		},
		// The probability of conflict is least 1 in 2 million.
		{
			name:  "digit = 3",
			digit: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 25_000_000; i++ {
				got1 := generateRand(tt.digit)
				got2 := generateRand(tt.digit)
				if got1 == got2 {
					t.Errorf("got same value: %s", got1)
				}
			}
		})
	}
}
