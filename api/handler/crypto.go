package handler

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"go-app-engine-demo/pkg/crypto"
	"log"
	"net/http"
)

func decrypt() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var crypt *crypto.Crypto
		err := json.NewDecoder(r.Body).Decode(&crypt)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		svc := crypto.NewService(crypt)
		err = svc.Decrypt()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(string(svc.GetRaw())); err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func MakeCryptoHandlers(r *mux.Router, n negroni.Negroni) {
	r.Handle("/v1/decrypt", n.With(
		negroni.Wrap(decrypt()),
	)).Methods("POST").Name("decrypt")
}
