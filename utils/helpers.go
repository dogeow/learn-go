package utils

import (
	"runtime"
)

func Add(x int, y int) int {
	return x + y
}

func Os() string {
	var goos string = runtime.GOOS
	return goos
}
