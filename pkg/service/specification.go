package service

import (
	"github.com/andrejtad/final"
	"github.com/andrejtad/final/pkg/repository"
)

type SpecificationService struct {
	repo repository.Specification
	datasetRepo repository.Dataset
}

func NewSpecificationService(repo repository.Specification, datasetRepo repository.Dataset) *SpecificationService {
	return &SpecificationService{repo: repo, datasetRepo: datasetRepo}
}

func (s *SpecificationService) Create(datasetId int, specification final.Specification) (int, error) {
	_, err := s.datasetRepo.GetById(datasetId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(datasetId, specification)
}

func (s *SpecificationService) GetAll(datasetId int) ([]final.Specification, error)  {
	return s.repo.GetAll(datasetId)
}


func (s *SpecificationService) GetById(specificationsId int) (final.Specification, error) {
	return s.repo.GetById(specificationsId)
}

func (s *SpecificationService) Delete(specificationsId int) error {
	return s.repo.Delete(specificationsId)
}

func (s *SpecificationService) Update(specificationsId int, input final.UpdateSpecificationInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(specificationsId, input)
}

