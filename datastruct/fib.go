package datastruct

// Fib 生成Fib数
func Fib(n int) int {
	var a, b int = 0, 1
	var i = 1
	if n == 1 {
		return 1
	}
	for i < n {
		b = a + b
		a = b - a
		i++
	}
	return b
}

// FibNum fib最贴近第几个值
func FibNum(val int) int {
	for i := 0; i < val; i++ {
		if Fib(i) > val {
			return i
		}
	}
	return -1
}
