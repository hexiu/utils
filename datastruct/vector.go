package datastruct

import (
	"fmt"
	"reflect"
)

// Vector 向量
type Vector struct {
	data   []interface{}
	length int
	class  string
}

// NewVector 创建一个新的向量
func NewVector() *Vector {
	return &Vector{
		data:   make([]interface{}, 1),
		length: 0,
		class:  "string",
	}
}

// SetType 设置类型
func (v *Vector) SetType(class string) {
	v.class = class
	return
}

// JudgeCap 判断是否满了
func (v *Vector) JudgeCap(i int) bool {
	capt := cap(v.data)
	var length int
	if i != 0 {
		length = i
	} else {
		length = v.length
	}
	if length < capt {
		return false
	}
	var vector = make([]interface{}, capt*2)
	for i := 0; i < v.length; i++ {
		vector[i] = v.data[i]
	}
	v.data = vector
	return true
}

// Insert 插入数据
func (v *Vector) Insert(i int, x interface{}) bool {
	class := reflect.TypeOf(x).String()
	if v.length == 0 {
		v.SetType(class)
	}
	if i > v.length || class != v.class {
		return false
	}
	length := v.length
	for ; length > i; length-- {
		v.Put(length, v.data[length-1])
	}
	v.Put(i, x)
	v.length++
	return true
}

// Put 修改数据
func (v *Vector) Put(i int, x interface{}) bool {
	if i+1 > v.length {
		v.JudgeCap(i)
	}
	v.data[i] = x
	return true
}

// Get 获取数据
func (v *Vector) Get(i int) interface{} {
	if v.length > i {
		return v.data[i]
	}
	return -1
}

// Remove 移除元素
func (v *Vector) Remove(i int) interface{} {
	if i > v.length {
		return -1
	}
	val := v.data[i]
	if v.length == 1 {
		v.length--
		return val
	}
	for ; i < v.length; i++ {
		v.data[i] = v.data[i+1]
	}
	v.length--
	return val
}

// Size 返回向量大小
func (v *Vector) Size() int {
	return v.length
}

// Find 发现一个元素所在的位置
func (v *Vector) Find(val interface{}) int {

	for i := 0; i < v.length; i++ {
		if v.data[i] == val {
			return i
		}
	}
	return -1
}

// Search 有序查找
func (v *Vector) Search(val interface{}) int {
	var class reflect.Type
	if v.length > 0 {
		class = reflect.TypeOf(v.data[0])
	}
	if class.Comparable() {
		return find(v.data, 0, v.length, val)
	}
	for i := 0; i < v.length; i++ {
		if v.data[i] == val {
			return i
		}
	}

	return -1
}

// FibOnAcci 有序查找
func (v *Vector) FibOnAcci(val interface{}) int {
	var class reflect.Type
	if v.length > 0 {
		class = reflect.TypeOf(v.data[0])
	}
	if class.Comparable() {
		fibn := FibNum(v.length)
		return findOnFibAcci(v.data, 0, v.length, fibn, val)
	}
	for i := 0; i < v.length; i++ {
		if v.data[i] == val {
			return i
		}
	}

	return -1
}

func find(v []interface{}, lo, hi int, val interface{}) int {
	fmt.Println("start:", v, lo, hi, val)
	if lo > hi {
		return -1
	}
	if v[lo-lo] == val {
		return lo
	}
	if v[hi-lo-1] == val {
		return hi
	}
	if lo == hi-1 {
		return -1
	}
	mid := (lo + hi) / 2

	val1 := reflect.ValueOf(val)
	val2 := reflect.ValueOf(v[mid-lo])
	switch val.(type) {
	case int:
		v1 := val1.Int()
		v2 := val2.Int()
		if v1 < v2 {
			return find(v[lo-lo:mid-lo], lo, mid, val)
		}
		return find(v[mid-lo:hi-lo], mid, hi, val)

	case string:
		v1 := val1.String()
		v2 := val2.String()
		if v1 < v2 {
			return find(v[lo-lo:mid-lo], lo, mid, val)
		}
		return find(v[mid-lo:hi-lo], mid, hi, val)

	default:
		return -1
	}

}

func findOnFibAcci(v []interface{}, lo, hi, fibn int, val interface{}) int {
	fmt.Println("start:", v, lo, hi, val)
	if lo > hi {
		return -1
	}
	mid := int(Fib(fibn - 1))
	if hi-lo == 1 && v[mid] == val {
		return mid
	}
	val1 := reflect.ValueOf(val)
	val2 := reflect.ValueOf(v[mid-lo])
	switch val.(type) {
	case int:
		v1 := val1.Int()
		v2 := val2.Int()
		if v1 < v2 {
			return find(v[lo-lo:mid], lo, mid, val)
		}
		return find(v[mid-lo:hi-lo], mid, hi, val)

	case string:
		v1 := val1.String()
		v2 := val2.String()
		if v1 < v2 {
			return find(v[lo-lo:mid-lo], lo, mid, val)
		}
		return find(v[mid-lo:hi-lo], mid, hi, val)

	default:
		return -1
	}

}

// Sort 对元素进行排序
func (v *Vector) Sort(f func(v1, v2 interface{}) bool) bool {
	if v.class == "" {
		return false
	}
	// return bubbleSort(v.data, f)
	lo, hi := 0, v.length
	return mergeSort(v.data, lo, hi, f)
}

// Disordered 发现有多少无序对
func (v *Vector) Disordered(f func(v1, v2 interface{}) bool) (count int) {
	for i := 1; i < v.length; i++ {
		if f(v.data[i-1], v.data[i]) {
			count++
		}
	}
	return
}

// SortFunc 排序函数
func (v *Vector) SortFunc(f func(x1, x2 interface{}) bool) bool {
	if v.class == "" {
		return false
	}
	for sorted := false; !sorted; {
		sorted = true
		for i := 1; i < v.length; i++ {
			val1 := v.data[i-1]
			val2 := v.data[i]
			if f(val1, val2) {
				v.data[i-1], v.data[i] = v.data[i], v.data[i-1]
				sorted = false
			}
		}
	}
	return true
}

// Uniquify 只要唯一项
func (v *Vector) Uniquify(f func(x1, x2 interface{}) bool) bool {
	if v.Disordered(f) == 0 {
		for i := 1; i < v.length; i++ {
			if v.data[i-1] == v.data[i] {
				v.Remove(i)
			}
		}
		return true
	}
	return false
}
