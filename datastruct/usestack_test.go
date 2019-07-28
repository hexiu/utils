package datastruct

import "testing"

func Test_KuoHao(t *testing.T) {
	t.Log(FindKuoHao("{()ggggg}([]([{}]))()"))
	t.Log(byte('('), byte(')'))
	t.Log(byte('['), byte(']'))
	t.Log(byte('{'), byte('}'))
}
