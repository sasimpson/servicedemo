package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sasimpson/servicedemo/models"
)

//UserAPI -
type UserAPI struct{}

//InitUserRoutes will initialize the routes for the user API endpoint
func (h *Handler) InitUserRoutes(r *mux.Router) {
	sr := r.PathPrefix("/user").Subrouter()
	sr.Handle("", h.UserAllHandler()).Methods("GET").Name("UserGetAll")
	sr.Handle("/{id}", h.UserGetHandler()).Methods("GET").Name("UserGetByID")
}

//UserAllHandler will return the data for a single user record
func (h *Handler) UserAllHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var users *[]models.User
		users, err := h.User.All()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	})
}

//UserGetHandler will return the data for a single user record
func (h *Handler) UserGetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Implemented", http.StatusNotImplemented)
	})
}

//UserNewHandler will return the data for a single user record
func (h *Handler) UserNewHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Implemented", http.StatusNotImplemented)
	})
}
