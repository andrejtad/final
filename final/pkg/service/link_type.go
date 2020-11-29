package service

import (
	"github.com/andrejtad/final"
	"github.com/andrejtad/final/pkg/repository"
)

type LinkTypeService struct {
	repo repository.LinkType
}

func NewLinkTypeService(repo repository.LinkType) *LinkTypeService {
	return &LinkTypeService{repo: repo}
}

func (s *LinkTypeService) GetAll() ([]final.LinkType, error)  {
	return s.repo.GetAll()
}
