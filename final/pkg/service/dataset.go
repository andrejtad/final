package service

import (
	"github.com/andrejtad/final"
	"github.com/andrejtad/final/pkg/repository"
)

type DatasetService struct {
	repo repository.Dataset
	dataOwnerRepo repository.DataOwner
}

func NewDatasetService(repo repository.Dataset, dataOwnerRepo repository.DataOwner) *DatasetService {
	return &DatasetService{repo: repo, dataOwnerRepo: dataOwnerRepo}
}

func (s *DatasetService) Create(dataOwnerId int, dataset final.Dataset) (int, error) {
    _, err := s.dataOwnerRepo.GetById(dataOwnerId)
    if err != nil {
    	return 0, err
	}
	return s.repo.Create(dataOwnerId, dataset)
}

func (s *DatasetService) GetAll(datasetId int) ([]final.Dataset, error)  {
	return s.repo.GetAll(datasetId)
}


func (s *DatasetService) GetById(datasetId int) (final.Dataset, error) {
	return s.repo.GetById(datasetId)
}

func (s *DatasetService) Delete(datasetId int) error {
	return s.repo.Delete(datasetId)
}

func (s *DatasetService) Update(datasetId int, input final.UpdateDatasetInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(datasetId, input)
}
