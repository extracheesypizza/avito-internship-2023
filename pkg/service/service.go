package service

import (
	"avito-app"
	"avito-app/pkg/repository"
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

type Service struct {
	User
	Segment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:    NewUserService(repos.User),
		Segment: NewSegmentService(repos.Segment),
	}
}
