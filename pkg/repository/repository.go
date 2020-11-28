package repository

import (
	"github.com/andrejtad/final"
	"github.com/jmoiron/sqlx"
)

type Authorisation interface {
	CreateUser(user final.User) (int, error)
	GetUser(username, password string) (final.User, error)
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

type Specification interface {
	Create(datasetId int, dataset final.Specification) (int, error)
	GetAll(datasetId int) ([]final.Specification, error)
	GetById(specificationsId int) (final.Specification, error)
	Delete(specificationsId int) error
	Update(specificationsId int, input final.UpdateSpecificationInput) error
}

type Repository struct {
	Authorisation
	DataOwner
	Dataset
	Specification
}

func NewRepository(db *sqlx.DB) *Repository  {
	return &Repository{
		Authorisation: NewAuthPostgres(db),
		DataOwner: NewDataOwnerPostgres(db),
		Dataset: NewDatasetPostgres(db),
		Specification: NewSpecificationPostgres(db),
	}
}
