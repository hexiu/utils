package datastruct

import (
	"fmt"
)

// Stack 栈实现
type Stack interface {
	Pop() (val interface{})
	Push(x interface{}) bool
	Top() interface{}
}

// NewStack 初始化栈
func NewStack() (st Stack) {
	st = &stack{
		*NewVector(),
	}
	return
}

type stack struct {
	Vector
}

// Pop 出栈
func (s *stack) Pop() (val interface{}) {
	if s.Size() == 0 {
		return nil
	}
	fmt.Println(s.Size())
	return s.Remove(s.Size() - 1)
}

// Push 入栈
func (s *stack) Push(x interface{}) bool {
	return s.Insert(s.Size(), x)
}

// Top 查看
func (s *stack) Top() interface{} {
	return s.Get(s.Size() - 1)
}

type nodestack struct {
	Posi
}

// NewNodeStack 列表类的栈
func NewNodeStack() Stack {
	return &nodestack{
		*NewPosi(),
	}
}

func (n *nodestack) Pop() (val interface{}) {
	if n.length == 0 {
		return nil
	}
	node := n.header.Succ()
	n.Remove(node)
	return node.Data()
}

func (n *nodestack) Top() interface{} {

	return n.header.Succ().Data()
}

func (n *nodestack) Push(x interface{}) bool {
	node := NewNode()
	node.data = x
	n.InsertAsPred(node)
	return true
}
