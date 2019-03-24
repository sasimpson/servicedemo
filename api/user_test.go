package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sasimpson/servicedemo/api"
	"github.com/sasimpson/servicedemo/interfaces/mock"
)

func TestUserAllHandler(t *testing.T) {
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
