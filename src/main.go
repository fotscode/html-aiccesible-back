package main

import (
	"html-aiccesible/middleware"
	"html-aiccesible/models"
	routes "html-aiccesible/routes"
)

func main() {
	models.CreateDefaultUser()
	r := routes.SetUpRouter()
	r.Use(middleware.SetupCors())
	r.Run()
}
