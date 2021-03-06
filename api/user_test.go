package api_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/sasimpson/servicedemo/models"

	"github.com/sasimpson/servicedemo/api"
	"github.com/sasimpson/servicedemo/interfaces/mock"
)

func TestUserAllHandler(t *testing.T) {
	testCases := []struct {
		desc         string
		handler      api.Handler
		responseCode int
		responseBody string
	}{
		{
			desc: "get no users",
			handler: api.Handler{
				User: &mock.UserMock{
					Users: &[]models.User{},
				},
			},
			responseCode: http.StatusOK,
			responseBody: "[]\n",
		},
		{
			desc: "get error",
			handler: api.Handler{
				User: &mock.UserMock{
					Error: errors.New("Unknown Error"),
				},
			},
			responseCode: http.StatusInternalServerError,
		},
		{
			desc: "get 2 users",
			handler: api.Handler{
				User: &mock.UserMock{
					Users: &[]models.User{
						models.User{
							ID:        1,
							FirstName: "Test",
							LastName:  "Tester",
							Birthday:  time.Date(1979, 1, 19, 0, 0, 0, 0, time.UTC),
						},
						models.User{
							ID:        2,
							FirstName: "Tester",
							LastName:  "Testing",
							Birthday:  time.Date(1978, 2, 20, 0, 0, 0, 0, time.UTC),
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
		handler      api.Handler
		responseCode int
	}{
		{
			desc: "",
			handler: api.Handler{
				User: &mock.UserMock{},
			},
			responseCode: http.StatusNotImplemented,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/user/1", nil)
			if err != nil {
				t.Errorf("handler returned an error when it shouldn't")
			}

			rr := httptest.NewRecorder()
			handler := tC.handler.UserGetHandler()

			handler.ServeHTTP(rr, req)
			if rr.Code != tC.responseCode {
				t.Errorf("return code does not match, wanted: %v got: %v", tC.responseCode, rr.Code)
			}
		})
	}
}

func TestUserPostHandler(t *testing.T) {
	testCases := []struct {
		desc         string
		handler      api.Handler
		responseCode int
	}{
		{
			desc: "",
			handler: api.Handler{
				User: &mock.UserMock{},
			},
			responseCode: http.StatusNotImplemented,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/user", nil)
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
		route   string
		methods []string
	}{
		{
			desc: "get ",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}
