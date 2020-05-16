package main

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	root := &TreeNode{Val: 6}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 8}
	root.Left = node1
	root.Right = node2

	node3 := &TreeNode{Val: 0}
	node4 := &TreeNode{Val: 4}
	node1.Left = node3
	node1.Right = node4

	node5 := &TreeNode{Val: 7}
	node6 := &TreeNode{Val: 9}
	node2.Left = node5
	node2.Right = node6

	node7 := &TreeNode{Val: 3}
	node8 := &TreeNode{Val: 5}
	node4.Left = node7
	node4.Right = node8

	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "root root",
			args: args{
				root: root,
				p:    root,
				q:    root,
			},
			want: root,
		},
		{
			name: "1 nodep covers nodeq",
			args: args{
				root: root,
				p:    node1,
				q:    node8,
			},
			want: node1,
		},
		{
			name: "2 nodep covers nodeq",
			args: args{
				root: root,
				p:    node1,
				q:    node7,
			},
			want: node1,
		},
		{
			name: "3 nodep covers nodeq",
			args: args{
				root: root,
				p:    node1,
				q:    node4,
			},
			want: node1,
		},
		{
			name: "nodeq covers nodep",
			args: args{
				root: root,
				p:    node2,
				q:    node5,
			},
			want: node2,
		},
		{
			name: "1 nodep can't cover nodep",
			args: args{
				root: root,
				p:    node3,
				q:    node8,
			},
			want: node1,
		},
		{
			name: "2 nodep can't cover nodep",
			args: args{
				root: root,
				p:    node7,
				q:    node6,
			},
			want: root,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
				t.Errorf("got:%d want:%d", got.Val, tt.want.Val)
			}
		})
	}
}
