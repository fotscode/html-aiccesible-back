package main

import (
	"html-aiccesible/models"
	routes "html-aiccesible/routes"
)

func main() {
	models.CreateDefaultUser()
	r := routes.SetUpRouter()
	r.Run()
}
