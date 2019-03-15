package datastruct

import (
	"fmt"
	"testing"
)

func Test_ListNode(t *testing.T) {
	node := NewNode()
	cur := node
	for i := 1; i < 5; i++ {
		fmt.Println(i)
		node := new(Node)
		node.data = i
		cur.InsertAsSucc(node)
		cur = node
	}
	cur = node
	for cur != nil {
		fmt.Println(cur.Data())
		cur = cur.Succ()
	}
	fmt.Println(node.Succ().Succ().Data())
}
