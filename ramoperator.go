package db

import (
	"reflect"
)

type ramOperator struct {
}

func NewRAMOperator() Operator {
	return ramOperator{}
}

var memTables map[string][]IRecord

func init() {
	memTables = make(map[string][]IRecord)
}

func getTableName(obj interface{}) string {
	return reflect.TypeOf(obj).Elem().Name()
}

//Insert adds a new record and returns the ID created
func (opr ramOperator) Insert(obj interface{}) (int64, error) {
	key := getTableName(obj)
	rec, ok := obj.(IRecord)

	if ok {
		memTables[key] = append(memTables[key], rec)
	}

	return 99, nil
}

// InsertMulti adds multiple records, returns the number of rows saved
func (opr ramOperator) InsertMulti(batchCount int, objs interface{}) (int64, error) {
	return 99, nil
}

// ReadOne reads one record filtered by the values already specified in obj
func (opr ramOperator) ReadOne(obj interface{}, related ...string) error {
	key := getTableName(obj)

	for _, v := range memTables[key] {
		if obj.(IRecord).GetID() == v.GetID() {
			obj = v
			break
		}
	}

	return nil
}

// Read reads all records filtered by the values specified in the filter
func (opr ramOperator) Read(filter interface{}, container interface{}) error {
	//key := getTableName(filter)

	return nil
}

// Update will persist any changes made to a record
func (opr ramOperator) Update(obj interface{}) (int64, error) {
	key := getTableName(obj)
	rec, ok := obj.(IRecord)

	if ok {
		for k, v := range memTables[key] {
			if v.GetID() == rec.GetID() {
				memTables[key][k] = rec
				break
			}
		}
	}

	return 99, nil
}
