package datastruct

// pred()
// succ()
// data()
// insertAsPred()
// insertAsSucc()

// Node 节点
type Node struct {
	header *Node
	pred   *Node
	data   interface{}
	succ   *Node
	tailer *Node
}

// NewNode 创建一个新的节点
func NewNode() *Node {
	var header = new(Node)
	var tailer = new(Node)
	node := new(Node)
	header.succ = tailer
	tailer.pred = header
	node.header = header
	node.tailer = tailer
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

	if succ != nil {
		succ.pred = node
	}
	node.succ = succ
	node.pred = n
	n.succ = node
}
