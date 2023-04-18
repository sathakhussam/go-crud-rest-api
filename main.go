package main

import (
	"log"
	"net/http"

	"github.com/sathakhussam/go-mod-crud/controllers"
)

var port string = "3000"

func main() {
	log.Println("Started On Port " + port)

	http.HandleFunc("/tasks", controllers.TasksMethodHandler)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
