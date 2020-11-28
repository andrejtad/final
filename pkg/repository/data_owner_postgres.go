package repository

import (
	"fmt"
	"github.com/andrejtad/final"
	"github.com/jmoiron/sqlx"
	"strings"
)

type DataOwnerPostgres struct {
	db *sqlx.DB
}

func NewDataOwnerPostgres(db *sqlx.DB) *DataOwnerPostgres {
	return &DataOwnerPostgres{db: db}
}

func (r *DataOwnerPostgres) Create(dataOwner final.DataOwner) (int, error)  {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createDataOwnerQuery := fmt.Sprintf("INSERT INTO %s (title, link_to_logo) VALUES ($1, $2) RETURNING id", dataOwnerTable)
	row := tx.QueryRow(createDataOwnerQuery, dataOwner.Title, dataOwner.LinkToLogo)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *DataOwnerPostgres) GetAll()  ([]final.DataOwner, error) {
	var dataOwners []final.DataOwner
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.link_to_logo FROM %s tl", dataOwnerTable)
	err := r.db.Select(&dataOwners, query)
	return dataOwners, err
}

func (r *DataOwnerPostgres) GetById(dataOwnerId int) (final.DataOwner, error)  {
	var dataOwner final.DataOwner
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.link_to_logo FROM %s tl WHERE tl.id = $1", dataOwnerTable)
	err := r.db.Get(&dataOwner, query, dataOwnerId )
	return dataOwner, err
}

func (r *DataOwnerPostgres) Delete(dataOwnerId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl WHERE tl.id = $1", dataOwnerTable)
	_, err := r.db.Exec(query, dataOwnerId )
	return err
}

func (r *DataOwnerPostgres) Update(dataOwnerId int, input final.UpdateDataOwnerInput) error  {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil{
		setValues = append(setValues, fmt.Sprintf("title = $%d ", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.LinkToLogo != nil{
		setValues = append(setValues, fmt.Sprintf("link_to_logo = $%d ", argId))
		args = append(args, *input.LinkToLogo)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s  WHERE tl.id = $%d", dataOwnerTable, setQuery, argId)
	args = append(args, dataOwnerId)
	_, err := r.db.Exec(query, args... )
	return err
}
