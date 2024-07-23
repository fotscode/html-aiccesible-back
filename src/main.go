 package main
 
import routes "html-aiccesible/routes"

func main() {
	r := routes.SetUpRouter()
	r.Run()
}
