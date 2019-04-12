package datastruct

import (
	"fmt"
	"testing"
)

func Test_ListNode(t *testing.T) {
	// node := NewNode()
	// cur := node
	// for i := 1; i < 5; i++ {
	// 	fmt.Println(i)
	// 	node := new(Node)
	// 	node.data = i
	// 	cur.InsertAsSucc(node)
	// 	cur = node
	// }
	// cur = node
	// for cur != nil {
	// 	fmt.Println(cur.Data())
	// 	cur = cur.Succ()
	// }
	// fmt.Println("3: ", node.GetNode(3))
	// cur = node.tailer.Pred()
	// // fmt.Println("Find: ", node.Find(3, 3, nil))
	// fmt.Println(node.Succ().Succ().Data())

	posi := NewPosi()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		node := new(Node)
		node.data = i
		posi.InsertAsSucc(node)
	}
	cur := posi.header
	var node3 = new(Node)
	for cur != nil {
		fmt.Println(cur.Data())
		if cur.Data() == 3 {
			node3 = cur
		}
		cur = cur.Succ()
	}
	fmt.Println("Find: ", node3, posi.Find(1, 3, node3))

	fmt.Println("3: ", posi.GetNode(2), posi.length)
	node6 := new(Node)
	node6.data = 6
	// node3.InsertAsSucc(node6)
	posi.InsertBefore(node3, node6)
	cur = posi.header
	for cur != nil {
		fmt.Println(cur.Data())
		if cur.Data() == 3 {
			node3 = cur
		}
		cur = cur.Succ()
	}
}

func Test_Copy(t *testing.T) {
	posi := NewPosi()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		node := new(Node)
		node.data = i
		posi.InsertAsSucc(node)
	}
	cur := posi.header
	var node2 = new(Node)
	for cur.Succ() != nil {
		fmt.Println(cur.Data())
		if cur.Data() == 2 {
			node2 = cur
		}
		cur = cur.Succ()
	}
	fmt.Println(node2)
	p1 := posi.CopyNodes(node2, 2)
	cur = p1.header
	for cur.Succ() != nil {
		fmt.Println("copy: ", cur.Data())
		cur = cur.Succ()
	}
	fmt.Println(p1.String())
	p1.Remove(node2)
	fmt.Println("new:", p1.String())
}

func Test_Unique(t *testing.T) {
	posi := NewPosi()
	j := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		node := new(Node)
		node.data = j
		if i%2 == 0 {
			j++
		}
		posi.InsertAsSucc(node)
	}
	t.Log("old: ", posi.String())
	posi.Unique()
	t.Log("new: ", posi.String())
}

func Test_LinkSort(t *testing.T) {
	posi := NewPosi()
	j := 10
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		node := new(Node)
		node.data = j
		if i%2 == 0 {
			j--
		}
		posi.InsertAsSucc(node)
	}
	t.Log("old: ", posi.String())
	f := func(a, b interface{}) bool {
		if a == nil || b == nil {
			return false
		}
		if a.(int) >= b.(int) {
			return true
		}
		return false
	}
	posi.SelectSort(f)
	t.Log("new: ", posi.String())
}

func Test_InsertSort(t *testing.T) {
	posi := NewPosi()
	j := 10
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		node := new(Node)
		node.data = j
		if i%2 == 0 {
			j--
		}
		posi.InsertAsSucc(node)
	}
	t.Log("old: ", posi.String())
	f := func(a, b interface{}) bool {
		if a == nil || b == nil {
			return false
		}
		if a.(int) < b.(int) {
			return true
		}
		return false
	}
	posi.InsertSort(f)
	t.Log("new: ", posi.String())
}
