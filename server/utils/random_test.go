package utils

import "testing"

func TestGenerateRandomString(t *testing.T) {
	type args struct {
		n int
	}

	cases := []struct {
		name       string
		args       args
		wantLength int
	}{
		{"0", args{0}, 0},
		{"1", args{1}, 1},
		{"10", args{10}, 10},
		{"100", args{100}, 100},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomString(tt.args.n); len(got) != tt.wantLength {
				t.Errorf("GenerateRandomString() = %v, wantLength %v", got, tt.wantLength)
			}
		})
	}
}
