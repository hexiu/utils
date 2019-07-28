package datastruct

// Vertex 图节点
type Vertex struct {
	// 出入度数，状态信息
	inDegree, OutDegree, Vstatus, status int
	// 时间标签
	dTime, fTime int
	// 数据信息
	Data interface{}
}

// Class 类型
type Class struct {
	class int
}

//Edge 边
type Edge struct {
	Data   interface{}
	weight int
	class  Class
}

// 边类型
const (
	NULL     = 0
	Forward  = 1
	Backward = 2
	Cross    = 3
)

// GraphPosi 图
type GraphPosi struct {
	Vertexs []*Vertex
	Sides   [][]*Edge
}

// NewEdge 新的边
func NewEdge() *Edge {
	return &Edge{
		class: Class{
			class: NULL,
		},
	}
}

// NewVertex 新的节点
func NewVertex() *Vertex {
	return &Vertex{}
}

// NewGraphPosi 新建图
func NewGraphPosi() *GraphPosi {
	return &GraphPosi{
		Vertexs: make([]*Vertex, 0),
		Sides:   make([][]*Edge, 0),
	}
}
