package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/abbeyhrt/keep-up-graphql/internal/database"
	"github.com/abbeyhrt/keep-up-graphql/internal/models"
	"github.com/gorilla/mux"
)

type errorResponse struct {
	code    int    `json:"code"`
	message string `json:"message"`
}

func createUser(store database.DAL) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			resp, _ := json.Marshal(errorResponse{http.StatusInternalServerError, "error reading body"})

			w.Write(resp)
			return
		}
		var u models.User
		err = json.Unmarshal(b, &u)
		if err != nil {
			log.Println(err)
			resp, _ := json.Marshal(errorResponse{http.StatusInternalServerError, "error unmarshaling body"})

			w.Write(resp)
			return
		}

		user, err := store.CreateUser(u)
		if err != nil {
			log.Println(err)
			resp, _ := json.Marshal(errorResponse{http.StatusInternalServerError, err.Error()})
			w.Write(resp)
			return
		}

		resp, err := json.Marshal(user)
		if err != nil {
			log.Println(err)
			resp, _ := json.Marshal(errorResponse{http.StatusInternalServerError, err.Error()})
			w.Write(resp)
			return
		}

		w.Write(resp)
	}
}

// New handler for the server
func New(store database.DAL) http.Handler {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world"))
	})
	r.HandleFunc("/users", createUser(store)).Methods("POST")
	return r
}
