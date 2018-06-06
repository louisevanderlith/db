package db

import (
	"testing"
)

func TestGetFilterValues_Pointers(t *testing.T) {
	input := tablea{}
	actual := getFilterValues(&input)

	if val, ok := actual["Relation"]; ok {
		t.Errorf("Pointer has value %s", val)
	}
}

func TestGetRelationships(t *testing.T) {
	input := tablea{}
	input.Collections = []*tablea{
		&tablea{
			Name: "TEST",
		},
	}
	input.Relation = &tablea{
		Name: "RELATE",
	}

	actual := getRelationships(&input)

	for _, v := range actual {
		fields := getReadColumns(v)

		if len(fields) <= 0 {
			t.Error("no fields found")
		}
	}
}

func TestNestedRelationshipsModel(t *testing.T) {
	input := tablea{}
	input.Relation = &tablea{
		Name: "RELATE",
	}

	relationships := getRelationships(&input)

	for _, v := range relationships {
		readCols := getReadColumns(v)

		if len(readCols) <= 0 {
			t.Error("Exptected items in readCols.")
		}
	}
}
