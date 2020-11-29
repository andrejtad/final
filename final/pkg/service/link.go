package service

import (
	"github.com/andrejtad/final"
	"github.com/andrejtad/final/pkg/repository"
)

type LinkService struct {
	repo repository.Link
}

func NewLinkService(repo repository.Link) *LinkService {
	return &LinkService{repo: repo}
}

func (s *LinkService) GetAll(childId int) ([]final.Link, error)  {
	return s.repo.GetAll(childId)
}
