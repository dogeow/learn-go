package learn

import (
	"fmt"
	"runtime"
)

func System() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Running on", runtime.NumCPU(), "cores")
}
