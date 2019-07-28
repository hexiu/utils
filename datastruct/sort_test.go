package datastruct

import "testing"

func Test_Qsort(t *testing.T) {
	list := []interface{}{2, 3, 1, 10, 5, 8, 10, 4, 7, 6}
	t.Log(Qsort(list, 0, len(list)-1))
	t.Log(list)

}
