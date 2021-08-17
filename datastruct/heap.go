/*
 * @Date: 2021-08-17 15:56:09
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-08-17 18:56:29
 * @FilePath: /utils/datastruct/heap.go
 */
package datastruct

// 简单的堆实现代码
type HEAP interface {
}

type heap struct {
	*Vector
	SortFunc func(interface{}, interface{}) bool
}

func NewHeap(v *Vector) *heap {
	return &heap{Vector: v}
}

func (h *heap) SetSortFunc(f func(v1, v2 interface{}) bool) {
	h.SortFunc = f
}

// 父亲节点
func (h *heap) Parent(i int) (v int) {
	v = (i - 1) >> 1
	if v < 0 {
		v = -1
		return
	}
	return
}

// 左子树
func (h *heap) LC(i int) (v int) {
	if i < 0 {
		v = -1
		return
	}
	v = (i + 1) << 1
	if v > h.length {
		v = -1
		return
	}
	return
}

// 右子树
func (h *heap) RC(i int) (v int) {
	if i < 0 {
		v = -1
		return
	}
	v = (i+1)<<1 - 1
	if v > h.length {
		v = -1
		return
	}
	return
}

func (h *heap) Insert(v interface{}) bool {
	h.JudgeCap(h.length)
	h.data[h.length] = v
	h.length += 1
	return h.updateFilter(h.length - 1)
}

func (h *heap) Remove() (v interface{}) {
	if h.length <= 0 {
		return
	}
	v = h.data[0]
	h.data[0] = h.data[h.length-1]
	h.data[h.length-1] = nil
	h.length--
	h.downFilter()
	return
}

func (h *heap) updateFilter(index int) bool {
	// 执行上滤
	for {
		if index == 0 {
			return true
		}
		// fmt.Println(index, h.Parent(index), h.data)
		pindex := h.Parent(index)
		if pindex < 0 {
			return true
		}
		if h.Parent(index) > -1 {
			if h.SortFunc == nil {
				return false
			}
			if h.SortFunc(h.data[h.Parent(index)], h.data[index]) {
				h.data[h.Parent(index)], h.data[index] = h.data[index], h.data[h.Parent(index)]
			}
			index = h.Parent(index)
		} else {
			break
		}
	}
	return true
}

func (h *heap) downFilter() bool {
	index := 0
	if index < 0 || index > h.length {
		return false
	}
	for {
		lindex := h.LC(index)
		if lindex < 0 {
			return true
		}
		rindex := h.RC(index)
		if rindex < 0 {
			return true
		}
		// 只会和最大的孩子交换
		if h.SortFunc(h.data[index], h.data[h.LC(index)]) || h.SortFunc(h.data[index], h.data[h.RC(index)]) {
			if h.SortFunc(h.data[rindex], h.data[lindex]) {
				h.data[index], h.data[h.LC(index)] = h.data[h.LC(index)], h.data[index]
				index = h.LC(index)
			} else {
				h.data[index], h.data[h.RC(index)] = h.data[h.RC(index)], h.data[index]
				index = h.RC(index)
			}

		} else {
			break
		}
	}
	return true
}
