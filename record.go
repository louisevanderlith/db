package db

import "time"

// IRecord provides functionality to datasets on a row level
type IRecord interface {
	// GetID gets the primary key value of a record
	GetID() int64

	// IsDeleted indicates if a record has been soft-deleted
	IsDeleted() bool

	// GetCreateDate gets the create data of a record
	GetCreateDate() time.Time

	// Disable will set the currect record as deleted
	Disable() Record

	// Validate checks that all validation rules pass before commiting changes
	Validate() (bool, error)

	// Exists indicates whether a similiar record exists
	Exists() (bool, error)
}

// Record is a base class providing ID, CreateDate, & Deleted columns to all tables
type Record struct {
	Id         int64     `orm:"auto;pk;unique;"`
	CreateDate time.Time `orm:"auto_now_add"`
	Deleted    bool      `orm:"default(false)"`
}

func (r Record) GetID() int64 {
	return r.Id
}

func (r Record) IsDeleted() bool {
	return r.Deleted
}

func (r Record) GetCreateDate() time.Time {
	return r.CreateDate
}

func (r Record) Disable() Record {
	r.Deleted = true

	return r
}

func (r Record) Exists() (bool, error) {
	return false, nil
}
