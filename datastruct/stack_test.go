package datastruct

import "testing"

func Test_Stack(t *testing.T) {
	s := NewStack()
	t.Log(s.Push(5))
	t.Log(s.Top())
	t.Log(s.Pop())
	t.Log(s.Pop())
	s = NewNodeStack()
	t.Log(s.Push(5))
	t.Log(s.Push(6))
	t.Log(s.Push(7))
	t.Log(s.Push(8))
	t.Log(s.Pop())
	t.Log(s.Pop())

	t.Log(s.Top())
}
