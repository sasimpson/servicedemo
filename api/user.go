package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	})
}

//UserGetHandler will return the data for a single user record
func (h *Handler) UserGetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			idVar string
			id    int
			err   error
			user  *models.User
		)

		vars := mux.Vars(r)
		log.Println(vars)

		if idVar := vars["id"]; idVar == "" {
			log.Printf("idVar: %s", idVar)
			http.Error(w, "no id for user", http.StatusBadRequest)
			return
		}

		if id, err = strconv.Atoi(idVar); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		if user, err = h.User.Get(id); err != nil {
			if err == models.ErrUserNotFound {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if user == nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

//UserNewHandler will return the data for a single user record
func (h *Handler) UserNewHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Implemented", http.StatusNotImplemented)
	})
}
