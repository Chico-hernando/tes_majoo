package main

import (
	"majoo/configs"
	"majoo/routes"
)

func main() {
	configs.InitDB()

	e := routes.New()
	e.Start(":8000")
}