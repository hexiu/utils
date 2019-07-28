package datastruct

import (
	"fmt"
	"testing"
)

// 队列
func Test_Queue(t *testing.T) {
	queue := NewNodeQueue()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		queue.EnQueue(i)
	}
	for j := 5; j > 0; j-- {
		// fmt.Println(j)
		fmt.Println(queue.DeQueue())
	}
	queue = NewNodeQueue()

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		queue.EnQueue(i)
	}
	for j := 5; j > 0; j-- {
		// fmt.Println(j)
		fmt.Println(queue.DeQueue())
	}
}
