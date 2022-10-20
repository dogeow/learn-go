package main

import (
	"github.com/dogeow/learn-go/routers"
)

func main() {
	router := routers.InitRouter()

	router.Run(":3333")
}
