package datastruct

import "fmt"

// pred()
// succ()
// data()
// insertAsPred()
// insertAsSucc()

// Node 节点
type Node struct {
	data interface{}
	succ *Node
	pred *Node
}

// Posi 列表
type Posi struct {
	header *Node
	tailer *Node
	length int
}

// NewPosi 创建一个列表
func NewPosi() (posi *Posi) {
	posi = new(Posi)
	var header = new(Node)
	var tailer = new(Node)
	header.succ = tailer
	tailer.pred = header
	posi.header = header
	posi.tailer = tailer
	return posi
}

// NewNode 创建一个新的节点
func NewNode() *Node {
	node := new(Node)
	return node
}

// Pred 前驱
func (n *Node) Pred() *Node {
	return n.pred
}

// Succ 后继
func (n *Node) Succ() *Node {
	return n.succ
}

// Data 数据
func (n *Node) Data() (d interface{}) {
	return n.data
}

// InsertAsPred 插入前驱
func (n *Node) InsertAsPred(node *Node) {
	pred := n.Pred()
	pred.succ = node
	node.pred = pred
	node.succ = n
	n.pred = node
}

// InsertAsSucc 插入后继
func (n *Node) InsertAsSucc(node *Node) {
	succ := n.Succ()

	succ.pred = node

	node.succ = succ
	node.pred = n
	n.succ = node
}

// GetNode 获取指定第几个数据
func (posi *Posi) GetNode(r int) (node *Node) {
	if r > posi.length {
		return nil
	}
	node = posi.header
	for r != 0 {
		node = node.Succ()
		r--
	}
	return node
}

// Find 查找
func (posi *Posi) Find(e interface{}, r int, p *Node) (node *Node) {
	if p != nil {
		node = p
	} else {
		node = posi.tailer
	}
	for r != 0 {
		node = node.Pred()
		if node.Data() == e {
			return node
		}
		r--
	}
	return nil
}

// InsertAsPred 插入
func (posi *Posi) InsertAsPred(node *Node) {
	if posi.length == 0 {
		posi.header.succ = node
		posi.tailer.pred = node
		node.pred = posi.header
		node.succ = posi.tailer
		posi.length++
	} else {
		next := posi.header.succ
		posi.header.succ = node
		node.pred = posi.header
		node.succ = next
		next.pred = node
		posi.length++
	}
}

// InsertAsSucc 插入在尾部
func (posi *Posi) InsertAsSucc(node *Node) {
	if posi.length == 0 {
		posi.header.succ = node
		posi.tailer.pred = node
		node.pred = posi.header
		node.succ = posi.tailer
		posi.length++
	} else {
		up := posi.tailer.pred
		posi.tailer.pred = node
		node.pred = up
		node.succ = posi.tailer
		up.succ = node
		posi.length++
	}
}

// InsertBefore 节点前插入
func (posi *Posi) InsertBefore(node *Node, new *Node) {
	node.InsertAsPred(new)
	posi.length++
}

// InsertAfter 节点后插入
func (posi *Posi) InsertAfter(node *Node, new *Node) {
	node.InsertAsSucc(new)
	posi.length++
}

// Remove 删除节点
func (posi *Posi) Remove(node *Node) {
	pred := node.Pred()
	succ := node.Succ()
	pred.succ = succ
	succ.pred = pred
	posi.length--
}

// CopyNodes 部分节点copy
func (posi *Posi) CopyNodes(node *Node, n int) (p *Posi) {
	p = NewPosi()
	for n != 0 {
		if node.Succ() == nil {
			return
		}
		nodei := node
		node = node.Succ()
		p.InsertAsSucc(nodei)
		n--
	}
	return
}

// String 打印
func (posi *Posi) String() (out string) {
	cur := posi.header.Succ()
	for cur.Succ() != nil {
		out += fmt.Sprint(cur.Data(), " ")
		cur = cur.Succ()
	}
	return
}

// Unique 去重
func (posi *Posi) Unique() {
	if posi.length < 2 {
		return
	}
	node := posi.header.succ
	for posi.tailer != node {
		tm := node.succ
		if tm.data == node.data {
			posi.Remove(tm)
		} else {
			node = tm
		}
	}
	return
}

// FindAll 查找
func (posi *Posi) FindAll(e interface{}) (node *Node) {
	node = posi.header.succ
	for posi.tailer != node {
		if e == node.data {
			return node
		}
	}
	return nil
}
