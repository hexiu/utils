package datastruct

import "testing"

func Test_Fib(t *testing.T) {
	var n = 10

	for i := 1; i < n; i++ {
		fibnum := Fib(i)
		t.Log(i, fibnum)
	}
}
