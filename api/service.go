package api

import "github.com/sasimpson/servicedemo/models"

// BaseHandler contains the model interface implementations.
type BaseHandler struct {
	Env  models.Env
	User models.UserDataInterface
}
