package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/vsantosalmeida/go-app-engine-demo/pkg/crypto"
)

func decrypt() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var crypt *crypto.Crypto
		err := json.NewDecoder(r.Body).Decode(&crypt)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		svc := crypto.NewService(crypt)
		err = svc.Decrypt()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(string(svc.GetDecryptRaw())); err != nil {
			log.Println(err)
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
