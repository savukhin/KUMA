package main

import (
	"server/api"
)

func main() {
	app := api.SetupRouter()

	app.Listen(":8080")
}
