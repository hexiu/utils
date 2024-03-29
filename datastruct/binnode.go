package datastruct

import (
	"fmt"
)

// BinNode 二叉树节点
type BinNode struct {
	LC      *BinNode
	RC      *BinNode
	Parents *BinNode
	Data    interface{}
	height  int
	color   interface{}
	npl     interface{}
}

// UpdateHeight 更新节点的高度
type UpdateHeight interface {
	// 更新节点的高度
	UpdateHeight(b *BinNode)
}

// BinPosi 二叉树
type BinPosi struct {
	root *BinNode
	size int
	// UpdateHeight
}

// NewBinNode 创建一个新的二叉树节点
func NewBinNode(e interface{}, parent *BinNode) *BinNode {
	return &BinNode{
		Data:    e,
		Parents: parent,
	}
}

// NewBinPosi 创建一个新的二叉树
func NewBinPosi(binnode *BinNode) *BinPosi {
	return &BinPosi{
		root: binnode,
		size: 1,
	}
}

// InsertLc 插入左子树
func (b *BinNode) InsertLc(x *BinNode) {
	x.Parents = b
	b.LC = x
	return
}

// InsertRc 插入右子树
func (b *BinNode) InsertRc(x *BinNode) {
	x.Parents = b
	b.RC = x
	return
}

// updateheight 更新高度
func updateheight(binnode *BinNode, h int) {
	if binnode == nil {
		return
	}
	binnode.height += h
	binnode = binnode.Parents
	// binnode.height
	for binnode != nil {
		if binnode.LC != nil && binnode.RC != nil {
			if binnode.LC.height > binnode.RC.height {
				binnode.height = binnode.LC.height + 1
				binnode = binnode.Parents
				continue
			}
			binnode.height = binnode.RC.height + 1
			binnode = binnode.Parents
			continue
		}
		if binnode.LC == nil {
			binnode.height = binnode.RC.height + 1
			binnode = binnode.Parents
			continue
		}
		binnode.height = binnode.LC.height + 1
		binnode = binnode.Parents
	}
}

// Size 数量
func (b *BinNode) Size() (size int) {
	// 设置为1 ，记录本身
	size = 1
	if b.LC != nil {
		size += b.LC.Size()
	}
	if b.RC != nil {
		size += b.RC.Size()
	}
	return
}

// updateHeight 更新节点的高度
func (bp *BinPosi) updateHeight(b *BinNode, h int) {
	if b == nil {
		return
	}
	updateheight(b, h)
}

// updatePHeight 更新b以及祖先的高度
func (bp *BinPosi) updatePHeight(b *BinNode, h int) {
	bp.updateHeight(b, h)
	return
}

// Size 树的节点数
func (bp *BinPosi) Size() (size int) {
	return bp.size
}

// InsertLC 插入左子树
func (bp *BinPosi) InsertLC(x *BinNode, e *BinNode) *BinNode {
	bp.size++
	fmt.Println("l update:", x, e)
	x.InsertLc(e)
	bp.updatePHeight(x, 1)
	return x.LC
}

// InsertRC 插入右子树
func (bp *BinPosi) InsertRC(x *BinNode, e *BinNode) *BinNode {
	bp.size++
	fmt.Println("r update:", x, e)
	x.InsertRc(e)
	bp.updatePHeight(x, 1)
	return x.RC
}

// Traverse 先序遍历
func (bp *BinPosi) Traverse(b *BinNode, f func(binnode *BinNode)) {
	if b == nil {
		return
	}
	f(b)
	bp.Traverse(b.LC, f)
	bp.Traverse(b.RC, f)
}

// MidTraverse 中序遍历
func (bp *BinPosi) MidTraverse(b *BinNode, f func(binnode *BinNode)) {
	if b == nil {
		return
	}
	bp.MidTraverse(b.LC, f)
	f(b)
	bp.MidTraverse(b.RC, f)
}

// AfterTraverse 中序遍历
func (bp *BinPosi) AfterTraverse(b *BinNode, f func(binnode *BinNode)) {
	if b == nil {
		return
	}
	bp.AfterTraverse(b.RC, f)
	bp.AfterTraverse(b.LC, f)
	f(b)
}

func (bp *BinPosi) Succ(n *BinNode) (b *BinNode) {
	n = n.RC
	for {
		if n != nil {
			if n.LC != nil {
				n = n.LC
				continue
			}
			b = n
			return
		}
		break
	}
	return
}

// func visit(b *BinNode) {
// 	fmt.Println(b.Data)
// 	return
// }
