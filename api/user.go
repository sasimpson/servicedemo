package api

import (
	"encoding/json"
	"errors"
	"github.com/swaggest/openapi-go/openapi3"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sasimpson/servicedemo/models"
)

// UserAPI -
type UserAPI struct {
	BaseHandler
}

func (api *UserAPI) InitOpenAPI(r *openapi3.Reflector) {
	//Paths:
	// GET /users
	getAllOp := openapi3.Operation{}
	getAllOp.WithDescription("Get all users")
	r.SetRequest(&getAllOp, nil, http.MethodGet)
	r.SetJSONResponse(&getAllOp, new([]models.User), http.StatusOK)
	r.SetJSONResponse(&getAllOp, new([]models.User), http.StatusInternalServerError)
	r.Spec.AddOperation(http.MethodGet, "/user", getAllOp)

	// GET /users/:id
	type getReq struct {
		ID string `path:"id"`
	}
	getOp := openapi3.Operation{}
	getOp.WithDescription("Get user by id")
	r.SetRequest(&getOp, new(getReq), http.MethodGet)
	r.SetJSONResponse(&getOp, new(models.User), http.StatusOK)
	r.SetJSONResponse(&getOp, nil, http.StatusBadRequest)
	r.SetJSONResponse(&getOp, nil, http.StatusNotFound)
	r.SetJSONResponse(&getOp, nil, http.StatusInternalServerError)
	err := r.Spec.AddOperation(http.MethodGet, "/user/{id}", getOp)
	if err != nil {
		panic(err)
	}
}

// InitRoutes will initialize the routes for the user API endpoint
func (api *UserAPI) InitRoutes(r *mux.Router) {
	sr := r.PathPrefix("/user").Subrouter()
	sr.Handle("", api.UserAllHandler()).Methods("GET").Name("UserGetAll")
	sr.Handle("/{id}", api.UserGetHandler()).Methods("GET").Name("UserGetByID")
}

// UserAllHandler will return the data for a single user record
func (api *UserAPI) UserAllHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var users *[]models.User
		users, err := api.User.All()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

// UserGetHandler will return the data for a single user record
func (api *BaseHandler) UserGetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idVar := mux.Vars(r)["id"]

		if idVar == "" {
			http.Error(w, "no id for user", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idVar)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := api.User.Get(id)
		if err != nil {
			if errors.Is(err, models.ErrNotFound) {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

// UserNewHandler will return the data for a single user record
func (api *BaseHandler) UserNewHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		returnedUser, err := api.User.New(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(returnedUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
