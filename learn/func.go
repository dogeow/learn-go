package learn

import "fmt"

func NextInt(b []byte, i int) (int, int) {
	x := 0
	for ; i < len(b) && isDigit(b[i]); i++ {
		x = x*10 + int(b[i]) - '0'
		fmt.Println(x)
	}
	return x, i
}

func isDigit(b byte) bool {
	return true
}
