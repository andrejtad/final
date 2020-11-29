package repository

import (
	"fmt"
	"github.com/andrejtad/final"
	"github.com/jmoiron/sqlx"
)

type LinkPostgres struct {
	db *sqlx.DB
}

func NewLinkPostgres(db *sqlx.DB) *LinkPostgres {
	return &LinkPostgres{db: db}
}

func (r *LinkPostgres) GetAll(childId int)  ([]final.Link, error) {
	var links []final.Link
	query := fmt.Sprintf("SELECT l.id, l.parent_id, l.child_id, lt.title as link_type FROM %s l, %s lt where l.link_type = lt.id and l.child_id = $1", specificationSpecificationTable, linkTypeTable)
	err := r.db.Select(&links, query, childId)
	return links, err
}

