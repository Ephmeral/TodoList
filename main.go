package main

import (
	"fmt"
	"github.com/Ephmeral/TodoList/conf"
	"github.com/Ephmeral/TodoList/routes"
)

func main() {
	conf.Init()
	r := routes.NewRoutes()
	fmt.Println("http port = ", conf.HttpPort)
	_ = r.Run(conf.HttpPort)
}
