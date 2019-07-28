package datastruct

// Queue 队列
type Queue interface {
	EnQueue(x interface{}) bool
	DeQueue() interface{}
	Empty() bool
}

type queue struct {
	Vector
}

// NewQueue 创建队列
func NewQueue() Queue {
	return &queue{
		*NewVector(),
	}
}

func (q *queue) EnQueue(x interface{}) bool {
	return q.Insert(q.Size(), x)
}

func (q *queue) DeQueue() interface{} {
	if !q.Empty() {
		return q.Remove(0)
	}
	return nil
}

func (q *queue) Empty() bool {
	if q.Size() > 0 {
		return false
	}
	return true
}

type nodequeue struct {
	Posi
}

// NewNodeQueue 列表队列
func NewNodeQueue() Queue {
	return &nodequeue{
		*NewPosi(),
	}
}

func (nq *nodequeue) EnQueue(x interface{}) bool {
	node := NewNode()
	node.data = x
	nq.InsertAsPred(node)
	return true
}

func (nq *nodequeue) DeQueue() interface{} {
	if !nq.Empty() {
		node := nq.tailer.Pred()
		nq.Remove(node)
		return node.Data()
	}
	return nil
}

func (nq *nodequeue) Empty() bool {
	if nq.length > 0 {
		return false
	}
	return true
}
