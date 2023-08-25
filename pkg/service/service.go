package service

import "avito-app/pkg/repository"

type UserList interface {
}

type Service struct {
	UserList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
