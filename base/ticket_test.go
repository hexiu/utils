package base

import (
	"fmt"
	"testing"
	"time"
)

var num = 10
var con = 3

func Test_Ticket(t *testing.T) {
	ticket, _ := NewTicket(con)
	startT := time.Now()
	sign := make(chan int, num)
	for i := 0; i < num; {
		if ticket.Take() == nil {
			fmt.Println(i, ticket.Num())
			go func(i int) {
				sign <- i
				// time.Sleep(time.Millisecond)
				defer ticket.Return()
			}(i)
			i++
		} else {
			fmt.Println(i, "sleep. ")
			time.Sleep(time.Millisecond)
		}
	}
	// time.Sleep(10 * time.Second)
	var list []int
	for i := 0; i < num; i++ {
		list = append(list, <-sign)
	}

	fmt.Println(list)
	fmt.Println("done. ", "time is :", time.Now().Sub(startT))
	t.Log("ok. ")

}
