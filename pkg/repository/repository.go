package repository

import "github.com/jmoiron/sqlx"

type UserList interface {
}

type Repository struct {
	UserList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
