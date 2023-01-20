package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sasimpson/servicedemo/api"
	"github.com/sasimpson/servicedemo/interfaces/mock"
)

// main will assemble and start our service as well as handle all the configuration bits
func main() {

	userAPI := api.UserAPI{
		BaseHandler: api.BaseHandler{
			User: &mock.UserMock{},
		},
	}

	routes := mux.NewRouter()
	userAPI.InitRoutes(routes)

	log.Fatal(http.ListenAndServe(":8080", routes))
}
