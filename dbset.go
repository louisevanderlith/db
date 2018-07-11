package db

import (
	"errors"
	"fmt"
	"reflect"
)

type dbSet struct {
	t reflect.Type
}

// NewDBSet returns a Setter for real database sets
func NewDBSet(t IRecord) Setter {
	result := new(dbSet)
	result.t = reflect.TypeOf(t)

	return result
}

// GetRecordType returns the Type of the IRecord it holds
func (s *dbSet) GetRecordType() reflect.Type {
	return s.t
}

func (set *dbSet) Create(item IRecord) (id int64, err error) {
	if reflect.ValueOf(item).Kind() != reflect.Ptr {
		return 0, errors.New("IRecord must be a pointer")
	}

	if elem := reflect.TypeOf(item).Elem(); elem == set.t {
		id, err = tryCreate(item)
	} else {
		msg := fmt.Sprintf("%s is not of type %s", elem, set.t)
		err = errors.New(msg)
	}

	return id, err
}

func tryCreate(item IRecord) (id int64, err error) {
	var valid bool
	valid, err = item.IsValid()

	if valid {
		var exists bool
		exists, err = item.Exists()

		if !exists {
			id, err = operator.Insert(item)
		}
	}

	return id, err
}

func (set *dbSet) CreateMulti(items IRecordCollection) (insertCount int64, err error) {
	count := items.Length()

	return operator.InsertMulti(count, items)
}

func (set *dbSet) ReadOne(filter IRecord, related ...string) (IRecord, error) {
	err := operator.ReadOne(filter, related...)

	return filter, err
}

func (set *dbSet) Read(filter IRecord, container IRecordCollection) error {
	err := operator.Read(filter, container)

	return err
}

func (set *dbSet) Fetch(page, pageSize int, expr func(obj IRecord) bool) (IRecordCollection, error) {
	var preContainer IRecordCollection
	var postContainer IRecordCollection
	tbl := reflect.New(set.GetRecordType())

	err := operator.Read(tbl, preContainer)

	if err == nil {
		skip := (page - 1) * pageSize

		for i := 0; i < preContainer.Length(); i++ {
			obj := preContainer.At(i)

			if expr(obj) && i > skip {
				if postContainer.Length() == pageSize {
					break
				}

				postContainer.Add(obj)
				//postContainer = append(postContainer, obj)
			}
		}
	}

	return postContainer, err
}

func (set *dbSet) Update(item IRecord) (err error) {
	var valid bool

	if valid, err = item.Validate(); valid {
		_, err = operator.Update(item)
	}

	return err
}

func (set *dbSet) Delete(item IRecord) error {
	_, err := operator.Update(item.Disable())

	return err
}
