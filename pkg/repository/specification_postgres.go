package repository

import (
	"fmt"
	"github.com/andrejtad/final"
	"github.com/jmoiron/sqlx"
	"strings"
)

type SpecificationPostgres struct {
	db *sqlx.DB
}

func NewSpecificationPostgres(db *sqlx.DB) *SpecificationPostgres {
	return &SpecificationPostgres{db: db}
}

func (r *SpecificationPostgres) Create(datasetId int, specification final.Specification) (int, error)  {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var specificationId int
	createSpecificationQuery := fmt.Sprintf("INSERT INTO %s (dataset_id, title, description) values ($1, $2, $3) RETURNING id", specificationTable)

	row := tx.QueryRow(createSpecificationQuery, datasetId, specification.Title, specification.Description)
	err = row.Scan(&specificationId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return specificationId, tx.Commit()
}

func (r *SpecificationPostgres) GetAll(datasetId int) ([]final.Specification, error) {
	var datasets []final.Specification
	query := fmt.Sprintf(`SELECT ts.id, ts.dataset_id, ts.title, ts.description, ts.is_primary, ts.is_reference FROM %s ts WHERE ts.dataset_id = $1`, specificationTable)
	if err := r.db.Select(&datasets, query, datasetId); err != nil {
		return nil, err
	}

	return datasets, nil
}

func (r *SpecificationPostgres) GetById(specificationsId int) (final.Specification, error) {
	var dataset final.Specification
	query := fmt.Sprintf(`SELECT tds.id, tds.data_owner_id, tds.title, tds.description FROM %s tds WHERE tds.id = $1`, specificationTable)
	if err := r.db.Get(&dataset, query, specificationsId); err != nil {
		return dataset, err
	}

	return dataset, nil
}

func (r *SpecificationPostgres) Delete(specificationsId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ts WHERE ts.id = $1`, specificationTable)
	_, err := r.db.Exec(query, specificationsId)
	return err
}

func (r *SpecificationPostgres)Update(specificationsId int, input final.UpdateSpecificationInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.IsPrimary != nil {
		setValues = append(setValues, fmt.Sprintf("is_primary=$%d", argId))
		args = append(args, *input.IsPrimary)
		argId++
	}

	if input.IsReference != nil {
		setValues = append(setValues, fmt.Sprintf("is_reference=$%d", argId))
		args = append(args, *input.IsReference)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ts SET %s WHERE ts.id = $%d`, specificationTable, setQuery, argId)
	args = append(args, specificationsId)

	_, err := r.db.Exec(query, args...)
	return err
}