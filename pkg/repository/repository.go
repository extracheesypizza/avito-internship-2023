package repository

import (
	"avito-app"

	"github.com/jmoiron/sqlx"
)

type Segment interface {
	CreateSegment(seg avito.Segment) (int, error)
	DeleteSegment(seg avito.Segment) (int, error)
}

type User interface {
	GetUserSegments(usr_id int) ([]string, error)
	AddUserToSegment(usr avito.User) (int, error)
	DeleteUserFromSegment(usr avito.User) (int, error)
}

type Repository struct {
	User
	Segment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUserPostgres(db),
		Segment: NewSegmentPostgres(db),
	}
}
