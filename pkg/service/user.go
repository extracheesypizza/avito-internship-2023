package service

import (
	"avito-app"
	"avito-app/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) AddUserToSegment(usr avito.User) (int, error) {
	return s.repo.AddUserToSegment(usr)
}

func (s *UserService) DeleteUserFromSegment(usr avito.User) (int, error) {
	return s.repo.DeleteUserFromSegment(usr)
}

func (s *UserService) GetUserSegments(usr_id int) ([]string, error) {
	return s.repo.GetUserSegments(usr_id)
}
