package main

import (
	"github.com/niqinge/order/model"
	"github.com/niqinge/order/server"
)

func main() {
	g := server.NewEngine(server.NewServer(
		&model.Order{},
		))
	if err := g.Run(":8081"); err != nil {
		panic(err)
	}
}
