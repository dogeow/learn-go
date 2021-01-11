package learn

import (
	"fmt"
	"os"
)

func Args() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[1])
}
