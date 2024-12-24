package api_test

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/sasimpson/servicedemo/models"

	"github.com/sasimpson/servicedemo/api"
	"github.com/sasimpson/servicedemo/interfaces/mock"
)

var (
	testUUID1 = uuid.New()
	testUUID2 = uuid.New()
)

func TestUserAllHandler(t *testing.T) {
	testCases := []struct {
		desc         string
		handler      api.UserAPI
		responseCode int
		responseBody string
	}{
		{
			desc: "get no users",
			handler: api.UserAPI{
				BaseHandler: api.BaseHandler{
					User: &mock.User{
						Users: &[]models.User{},
					},
				},
			},
			responseCode: http.StatusOK,
			responseBody: "[]\n",
		},
		{
			desc: "get error",
			handler: api.UserAPI{
				BaseHandler: api.BaseHandler{
					User: &mock.User{
						Error: errors.New("unknown error"),
					},
				},
			},
			responseCode: http.StatusInternalServerError,
		},
		{
			desc: "get 2 users",
			handler: api.UserAPI{
				BaseHandler: api.BaseHandler{
					User: &mock.User{
						Users: &[]models.User{
							{
								ID:        testUUID1,
								FirstName: "Test",
								LastName:  "Tester",
								Birthday:  time.Date(1979, 1, 19, 0, 0, 0, 0, time.UTC),
							},
							{
								ID:        testUUID2,
								FirstName: "Tester",
								LastName:  "Testing",
								Birthday:  time.Date(1978, 2, 20, 0, 0, 0, 0, time.UTC),
							},
						},
					},
				},
			},
			responseCode: http.StatusOK,
			responseBody: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/user", nil)
			if err != nil {
				t.Errorf("handler returned an error when it shouldn't")
			}

			rr := httptest.NewRecorder()
			handler := tC.handler.UserAllHandler()

			handler.ServeHTTP(rr, req)
			if rr.Code != tC.responseCode {
				t.Errorf("return code does not match, wanted: %v got: %v", tC.responseCode, rr.Code)
			}

			if tC.responseBody != "" {
				if rr.Body.String() != tC.responseBody {
					t.Errorf("returned body does not match, wanted: %v got: %v", tC.responseBody, rr.Body.String())
				}
			}
		})
	}
}

func TestUserGetHandler(t *testing.T) {
	testCases := []struct {
		desc         string
		handler      api.UserAPI
		requestID    string
		responseCode int
		responseBody string
	}{
		{
			desc: "get user",
			handler: api.UserAPI{
				BaseHandler: api.BaseHandler{
					User: &mock.User{
						User: &models.User{
							ID:        testUUID1,
							FirstName: "test",
							LastName:  "user",
							Birthday:  time.Date(1978, 2, 20, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
			requestID:    testUUID1.String(),
			responseCode: http.StatusOK,
		},
		{
			desc: "get none",
			handler: api.UserAPI{
				BaseHandler: api.BaseHandler{
					User: &mock.User{
						Error: models.ErrNotFound,
					},
				},
			},
			requestID:    testUUID1.String(),
			responseCode: http.StatusNotFound,
		},
		{
			desc: "get no id",
			handler: api.UserAPI{
				BaseHandler: api.BaseHandler{
					User: &mock.User{},
				},
			},
			requestID:    "",
			responseCode: http.StatusBadRequest,
		},
		{
			desc: "get bad id",
			handler: api.UserAPI{
				BaseHandler: api.BaseHandler{
					User: &mock.User{
						User: nil,
					},
				},
			},
			requestID:    "abc",
			responseCode: http.StatusBadRequest,
		},
		{
			desc: "get error",
			handler: api.UserAPI{
				BaseHandler: api.BaseHandler{
					User: &mock.User{
						Error: errors.New("unknown error"),
					},
				},
			},
			requestID:    testUUID1.String(),
			responseCode: http.StatusInternalServerError,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req, _ := http.NewRequest("GET", fmt.Sprintf("/user/%s", tC.requestID), nil)
			req = mux.SetURLVars(req, map[string]string{"id": tC.requestID})
			rr := httptest.NewRecorder()
			tC.handler.UserGetHandler().ServeHTTP(rr, req)
			if rr.Code != tC.responseCode {
				t.Errorf("return code does not match, wanted: %v got: %v", tC.responseCode, rr.Code)
			}
		})
	}
}

func TestUserPostHandler(t *testing.T) {
	testCases := []struct {
		desc         string
		userRequest  string
		handler      api.UserAPI
		responseCode int
	}{
		{
			desc: "add new user",
			userRequest: `
				{
					"first_name": "Test",
					"last_name": "User",
					"birthday": "2010-04-23T18:25:43.511Z",
					"email": "testuser@test.com"
				}
			`,
			handler: api.UserAPI{
				BaseHandler: api.BaseHandler{
					User: &mock.User{
						User: &models.User{
							ID:        testUUID1,
							FirstName: "Test",
							LastName:  "User",
							Birthday:  time.Date(2010, 4, 23, 18, 25, 43, 511, time.UTC),
							Email:     "testuser@test.com",
						},
					},
				},
			},
			responseCode: http.StatusCreated,
		},
		{
			desc: "new user bad json",
			userRequest: `
				{
					"first_name": "Test",
					"last_name": "User",
					"birthday": "2010-04-23T18:25:43.511Z",
					"email": 123
				}
			`,
			handler: api.UserAPI{
				BaseHandler: api.BaseHandler{
					User: &mock.User{},
				},
			},
			responseCode: http.StatusBadRequest,
		},
		{
			desc: "new user error",
			userRequest: `
				{
					"first_name": "Test",
					"last_name": "User",
					"birthday": "2010-04-23T18:25:43.511Z",
					"email": "testuser@test.com"
				}
			`,
			handler: api.UserAPI{
				BaseHandler: api.BaseHandler{
					User: &mock.User{
						Error: models.ErrAlreadyExists,
					},
				},
			},
			responseCode: http.StatusBadRequest,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			requestBody := strings.NewReader(tC.userRequest)

			req, err := http.NewRequest("POST", "/user", requestBody)
			if err != nil {
				t.Errorf("handler returned an error when it shouldn't")
			}

			rr := httptest.NewRecorder()
			handler := tC.handler.UserNewHandler()

			handler.ServeHTTP(rr, req)
			if rr.Code != tC.responseCode {
				t.Errorf("return code does not match, wanted: %v got: %v", tC.responseCode, rr.Code)
			}
		})
	}
}

func TestUserRoutes(t *testing.T) {
	testCases := []struct {
		desc    string
		name    string
		path    string
		methods []string
	}{
		{
			desc:    "all users",
			name:    "UserGetAll",
			path:    "/user",
			methods: []string{"GET"},
		},
		{
			desc:    "get user by id",
			name:    "UserGetByID",
			path:    "/user/{id}",
			methods: []string{"GET"},
		},
	}

	h := api.UserAPI{}
	routes := mux.NewRouter()
	h.InitRoutes(routes)

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			route := routes.Get(tC.name)
			path, _ := route.GetPathTemplate()
			// methods, _ := route.GetMethods()
			if tC.name != route.GetName() {
				t.Errorf("route name does not match, expected: %v and got: %v", tC.name, route.GetName())
			}
			if tC.path != path {
				t.Errorf("path does not match, expected: %v got: %v", tC.path, path)
			}
			// if tC.methods != methods {
			// 	t.Errorf("methods do not match, expected: %v got: %v", tC.methods, methods)
			// }
		})
	}
}
