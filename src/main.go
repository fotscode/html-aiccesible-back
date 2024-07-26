package main

import (
	"html-aiccesible/middleware"
	routes "html-aiccesible/routes"
)

func main() {
	r := routes.SetUpRouter()
	r.Use(middleware.SetupCors())
	r.Run()
}
