package common

import (
	"reflect"
	"testing"
)

func TestCreateSingleLinkListBySlice(t *testing.T) {
	type args struct {
		slice []interface{}
	}
	tests := []struct {
		name string
		args args
		want *LNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateSingleLinkListBySlice(tt.args.slice...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSingleLinkListBySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintSingleLinkList(t *testing.T) {
	type args struct {
		info string
		node *LNode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test print no node",
			args: args{
				info: "no node",
				node: CreateSingleLinkListBySlice(nil),
			},
		},
		{
			name: "test print one node",
			args: args{
				info: "one node",
				node: CreateSingleLinkListBySlice(1),
			},
		},
		{
			name: "test print two node",
			args: args{
				info: "two node",
				node: CreateSingleLinkListBySlice(1, 2),
			},
		},
		{
			name: "test print three node",
			args: args{
				info: "three node",
				node: CreateSingleLinkListBySlice(1, 2, 3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintSingleLinkList(tt.args.info, tt.args.node)
		})
	}
}
