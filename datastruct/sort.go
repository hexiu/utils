package datastruct

import "fmt"

// bubbleSort 起泡排序
func bubbleSort(v []interface{}, f func(v1, v2 interface{}) bool) bool {
	lo, hi := 0, len(v)
	if hi < 2 {
		return false
	}
	for lo != hi {
		hi = bubble(v, lo, hi, f)
	}
	return true
}

// bubble 起泡交换
func bubble(v []interface{}, lo, hi int, f func(v1, v2 interface{}) bool) (last int) {
	last = lo
	for ; lo < hi-1; lo++ {
		if f(v[lo], v[lo+1]) {
			v[lo], v[lo+1] = v[lo+1], v[lo]
			last = lo + 1
		}
	}
	return last
}

// mergeSort 归并排序
func mergeSort(v []interface{}, lo, hi int, f func(v1, v2 interface{}) bool) bool {
	if hi-lo < 2 {
		return true
	}
	mid := (hi + lo + 1) >> 1
	mergeSort(v, lo, mid, f)
	mergeSort(v, mid, hi, f)
	merge2(v, lo, mid, hi, f)

	return true
}

// merge 合并
func merge(v []interface{}, lo, mi, hi int, f func(v1, v2 interface{}) bool) {
	vsort := make([]interface{}, 0)
	fmt.Println(lo, mi, hi, v[lo:mi], v[mi:hi], vsort)
	low := lo
	mid := mi
	length := hi - low
	for len(vsort) < length {
		if f(v[lo], v[mi]) {
			if mi < hi {
				vsort = append(vsort, v[mi])
				mi++
			} else {
				vsort = append(vsort, v[lo])
				lo++
			}
		} else {
			if lo < mid {
				vsort = append(vsort, v[lo])
				lo++
			} else {
				vsort = append(vsort, v[mi])
				mi++
			}
		}
	}
	lo = low
	for i := 0; i < hi-low; i++ {
		v[lo] = vsort[i]
		lo++
	}
	// v = vsort
}

// merge2
func merge2(v []interface{}, lo, mi, hi int, f func(v1, v2 interface{}) bool) bool {
	a := make([]interface{}, mi-lo)
	a = v[lo:mi]

	b := make([]interface{}, hi-mi)
	b = v[mi:hi]
	c := make([]interface{}, hi-lo)
	la := mi - lo
	lb := hi - mi
	for i, j, k := 0, 0, 0; j < la; {
		// fmt.Println("i:", i, "j:", j, "k:", k, "la:", la, "lb:", lb, "hi:", hi, a, b, c)

		if k < lb && f(a[j], b[k]) {
			c[i] = b[k]
			k++
			i++
		}
		// fmt.Println("mid: i:", i, "j:", j, "k:", k, "lo:", lo, "mi:", mi, "hi:", hi, a, b, c)
		// fmt.Println("compare:", k, b, j, a, la, lb)
		if lb <= k || !f(a[j], b[k]) {
			c[i] = a[j]
			j++
			i++
		}
		if j == la {
			for k < lb {
				c[i] = b[k]
				i++
				k++
			}
			break
		}
		// fmt.Println("over:  i:", i, "j:", j, "k:", k, "lo:", lo, "mi:", mi, "hi:", hi, a, b, c)
	}
	length := hi - lo
	for i := 0; i < length; i++ {
		v[lo] = c[i]
		lo++
	}
	// fmt.Println("c: ", c)
	// fmt.Println("v: ", v)
	// fmt.Println("over: lo:", lo, "j:", j, "k:", k, "lo:", lo, "mi:", mi, "hi:", hi, a, b, v[:hi], c)

	return true
}
