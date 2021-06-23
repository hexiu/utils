/*
 * @Date: 2021-06-16 11:10:01
 * @LastEditors: jaxiu
 * @LastEditTime: 2021-06-23 14:59:05
 * @FilePath: /utils/imath/num_test.go
 */
package imath

import (
	"fmt"
	"testing"
)

func TestNum(t *testing.T) {
	var m map[string]int = make(map[string]int, 0)
	for i := 0; i < 100000000; i++ {
		code, _ := Num2Code(int64(i))
		m[code]++
		if i > 99999990 {
			fmt.Println(code, i)
		}
		if code == "YDFD6V" {
			fmt.Println(code, int64(i)+baseId, length)

		}
		if len(code) != 6 {
			fmt.Println(code, int64(i)+baseId, length)

		}
	}
	var count int
	for k, v := range m {
		if v > 1 {
			fmt.Println(k, v)
			count++
		}
	}
	fmt.Println("lv:", float64(count)/10000000.0)
}
