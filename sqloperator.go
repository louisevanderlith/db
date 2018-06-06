package db

import (
	"strings"

	"github.com/astaxie/beego/orm"
)

type sqlOperator struct {
}

func NewSQLOperator() Operator {
	return sqlOperator{}
}

//Insert adds a new record and returns the ID created
func (opr sqlOperator) Insert(obj interface{}) (int64, error) {
	o := orm.NewOrm()

	relationships := getRelationships(obj)

	id, err := o.Insert(obj)

	if err == nil {
		for _, v := range relationships {
			o.Insert(v)
		}
	}

	return id, err
}

// InsertMulti adds multiple records, returns the number of rows saved
func (opr sqlOperator) InsertMulti(batchCount int, objs interface{}) (int64, error) {
	o := orm.NewOrm()

	return o.InsertMulti(batchCount, objs)
}

// ReadOne reads one record filtered by the values already specified in obj
func (opr sqlOperator) ReadOne(obj interface{}, related ...string) error {
	readColumns := getReadColumns(obj)

	o := orm.NewOrm()
	err := o.Read(obj, readColumns...)

	for _, v := range related {
		o.LoadRelated(obj, v)
	}

	return err
}

// Read reads all records filtered by the values specified in the filter
func (opr sqlOperator) Read(filter interface{}, container interface{}) error {
	filterVals := getFilterValues(filter)

	o := orm.NewOrm()
	qt := o.QueryTable(filter)

	for k, v := range filterVals {
		qt = qt.Filter(strings.ToLower(k), v)
	}

	qt = qt.Filter("deleted", false)
	_, err := qt.All(container)

	return err
}

// Update will persist any changes made to a record
func (opr sqlOperator) Update(obj interface{}) (int64, error) {
	o := orm.NewOrm()
	relationships := getRelationships(obj)

	for _, v := range relationships {
		o.Update(v, getReadColumns(v)...)
	}

	return o.Update(obj, getReadColumns(obj)...)
}
