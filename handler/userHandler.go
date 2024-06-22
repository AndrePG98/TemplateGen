package handler

import (
	"TemplateGen/config"
	"TemplateGen/service"
	"database/sql"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(config *config.Config, db *sql.DB) *UserHandler {
	return &UserHandler{
		UserService: service.NewUserService(db),
	}
}
