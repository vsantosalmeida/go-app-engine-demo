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
	os.Setenv("DATASTORE_EMULATOR_HOST", "localhost:8081")
	os.Setenv("DATASTORE_PROJECT_ID", "app-engine-demo")
	port := getPort()
	router := mux.NewRouter()
	router.Use(commonMiddleware)

	router.HandleFunc("/contato", controller.GetPersons).Methods("GET")
	router.HandleFunc("/contato/{key}", controller.GetPerson).Methods("GET")
	router.HandleFunc("/contato", controller.CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{key}", controller.DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
