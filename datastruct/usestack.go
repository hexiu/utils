package datastruct

import (
	"fmt"
)

type Flag struct {
	str   byte
	index int
}

// FindKuoHao 发现括号是否和对
func FindKuoHao(str string) bool {
	if str == "" {
		str = "(())()())()"
	}
	st := NewNodeStack()
	for i := 0; i < len(str); i++ {
		cur := str[i]
		if cur == '(' || cur == '[' || cur == '{' {

			f := new(Flag)
			f.index = i
			f.str = str[i]
			st.Push(f)
			continue
		}
		if cur == ')' || cur == ']' || cur == '}' {

			val := st.Top()
			var rel *Flag
			if val != nil {
				rel = val.(*Flag)
				if rel.str == '(' && cur == ')' || int(rel.str)-2 != int(byte(cur)) {
					fmt.Println("str:", str[rel.index:i+1])
					st.Pop()
					continue
				}
				return false
			}
			return false
		}
	}
	if st.Empty() {
		return true
	}
	return false
}
