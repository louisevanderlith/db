package db

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/astaxie/beego/orm"
)

var operator Operator

// Sync will attempt to initialize tables and setup the database connection
func SyncDatabase(context interface{}, dbSource string) error {
	var err error

	if len(dbSource) < 6 && dbSource != "memory" {
		err = errors.New("dbSource is invalid")
		return err
	}

	if dbSource == "memory" {
		operator = NewRAMOperator()
	} else {
		operator = NewSQLOperator()
		err = spinUpORM(context, dbSource)
	}

	return err
}

func spinUpORM(context interface{}, dbSource string) error {
	err := registerModels(context)

	if err == nil {
		name := "default"
		driverName := "postgres"
		err = orm.RegisterDataBase(name, driverName, dbSource)

		if err == nil {
			err = orm.RunSyncdb(name, false, false)
		}
	}

	return err
}

func registerModels(context interface{}) error {
	models, err := findModels(context)

	if err == nil {
		orm.RegisterModel(models...)
	}

	return err
}

func findModels(context interface{}) (models []interface{}, err error) {
	val := reflect.ValueOf(context)

	if val.Kind() != reflect.Ptr {
		return nil, errors.New("context must be a pointer")
	}

	elem := val.Elem()

	for i := 0; i < elem.NumField(); i++ {
		valueField := elem.Field(i)
		value := valueField.Interface()
		setter, ok := value.(Setter)

		if ok {
			recType := setter.GetRecordType()
			inst := reflect.New(recType).Interface()

			models = append(models, inst)
		} else {
			msg := fmt.Sprintf("field %v is not a Setter", i)
			err = errors.New(msg)
		}
	}

	return models, err
}
