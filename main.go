package main

import (
	"flag"
	"fmt"
	"shop-api/initialize"
)

func main() {
	port := flag.Int("port", 8021, "端口号")

	g := initialize.InitRouter()

	if err := g.Run(fmt.Sprintf(":%d", *port)); err != nil {
		panic("err:" + err.Error())
	}
}
