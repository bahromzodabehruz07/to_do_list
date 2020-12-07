package main

import (
	"github.com/to_do_list/controllers"
)

var server = controllers.Server{}

func main() {

	server.Initialize()
	server.Run(":8080")

}
