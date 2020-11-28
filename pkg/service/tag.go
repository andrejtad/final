package service

import (
	"github.com/andrejtad/final"
	"github.com/andrejtad/final/pkg/repository"
)

type TagService struct {
	repo repository.Tag
}

func NewTagService(repo repository.Tag) *TagService {
	return &TagService{repo: repo}
}

func (s *TagService) Create(tag final.Tag) (int, error) {
	return s.repo.Create(tag)
}

func (s *TagService) GetAll() ([]final.Tag, error)  {
	return s.repo.GetAll()
}

func (s *TagService) GetById(tagId int) (final.Tag, error) {
	return s.repo.GetById(tagId)
}

func (s *TagService) Delete(tagId int) error  {
	return s.repo.Delete(tagId)
}

func (s *TagService) Update(tagId int, input final.UpdateTagInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(tagId, input)
}
