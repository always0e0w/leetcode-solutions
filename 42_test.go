package main

import "testing"

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			want: 6,
		},
		{
			name: "2",
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			want: 9,
		},
		{
			name: "3",
			args: args{
				height: []int{1},
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				height: []int{1, 2},
			},
			want: 0,
		},
		{
			name: "5",
			args: args{
				height: []int{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
