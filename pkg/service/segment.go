package service

import (
	"avito-app"
	"avito-app/pkg/repository"
)

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) CreateSegment(seg avito.Segment) (int, error) {
	return s.repo.CreateSegment(seg)
}

func (s *SegmentService) DeleteSegment(seg avito.Segment) (int, error) {
	return s.repo.DeleteSegment(seg)
}
