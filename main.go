package main

import (
	"github.com/alamansson/basic-CRUD-rest-api/handler"
)

func main() {
	run := handler.InitRoutes()
	run.Run(":3000")
}
