package api

import "github.com/sasimpson/servicedemo/models"

//Handler contains the model interface implementations.
type Handler struct {
	Env  models.Env
	User models.UserModel
}
