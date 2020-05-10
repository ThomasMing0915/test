package main

import "testing"

func Test_exchangeBits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{
				num: 0,
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				num: 2,
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				num: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exchangeBits(tt.args.num); got != tt.want {
				t.Errorf("exchangeBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
