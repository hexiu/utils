package datastruct

import "fmt"

/*
实现一个二叉搜索树
二叉搜索树 本质是一个 中序遍历结果 有序的树的实现，所以插入节点时要考虑有序性
*/

type SortFunc func(interface{}, interface{}) bool

type bst struct {
	*BinPosi
	SortFunc SortFunc
}

func NewBST(b *BinPosi, f SortFunc) *bst {
	return &bst{
		BinPosi:  b,
		SortFunc: f,
	}
}

func (b *bst) SetSortFunc(f SortFunc) {
	b.SortFunc = f
}

func (b *bst) Insert(v interface{}) {
	if b.root == nil {
		b.root = &BinNode{Data: v, height: 0}
		b.size = 1
		return
	}
	node := b.root
	for {
		if node == nil {
			return
		}
		fmt.Println("insert: ", v)
		if b.SortFunc(v, node.Data) {
			if node.LC == nil {
				b.InsertLC(node, NewBinNode(v, node))
				break
			}
			node = node.LC
			continue
		}
		if node.RC == nil {
			b.InsertRC(node, NewBinNode(v, node))
			break
		}
		node = node.RC
		continue
	}

}

// 节点删除之后 需要链接起来
func (b *bst) Remove(v interface{}) bool {
	return b.removeAt(v)

}

func (b *bst) removeAt(v interface{}) bool {
	// 根结点处理
	node := b.root
	for {
		if node == nil {
			break
		}
		if node.Data == v {
			if node.LC == nil {
				node.RC.Parents = node.Parents
				node.Parents.RC = node.RC
				if node.Parents == nil {
					b.root = node.RC
				}
			} else if node.RC == nil {
				node.LC.Parents = node.Parents
				node.Parents.LC = node.LC
				if node.Parents == nil {
					b.root = node.LC
				}
			} else {
				succ := b.Succ(node)
				fmt.Println("succ:", succ)
				{
					// succ 和当前节点做交换
					succ.Data, node.Data = node.Data, succ.Data
					// 所有的指针节点信息交换
					// 交换之后 对 node 信息进行处理
					if succ.LC == nil {
						succ.RC.Parents = succ.Parents
						fmt.Println("succ.RC.P:", succ.Parents.Data, succ.Data)
						succ.Parents.RC = succ.RC
						fmt.Println("succ.P.LC", succ.RC.Data)
						fmt.Println("root", b.root.LC.Data, b.root.RC.Data)
						return true
					}
					return true
				}
			}
		}
		if b.SortFunc(v, node.Data) {
			node = node.LC
			continue
		}
		node = node.RC
	}
	return false
}
