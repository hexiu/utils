/*
 * @Date: 2021-08-18 15:27:04
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-08-23 10:50:36
 * @FilePath: /utils/datastruct/bst_test.go
 */

package datastruct

import (
	"fmt"
	"testing"
)

func Test_BST(t *testing.T) {
	bst := NewBST(&BinPosi{}, func(v1, v2 interface{}) bool {
		if v1 == nil || v2 == nil {
			return false
		}
		if v1.(int) < v2.(int) {
			return true
		}
		return false
	})
	bst.root = nil
	bst.Insert(3)
	bst.Insert(5)
	bst.Insert(6)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(9)
	fmt.Println("mid:")
	bst.MidTraverse(bst.root, func(binnode *BinNode) {
		if binnode == nil {
			return
		}
		fmt.Println(binnode.Data, binnode.height)
	})
	fmt.Println("remove 6: ")
	bst.Remove(6)
	fmt.Println("mid:")
	bst.MidTraverse(bst.root, func(binnode *BinNode) {
		fmt.Println(binnode.Data, binnode.height)
	})
}
