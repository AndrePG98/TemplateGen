package service

import (
	"database/sql"
	// . "github.com/go-jet/jet/v2/mysql"
)

type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		DB: db,
	}
}
