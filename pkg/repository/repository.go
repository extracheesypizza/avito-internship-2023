package repository

type UserList interface {
}

type Repository struct {
	UserList
}

func NewRepository() *Repository {
	return &Repository{}
}
