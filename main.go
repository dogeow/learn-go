package main

import (
	"github.com/likunyan/learn-go/routers"
)

func main() {
	router := routers.InitRouter()

	router.Run(":3333")
}
