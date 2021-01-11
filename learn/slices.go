package learn

import (
	"fmt"
)

func Slices() {
	a := make([]string, 10)
	fmt.Println("a:", a)
	fmt.Println("a 长度", len(a))
	fmt.Println("a 容量", cap(a))

	b := make([]int, 0, 5)
	fmt.Println("a:", a)
	fmt.Println("b 长度", len(b))
	fmt.Println("b 容量", cap(b))

	c := cap(b)
	for i := 0; i < 25; i++ {
		b = append(b, i)

		// 如果容量改变了
		// Go 必须增加数组长度来容纳新的数据
		if cap(b) != c {
			c = cap(b)
			fmt.Println(c)
		}
	}

	// copy
	d := make([]int, len(b))
	copy(d, b)
	fmt.Println("cpy:", d)

	// 切片
	l := d[5:10]
	fmt.Println("sl1:", l)
}
