package main

import (
	"fmt"
	"github.com/sasimpson/servicedemo/models"
	"github.com/swaggest/assertjson"
	"github.com/swaggest/openapi-go/openapi3"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sasimpson/servicedemo/api"
	"github.com/sasimpson/servicedemo/interfaces/mock"
)

// main will assemble and start our service as well as handle all the configuration bits
func main() {
	userAPI := api.UserAPI{
		BaseHandler: api.BaseHandler{
			User: &mock.UserMock{
				User: &models.User{
					ID:        1,
					FirstName: "Bob",
					LastName:  "Demo",
					Email:     "demo.bob@demo.com",
					Birthday:  time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				Users: &[]models.User{
					{
						ID:        1,
						FirstName: "Bob",
						LastName:  "Demo",
						Email:     "demo.bob@demo.com",
						Birthday:  time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				},
			},
		},
	}

	reflector := openapi3.Reflector{}
	reflector.Spec = &openapi3.Spec{Openapi: "3.0.3"}
	reflector.Spec.Info.
		WithTitle("Service Demo").
		WithDescription("This is a demo of one way to build services in Go.").
		WithVersion("v0.0.1")

	routes := mux.NewRouter()

	userAPI.InitRoutes(routes)
	userAPI.InitOpenAPI(&reflector)

	routes.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	routes.HandleFunc("/spec", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		j, err := assertjson.MarshalIndentCompact(reflector.Spec, "", " ", 120)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})

	log.Fatal(http.ListenAndServe(":8080", routes))
}
