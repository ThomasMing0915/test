package main

import (
	"LeetCode/common"
	"testing"
)

func TestReSortSingleLinkList(t *testing.T) {
	type args struct {
		node *common.LNode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "nil node",
			args: args{
				nil,
			},
		},
		{
			name: "head node",
			args: args{common.CreateSingleLinkListBySlice(nil)},
		},
		{
			name: "one node",
			args: args{common.CreateSingleLinkListBySlice(1)},
		},
		{
			name: "two node",
			args: args{common.CreateSingleLinkListBySlice(1, 2)},
		},
		{
			name: "three node",
			args: args{common.CreateSingleLinkListBySlice(1, 2, 3)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReSortSingleLinkList(tt.args.node)
		})
	}
}
