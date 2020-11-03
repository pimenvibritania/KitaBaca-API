package main

import (
	"github.com/pimenvibritania/rest-api/db"
	"github.com/pimenvibritania/rest-api/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
