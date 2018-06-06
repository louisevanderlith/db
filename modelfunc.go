package db

import (
	"reflect"
	"time"
)

// getRelationships returns the objects related to the current object
func getRelationships(obj interface{}) []interface{} {
	var result []interface{}

	val := reflect.ValueOf(obj).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		if typeField.Type.Kind() == reflect.Ptr || typeField.Type.Kind() == reflect.Slice {
			value := valueField.Interface()

			if !valueField.IsNil() && value != nil {
				switch reflect.TypeOf(value).Kind() {
				case reflect.Slice:
					s := reflect.ValueOf(value)

					for j := 0; j < s.Len(); j++ {
						result = append(result, s.Index(j).Interface())
					}
				default:
					result = append(result, value)
				}
			}
		}
	}

	return result
}

// getReadColumns returns a list of column names, used to search
func getReadColumns(obj interface{}) []string {
	var result []string

	val := reflect.ValueOf(obj).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		value := valueField.Interface()
		kind := valueField.Kind()

		if !isFieldEmpty(value, kind) {
			result = append(result, typeField.Name)
		}
	}

	return result
}

// getFilterValues returns a map of parameters and their values, used to filter.
func getFilterValues(filter interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	val := reflect.ValueOf(filter).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		value := valueField.Interface()
		kind := valueField.Kind()

		if !isFieldEmpty(value, kind) {
			result[typeField.Name] = value
		}
	}

	return result
}

func isFieldEmpty(val interface{}, kind reflect.Kind) bool {
	var result bool

	switch kind {
	case reflect.Int:
		iField := val.(int)
		result = intRule(iField)
	case reflect.Int64:
		i64Field := val.(int64)
		result = int64Rule(i64Field)
	case reflect.String:
		strField := val.(string)
		result = strRule(strField)
	case reflect.Ptr, reflect.Struct, reflect.Slice:
		result = true
	case reflect.Bool:
		result = boolRule(val)
	default:
		result = nilRule(val)
	}

	if tField, ok := val.(time.Time); ok {
		result = tField.IsZero()
	}

	return result
}

func sliceRule(val interface{}) bool {
	records := reflect.ValueOf(val)
	return records.Len() <= 0
}

func nilRule(val interface{}) bool {
	return val == nil
}

func strRule(val interface{}) bool {
	return val == ""
}

func intRule(val int) bool {
	return val < 1
}

func int64Rule(val int64) bool {
	return val < 1
}

func boolRule(val interface{}) bool {
	return val == false
}

func addToMap(target map[string]interface{}, items map[string]interface{}) {
	for k, v := range items {
		target[k] = v
	}
}
