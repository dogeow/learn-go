package learn

import "fmt"

// 这个 `intSeq` 函数返回另一个在 `intSeq` 函数体内定义的匿名函数。
// 这个返回的函数使用闭包的方式 _隐藏_ 变量 `i`。
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func Closures() {

	// 我们调用 `intSeq` 函数，将返回值（一个函数）赋给`nextInt`。
	// 这个函数的值包含了自己的值 `i`，
	// 这样在每次调用 `nextInt` 时都会更新 `i` 的值。
	nextInt := intSeq()

	// 通过多次调用 `nextInt` 来看看闭包的效果。
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
