package api

import "net/http"

//UserAPI -
type UserAPI struct {
}

//UserAllHandler will return the data for a single user record
func (h *Handler) UserAllHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Implemented", http.StatusNotImplemented)
	})
}

//UserGetHandler will return the data for a single user record
func (h *Handler) UserGetHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Implemented", http.StatusNotImplemented)
	})
}
