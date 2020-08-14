package main

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	root := &TreeNode{
		Val: 1,
	}
	node1 := &TreeNode{
		Val: 2,
	}
	node2 := &TreeNode{
		Val: 3,
	}
	root.Left = node1
	root.Right = node2

	tests := []struct {
		name string
		args args
		want []int
	}{
		//{
		//	name: "root is null",
		//	args: args{
		//		root: nil,
		//	},
		//	want: []int{},
		//},
		{
			name: "one node",
			args: args{
				root,
			},
			want: []int{2, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
