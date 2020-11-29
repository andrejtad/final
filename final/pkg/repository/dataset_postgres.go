package repository

import (
	"fmt"
	"github.com/andrejtad/final"
	"github.com/jmoiron/sqlx"
	"strings"
)

type DatasetPostgres struct {
	db *sqlx.DB
}

func NewDatasetPostgres(db *sqlx.DB) *DatasetPostgres {
	return &DatasetPostgres{db: db}
}

func (r *DatasetPostgres) Create(dataOwnerId int, dataset final.Dataset) (int, error)  {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var datasetId int
	createDatasetQuery := fmt.Sprintf("INSERT INTO %s (data_owner_id, title, description) values ($1, $2, $3) RETURNING id", datasetTable)

	row := tx.QueryRow(createDatasetQuery, dataOwnerId, dataset.Title, dataset.Description)
	err = row.Scan(&datasetId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return datasetId, tx.Commit()
}

func (r *DatasetPostgres) GetAll(dataOwnerId int) ([]final.Dataset, error) {
	var datasets []final.Dataset
	query := fmt.Sprintf(`SELECT tds.id, tds.data_owner_id, tds.title, tds.description FROM %s tds WHERE tds.data_owner_id = $1`, datasetTable)
	if err := r.db.Select(&datasets, query, dataOwnerId); err != nil {
		return nil, err
	}

	return datasets, nil
}

func (r *DatasetPostgres) GetById(datasetId int) (final.Dataset, error) {
	var dataset final.Dataset
	query := fmt.Sprintf(`SELECT tds.id, tds.data_owner_id, tds.title, tds.description FROM %s tds WHERE tds.id = $1`, datasetTable)
	if err := r.db.Get(&dataset, query, datasetId); err != nil {
		return dataset, err
	}

	return dataset, nil
}

func (r *DatasetPostgres) Delete(datasetId int) error {
	query := fmt.Sprintf(`DELETE FROM %s tds WHERE tds.id = $1`, datasetTable)
	_, err := r.db.Exec(query, datasetId)
	return err
}

func (r *DatasetPostgres)Update(datasetId int, input final.UpdateDatasetInput) error {
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

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s tds SET %s WHERE tds.id = $%d`, datasetTable, setQuery, argId)
	args = append(args, datasetId)

	_, err := r.db.Exec(query, args...)
	return err
}