package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sasimpson/servicedemo/api"
	"github.com/sasimpson/servicedemo/interfaces/mock"
	"github.com/stretchr/testify/assert"
)

func TestUserAllHandler(t *testing.T) {
	testCases := []struct {
		desc    string
		handler api.Handler
	}{
		{
			desc: "",
			handler: api.Handler{
				User: &mock.UserMock{},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert := assert.New(t)
			req, err := http.NewRequest("GET", "/user", nil)
			assert.Nil(err)

			rr := httptest.NewRecorder()
			handler := tC.handler.UserAllHandler()

			handler.ServeHTTP(rr, req)
			assert.Equal(http.StatusNotImplemented, rr.Code)
		})
	}
}

func TestUserGetHandler(t *testing.T) {
	testCases := []struct {
		desc    string
		handler api.Handler
	}{
		{
			desc: "",
			handler: api.Handler{
				User: &mock.UserMock{},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert := assert.New(t)
			req, err := http.NewRequest("GET", "/user/1", nil)
			assert.Nil(err)

			rr := httptest.NewRecorder()
			handler := tC.handler.UserGetHandler()

			handler.ServeHTTP(rr, req)
			assert.Equal(http.StatusNotImplemented, rr.Code)
		})
	}
}
