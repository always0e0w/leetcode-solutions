package main

import "testing"

func Test_nthSuperUglyNumber(t *testing.T) {
	type args struct {
		n      int
		primes []int
	}
	tests := []struct {
		name     string
		args     args
		wantUgly int
	}{
		{
			name: "1",
			args: args{
				n:      12,
				primes: []int{2, 7, 13, 19},
			},
			wantUgly: 32,
		},
		{
			name: "2",
			args: args{
				n:      1,
				primes: []int{2, 3, 5},
			},
			wantUgly: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUgly := nthSuperUglyNumber(tt.args.n, tt.args.primes); gotUgly != tt.wantUgly {
				t.Errorf("nthSuperUglyNumber() = %v, want %v", gotUgly, tt.wantUgly)
			}
		})
	}
}
