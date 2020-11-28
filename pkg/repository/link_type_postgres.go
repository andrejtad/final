package repository

import (
	"fmt"
	"github.com/andrejtad/final"
	"github.com/jmoiron/sqlx"
)

type LinkTypePostgres struct {
	db *sqlx.DB
}

func NewLinkTypePostgres(db *sqlx.DB) *LinkTypePostgres {
	return &LinkTypePostgres{db: db}
}

func (r *LinkTypePostgres) GetAll()  ([]final.LinkType, error) {
	var linkTypes []final.LinkType
	query := fmt.Sprintf("SELECT lt.id, lt.title FROM %s lt", linkTypeTable)
	err := r.db.Select(&linkTypes, query)
	return linkTypes, err
}

