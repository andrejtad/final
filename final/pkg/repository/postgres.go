package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable                      = "users"
	dataOwnerTable                  = "data_owner"
	datasetTable                    = "dataset"
	specificationTable              = "specification"
	tagTable                        = "tag"
	linkTypeTable                   = "link_type"
	datasetTagTable                 = "dataset_tag"
	datasetSpecificationTable       = "dataset_specification"
	specificationSpecificationTable = "specification_specification"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
