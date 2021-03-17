package main

import (
	"github.com/gorilla/mux"
	"go-app-engine-demo/pkg/controller"
	"log"
	"net/http"
	"os"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	return port
}

func main() {
	router := mux.NewRouter()
	port := getPort()

	router.HandleFunc("/contato", controller.GetPersons).Methods("GET")
	//router.HandleFunc("/contato/{id}", controller.GetPerson).Methods("GET")
	router.HandleFunc("/contato", controller.CreatePerson).Methods("POST")
	//router.HandleFunc("/contato/{id}", controller.DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":" + port, router))

}
