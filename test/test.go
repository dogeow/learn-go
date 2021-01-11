/**
结构体上的函数，像 JavaScript 的原型
*/
package test

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func (s *Saiyan) Super() {
	s.Power += 10000
}

func Test() {
	goku := &Saiyan{"Goku", 9001}
	goku.Super()

	fmt.Println(goku.Power) // 将会打印出 19001
}
