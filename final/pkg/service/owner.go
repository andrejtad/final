package service

import (
	"github.com/andrejtad/final"
	"github.com/andrejtad/final/pkg/repository"
)

type DataOwnerService struct {
	repo repository.DataOwner
}

func NewDataOwnerService(repo repository.DataOwner) *DataOwnerService {
	return &DataOwnerService{repo: repo}
}

func (s *DataOwnerService) Create(dataOwner final.DataOwner) (int, error) {
	return s.repo.Create(dataOwner)
}

func (s *DataOwnerService) GetAll() ([]final.DataOwner, error)  {
	return s.repo.GetAll()
}

func (s *DataOwnerService) GetById(dataOwnerId int) (final.DataOwner, error) {
	return s.repo.GetById(dataOwnerId)
}

func (s *DataOwnerService) Delete(dataOwnerId int) error  {
	return s.repo.Delete(dataOwnerId)
}

func (s *DataOwnerService) Update(dataOwnerId int, input final.UpdateDataOwnerInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(dataOwnerId, input)
}
