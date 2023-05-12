package main

import (
	"os"
)

func main() {
	_, err := os.Stat("/tmp/not_found_main.go")
	if err != nil {
		println("233")
	}
}
