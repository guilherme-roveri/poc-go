package api

import (
	"encoding/json"
	"net/http"

	"github.com/gbroveri/users/domains"
	"github.com/gbroveri/users/usecases"
	"github.com/gorilla/mux"
)

type UserAPI struct {
	crud usecases.UserUsecases
}

//NewUserAPIHandler factory
func NewUserAPIHandler(c usecases.UserUsecases) *UserAPI {
	return &UserAPI{crud: c}
}

//CreateUserHandler creates a new user
func (a *UserAPI) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var u domains.User
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if u, err = a.crud.Create(r.Context(), u); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

//GetHandler returns one user by iD
func (a *UserAPI) GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if user, err := a.crud.Get(r.Context(), id); err != nil {
		http.Error(w, err.Error(), 400)
		return
	} else {
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

}
