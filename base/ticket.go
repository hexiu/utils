package base

import "errors"

type ticket struct {
	ticket chan int
	num    int
	cap    int
	active bool
}

// Ticket 票盘
type Ticket interface {
	Take() (err error)
	Return() (err error)
	Num() (num int)
	Active() (active bool)
}

// NewTicket 新的票盘
func NewTicket(num int) (t Ticket, err error) {
	if num == 0 {
		return nil, errors.New("num is not 0, this is error! ")
	}
	gt := new(ticket)
	if gt.init(num) {
		return gt, nil
	}
	return nil, errors.New("init error")

}

// Init init
func (t *ticket) init(num int) (ok bool) {
	if num < 1 {
		num = 10
	}
	t.ticket = make(chan int, num)
	t.active = true
	t.num = num
	t.cap = num
	for i := 0; i < num; i++ {
		t.ticket <- 1
	}
	ok = true
	return
}

// Take get
func (t *ticket) Take() (err error) {
	if !t.active {
		return errors.New("ticket is not init. ")
	}
	// if t.num < 1 {
	// 	return errors.New("no ticket. ")
	// }
	<-t.ticket
	t.num--
	return
}

// Return return
func (t *ticket) Return() (err error) {
	if !t.active {
		return errors.New("ticket is not init. ")
	}
	// if t.num+1 > t.cap {
	// 	return errors.New("ticket ok. ")
	// }
	t.ticket <- 1
	t.num++
	return
}

// Num num
func (t *ticket) Num() (num int) {
	return t.num
}

// Active active
func (t *ticket) Active() (active bool) {
	return t.active
}
