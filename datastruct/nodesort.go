package datastruct

import "fmt"

// SelectSort 选择排序(两种方式一种搬动，一种是复制)
func (posi *Posi) SelectSort(f func(a, b interface{}) bool) {
	length := posi.length
	tailer := posi.tailer

	for length != 0 {
		node := posi.header.Succ()
		fmt.Println("node:", node)
		// fmt.Println("length: ", length, "a:", node.data, "b:", node.succ.data)
		fornum := length - 1
		maxnode := node
		for fornum != 0 {
			// fmt.Println("length: ", fornum, "a:", node.data, "b:", node.succ.data)
			if node.Succ() == posi.tailer {
				break
			}
			if f(maxnode.data, node.succ.data) {
				maxnode = node.succ
				fmt.Println("max & node:", maxnode.data, node.succ.data)
			}
			fmt.Println("M:", maxnode, fornum)
			node = node.Succ()
			fornum--
		}
		if maxnode.succ == tailer {
			return
		}
		fmt.Println("max: ", maxnode, length)
		maxnode.pred.succ = maxnode.succ
		maxnode.succ.pred = maxnode.pred
		tailer.InsertAsPred(maxnode)
		// posi.Remove(maxnode)
		// fmt.Println("max: ", maxnode)
		// newposi.InsertAsPred(maxnode)
		fmt.Println("newposi:", posi.String())
		length--
	}
	return
}

// InsertSort 插入排序
func (posi *Posi) InsertSort(f func(a, b interface{}) bool) {
	// rnode := posi.header
	if posi.length < 2 {
		return
	}
	var leftnode, rightnode *Node
	rightnode = posi.header.Succ().Succ()
	rnode := rightnode
	leftnode = posi.header.succ

	for leftnode != posi.tailer {
		node := posi.header.succ
		for node != leftnode {
			if f(node.data, rnode.data) {
				if node.pred == posi.header {
					rnode.pred.succ = rnode.succ
					rnode.succ.pred = rnode.pred
					rightnode = rightnode.Succ()
					fmt.Println("data:", node.data, rnode.data, node.succ.data)
					node.InsertAsPred(rnode)
					fmt.Println("time1:", rnode.data, posi.String())
					break
				}
			} else {
				if f(node.succ.data, rnode.data) {
					rightnode = rightnode.Succ()
					rnode.pred.succ = rnode.succ
					rnode.succ.pred = rnode.pred
					fmt.Println("data:", node.data, rnode.data, node.succ.data)
					node.InsertAsSucc(rnode)
					fmt.Println("time1:", rnode.data, posi.String())
					break
				}
			}
			// fmt.Println("string:", posi.String())
			node = node.Succ()
		}
		leftnode = leftnode.Succ()

		// rightnode = rightnode.Succ()
		rnode = rightnode
		fmt.Println("time:", rnode.data, posi.String())
	}
	fmt.Println("rel:", posi.String())
}
