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
func updateheight(binnode *BinNode) {

	for binnode.height != -1 {
		// binnode.height
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
func (bp *BinPosi) updateHeight(b *BinNode) {
	if b == nil {
		return
	}
	if b.LC == nil {
		return
	}
	if b.RC == nil {
		return
	}
	if b.LC.height > b.RC.height {
		b.height = b.LC.height + 1
	} else {
		b.height = b.RC.height + 1
	}
}

// updatePHeight 更新b以及祖先的高度
func (bp *BinPosi) updatePHeight(b *BinNode) {
	for b != nil {
		bp.updateHeight(b.Parents)
		b = b.Parents
	}
	return
}

// Size 树的节点数
func (bp *BinPosi) Size() (size int) {
	return bp.size
}

// InsertLC 插入左子树
func (bp *BinPosi) InsertLC(x *BinNode, e *BinNode) *BinNode {
	bp.size++
	fmt.Println("update:", x, e)
	x.InsertLc(e)
	bp.updatePHeight(e)
	return x.LC
}

// InsertRC 插入右子树
func (bp *BinPosi) InsertRC(x *BinNode, e *BinNode) *BinNode {
	bp.size++
	x.InsertRc(e)
	bp.updatePHeight(x)
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
