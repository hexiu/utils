package datastruct

import (
	"fmt"
	"strconv"
	"time"
)

func JinZhiChange(num, jz int) (rel string) {
	if jz == 1 {
		return strconv.Itoa(num)
	}
	s := NewNodeStack()
	var shang, yu int = 1, 0
	for shang > 0 {
		shang, yu = ShangYu(num, jz)
		fmt.Println(shang, yu)
		num = shang
		time.Sleep(1 * time.Second)
		s.Push(yu)
	}
	for ok := s.Pop(); ok != nil; {
		rel += fmt.Sprint(ok)
		ok = s.Pop()
	}

	return
}

func ShangYu(num, jz int) (shang, yu int) {
	return num / jz, num % jz
}
