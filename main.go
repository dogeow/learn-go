package main

import (
	"github.com/likunyan/learn-go/learn"
	"github.com/likunyan/learn-go/routers"
)

func main() {
	learn.Db()
	return
	router := routers.InitRouter()

	router.Run(":3333")
}
