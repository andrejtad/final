package repository

import (
	"fmt"
	"github.com/andrejtad/final"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TagPostgres struct {
	db *sqlx.DB
}

func NewTagPostgres(db *sqlx.DB) *TagPostgres {
	return &TagPostgres{db: db}
}

func (r *TagPostgres) Create(tag final.Tag) (int, error)  {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createTagQuery := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1) RETURNING id", tagTable)
	row := tx.QueryRow(createTagQuery, tag.Title)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *TagPostgres) GetAll()  ([]final.Tag, error) {
	var tags []final.Tag
	query := fmt.Sprintf("SELECT tg.id, tg.title FROM %s tg", tagTable)
	err := r.db.Select(&tags, query)
	return tags, err
}

func (r *TagPostgres) GetById(tagId int) (final.Tag, error)  {
	var tag final.Tag
	query := fmt.Sprintf("SELECT tg.id, tg.title FROM %s tg WHERE tg.id = $1", tagTable)
	err := r.db.Get(&tag, query, tagId)
	return tag, err
}

func (r *TagPostgres) Delete(tagId int) error {
	query := fmt.Sprintf("DELETE FROM %s tg WHERE tg.id = $1", tagTable)
	_, err := r.db.Exec(query, tagId)
	return err
}

func (r *TagPostgres) Update(tagId int, input final.UpdateTagInput) error  {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil{
		setValues = append(setValues, fmt.Sprintf("title = $%d ", argId))
		args = append(args, *input.Title)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tg SET %s  WHERE tg.id = $%d", tagTable, setQuery, argId)
	args = append(args, tagId)
	_, err := r.db.Exec(query, args... )
	return err
}
