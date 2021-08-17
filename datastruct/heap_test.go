/*
 * @Date: 2021-08-17 17:12:40
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-08-17 19:01:10
 * @FilePath: /utils/datastruct/heap_test.go
 */
package datastruct

import (
	"fmt"
	"testing"
)

func Test_Heap(t *testing.T) {
	v := NewVector()
	v.length = 0
	h := NewHeap(v)
	h.SetSortFunc(func(v1, v2 interface{}) bool {
		if v1 == nil || v2 == nil {
			return false
		}
		if v1.(int) > v2.(int) {
			return true
		}
		return false
	})
	h.Insert(1)
	h.Insert(8)
	h.Insert(5)
	h.Insert(3)
	h.Insert(9)
	h.Insert(7)
	fmt.Println(h.data)
	fmt.Println(h.Remove())
	fmt.Println(h.data)
	fmt.Println(h.Remove())
	fmt.Println(h.data)
	fmt.Println(h.Remove())
	fmt.Println(h.data)
	fmt.Println(h.Remove())
	fmt.Println(h.data)
	fmt.Println(h.Remove())
	fmt.Println(h.data)
	fmt.Println(h.Remove())
	fmt.Println(h.data)
	fmt.Println(h.Remove())
	fmt.Println(h.data)
	h.Insert(1)
	h.Insert(8)
	h.Insert(5)
	h.Insert(3)
	h.Insert(9)
	h.Insert(7)
	fmt.Println(h.data)

}
