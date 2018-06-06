package db

import (
	"reflect"
	"testing"
)

type fakeContext struct {
	TableA Setter
}

type tablea struct {
	Record
	Name        string
	Collections []*tablea
	Relation    *tablea
}

func (t tablea) Validate() (bool, error) {
	return true, nil
}

func TestFindModels_ReturnsModelTypes(t *testing.T) {
	ctx := fakeContext{
		TableA: NewDBSet(tablea{}),
	}

	actual, err := findModels(&ctx)

	t.Log(actual)

	if err != nil {
		t.Error(err)
	}

	for _, v := range actual {
		val := reflect.ValueOf(v)
		typ := reflect.Indirect(val).Type()

		if val.Kind() != reflect.Ptr {
			t.Error("value can't be a pointer")
		}

		if typ.Kind() == reflect.Ptr {
			t.Errorf("too many references to struct")
		}

		name := reflect.Indirect(val).Type().Name()

		if name != "tablea" {
			t.Errorf("Expected 'tablea', got %v", name)
		}
	}

	/*for _, v := range actual {
		val := reflect.ValueOf(v)
		kind := val.Kind()

		if kind != reflect.Ptr {
			t.Errorf("Expected %v, got %v", reflect.Ptr, kind)
		}

		_, ok := v.(IRecord)

		if !ok {
			t.Errorf("item %v is not IRecord", v)
		}
	}*/
}
