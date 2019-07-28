package datastruct

import (
	"fmt"
	"testing"
)

func Test_BinTree(t *testing.T) {
	node := new(BinNode)
	node.Data = 0
	bintree := NewBinPosi(node)
	lc := new(BinNode)
	lc.Data = 1
	fmt.Println(bintree.InsertLC(bintree.root, lc))
	lc = new(BinNode)
	lc.Data = 2
	fmt.Println(bintree.InsertRC(bintree.root, lc))
	t.Log(bintree)
	t.Log("lc:", bintree.root, bintree.root.LC)
	visit := func(binnode *BinNode) {
		fmt.Println(binnode.Data)
	}
	bintree.Traverse(bintree.root, visit)
	bintree.MidTraverse(bintree.root, visit)
	bintree.AfterTraverse(bintree.root, visit)
	t.Log("ok")
}
