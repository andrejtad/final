package final

import "errors"

type DataOwner struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	LinkToLogo  string `json:"linktologo" db:"link_to_logo"`
}

type Dataset struct {
	Id          int    `json:"id" db:"id"`
	DataOwnerId int    `db:"data_owner_id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type Specification struct {
	Id          int    `json:"id" db:"id"`
	DatasetId   int    `db:"dataset_id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	IsPrimary   string `json:"isprimary" db:"is_primary"`
	IsReference string `json:"isreference" db:"is_reference"`
}

type Tag struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
}

type Link struct {
	Id          int    `json:"id" db:"id"`
	ChildId     int    `json:"childid" db:"child_id"`
	ParentId    int    `json:"parentid" db:"parent_id"`
	LinkType    string `json:"linktype" db:"link_type"`
}

type LinkType struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
}


type UpdateDataOwnerInput struct {
	Title       *string `json:"title"`
	LinkToLogo  *string `json:"linktologo"`
}

type UpdateDatasetInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type UpdateSpecificationInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	IsPrimary   *bool   `json:"isprimary"`
	IsReference *bool   `json:"isreference"`
}

type UpdateTagInput struct {
	Title       *string `json:"title"`
}

func (i UpdateDataOwnerInput) Validate() error {
	if i.Title == nil && i.LinkToLogo == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

func (i UpdateDatasetInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

func (i UpdateSpecificationInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.IsPrimary == nil && i.IsReference == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

func (i UpdateTagInput) Validate() error {
	if i.Title == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
