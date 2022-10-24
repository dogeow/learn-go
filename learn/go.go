package learn

import (
	"fmt"
	"runtime"
)

func Go() {
	fmt.Printf("%s", runtime.Version())
	fmt.Printf("%s", runtime.GOOS)
}
