package learn

import "fmt"

func Datatype() {
	// 类型开关 (`datetype switch`) 比较类型而非值。可以用来发现一个接口值的类型。
	// 在这个例子中，变量 `t` 在每个分支中会有相应的类型。
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know datetype %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
