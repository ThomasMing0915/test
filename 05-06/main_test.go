package main

import "testing"

func Test_convertInteger(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1,1",
			args: args{
				A: 1,
				B: 1,
			},
			want: 0,
		},
		{
			name: "29,15",
			args: args{
				A: 15,
				B: 29,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertInteger(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("convertInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
