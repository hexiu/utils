package datastruct

/*
实现一个二叉搜索树
二叉搜索树 本质是一个 中序遍历结果 有序的树的实现，所以插入节点时要考虑有序性
*/

type SortFunc func(interface{}, interface{}) bool

type bst struct {
	BinPosi
	SortFunc SortFunc
}

func (b *bst) SetSortFunc(f SortFunc) {
	b.SortFunc = f
}

func (b *bst) Insert(v interface{}) {
	if b.root == nil {
		b.root = &BinNode{Data: v, height: 0}
		b.size = 1
	}
	node := b.root
	for {
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
func (b *bst) Remove(v interface{}) (n *BinNode) {
	if b.root == nil {
		return
	}
	return
}
