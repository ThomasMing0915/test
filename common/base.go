package common

// single link list node
type LNode struct {
	Data interface{}
	Next *LNode
}

func NewLNode(i interface{}) *LNode {
	return &LNode{Data: i, Next: nil}
}
