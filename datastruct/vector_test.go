package datastruct

import "testing"

func Test_Insert(t *testing.T) {
	vector := NewVector()
	for i := 0; i < 28; i++ {
		vector.Insert(i, i*i)
	}
	t.Log(vector, cap(vector.data))
}

func Test_Get(t *testing.T) {
	vector := NewVector()
	for i := 0; i < 28; i++ {
		vector.Insert(i, i*i)
	}
	for i := 0; i < vector.length; i++ {
		t.Log(i, vector.Get(i))
	}
}

func Test_Put(t *testing.T) {
	vector := NewVector()
	for i := 0; i < 28; i++ {
		vector.Insert(i, i*i)
	}
	t.Log("old: ", vector)
	vector.Put(3, 3*3*3)
	t.Log("new: ", vector)
}

func Test_Remove(t *testing.T) {
	vector := NewVector()
	for i := 0; i < 28; i++ {
		vector.Insert(i, i*i)
	}
	t.Log("old: ", vector)
	vector.Remove(3)
	t.Log("Remove: ", vector)
}

func Test_Find(t *testing.T) {
	vector := NewVector()
	for i := 0; i < 28; i++ {
		vector.Insert(i, i*i)
	}
	val := 9
	val1 := 168
	t.Log(val, "index:", vector.Find(val))
	vector.Insert(3, 27)
	t.Log(val1, "index:", vector.Find(val1))
}

func Test_Sort(t *testing.T) {
	vector := NewVector()
	// for i := 0; i < 28; i++ {
	// 	vector.Insert(i, i*i)
	// }
	vector.Insert(0, 52)
	vector.Insert(0, 1000)
	vector.Insert(0, 34)
	vector.Insert(0, 45)
	vector.Insert(0, 50)
	vector.Insert(0, 60)
	f := func(v1, v2 interface{}) bool {
		if v1 == nil || v2 == nil {
			return false
		}
		if v1.(int) > v2.(int) {
			return true
		}
		return false
	}
	t.Log(vector.Sort(f))
	t.Log(vector)
	vector1 := NewVector()
	for i := 0; i < 3; i++ {
		vector1.Insert(i, "a")
	}
	vector1.Insert(3, "a")
	vector1.Insert(3, "e")
	vector1.Insert(3, "f")
	vector1.Insert(3, "d")
	vector1.Insert(3, "c")
	vector1.Insert(3, "b")
	f1 := func(v1, v2 interface{}) bool {
		if v1 == nil || v2 == nil {
			return false
		}
		if v1.(string) > v2.(string) {
			return true
		}
		return false
	}
	vector1.Sort(f1)
	t.Log(vector1)
}

func Test_Disordered(t *testing.T) {
	vector := NewVector()
	for i := 0; i < 28; i++ {
		vector.Insert(i, i*i)
	}
	vector.Insert(5, 1000)
	vector.Insert(5, 34)
	vector.Insert(5, 45)
	t.Log(vector)
	f := func(v1, v2 interface{}) bool {
		if v1 == nil || v2 == nil {
			return false
		}
		if v1.(int) > v2.(int) {
			return true
		}
		return false
	}
	t.Log(vector.Disordered(f))
}

func Test_Uniquify(t *testing.T) {
	vector := NewVector()
	for i := 0; i < 28; i++ {
		vector.Insert(i, i*i)
	}
	vector.Insert(3, 9)
	vector.Insert(7, 36)
	t.Log(vector)
	f := func(v1, v2 interface{}) bool {
		if v1 == nil || v2 == nil {
			return false
		}
		if v1.(int) > v2.(int) {
			return true
		}
		return false
	}
	t.Log(vector.Uniquify(f), vector)

}

func Test_Search(t *testing.T) {
	vector := NewVector()
	for i := 0; i < 28; i++ {
		vector.Insert(i, i*i)
	}
	val := 9
	val1 := 168
	t.Log(val, "index:", vector.Search(val))
	vector.Insert(3, 27)
	t.Log(val1, "index:", vector.Search(val1))
}

func Test_FibOnAcci(t *testing.T) {
	vector := NewVector()
	for i := 0; i < 28; i++ {
		vector.Insert(i, i*i)
	}
	val := 9
	val1 := 168
	t.Log(val, "index:", vector.FibOnAcci(val))
	vector.Insert(3, 27)
	t.Log(val1, "index:", vector.FibOnAcci(val1))
}
