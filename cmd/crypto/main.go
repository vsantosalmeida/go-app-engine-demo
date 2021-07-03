package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/vsantosalmeida/go-app-engine-demo/api/handler"
	"github.com/vsantosalmeida/go-app-engine-demo/config"
	"github.com/vsantosalmeida/go-app-engine-demo/pkg/midleware"
)

func main() {
	router := mux.NewRouter()

	//handlers
	n := negroni.New(
		negroni.HandlerFunc(midleware.Cors),
		negroni.NewLogger(),
	)

	//crypto
	handler.MakeCryptoHandlers(router, *n)

	http.Handle("/", router)
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.CryptoApiPort),
		ErrorLog:     logger,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
