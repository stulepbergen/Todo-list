package main

import (
	"WebApp/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	rout := mux.NewRouter()

	rout.HandleFunc("/todo", controllers.Todo)
	rout.HandleFunc("/add", controllers.Add)
	rout.HandleFunc("/delete/{id}", controllers.Delete)
	rout.HandleFunc("/done/{id}", controllers.Done)

	log.Fatal(http.ListenAndServe(":8080", rout))
}
