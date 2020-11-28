package service

import (
	"github.com/andrejtad/final"
	"github.com/andrejtad/final/pkg/repository"
)

type Authorisation interface {
	CreateUser(user final.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type DataOwner interface {
	Create(dataOwner final.DataOwner) (int, error)
	GetAll() ([]final.DataOwner, error)
	GetById(dataOwnerId int) (final.DataOwner, error)
	Delete(dataOwnerId int) error
	Update(dataOwnerId int, input final.UpdateDataOwnerInput) error
}

type Dataset interface {
	Create(dataOwnerId int, dataset final.Dataset) (int, error)
	GetAll(dataOwnerId int) ([]final.Dataset, error)
	GetById(datasetId int) (final.Dataset, error)
	Delete(datasetId int) error
	Update(datasetId int, input final.UpdateDatasetInput) error
}

type Specification  interface {
	Create(datasetId int, dataset final.Specification) (int, error)
	GetAll(datasetId int) ([]final.Specification, error)
	GetById(specificationsId int) (final.Specification, error)
	Delete(specificationsId int) error
	Update(specificationsId int, input final.UpdateSpecificationInput) error
}

type Service struct {
	Authorisation
	DataOwner
	Dataset
	Specification
}

func NewService(repos *repository.Repository) *Service  {
	return &Service{
		Authorisation: NewAuthService(repos.Authorisation),
		DataOwner: NewDataOwnerService(repos.DataOwner),
		Dataset: NewDatasetService(repos.Dataset, repos.DataOwner),
		Specification: NewSpecificationService(repos.Specification, repos.Dataset),
	}
}