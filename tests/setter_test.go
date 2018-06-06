package tests

import (
	"testing"

	"github.com/louisevanderlith/db"
)

func TestNewContext(t *testing.T) {
	if utCTX.People == nil {
		t.Error("Context didn't initialize.")
	}
}

func TestSet_Create(t *testing.T) {
	row := newPersonTable()
	row.Name = "ABC"
	row.Id = 99

	_, err := utCTX.People.Create(&row)

	if err != nil {
		t.Error(err)
	}

	var filter = &personTable{}
	filter.Id = 99

	person, perr := utCTX.People.ReadOne(filter)

	if perr != nil {
		t.Error(perr)
	}

	if person == nil || person.GetID() != row.GetID() {
		t.Errorf("Expected %v got %v", row.GetID(), person.GetID())
	}
}

func TestSet_Read(t *testing.T) {
	row := newPersonTable()
	row.Name = "ABC"
	row.Id = 99
	utCTX.People.Create(row)

	rowb := newPersonTable()
	rowb.Name = "DEF"
	rowb.Id = 98
	rowb.Deleted = true
	utCTX.People.Create(rowb)

	var records People
	err := utCTX.People.Read(&personTable{}, &records)

	if err != nil {
		t.Error(err)
	}

	for _, v := range records {
		if v.IsDeleted() {
			t.Error("records shouldn't contain deleted records.")
			break
		}
	}
}

func TestSet_ReadOne_Nil(t *testing.T) {
	row := newPersonTable()
	row.Name = "ABC"
	row.Id = 55
	row.Deleted = true
	utCTX.People.Create(row)

	record, err := utCTX.People.ReadOne(&personTable{Record: db.Record{Id: 55}})

	if err != nil {
		t.Error(err)
	}

	if record != nil {
		t.Error("Deleted records shouldn't be returned.")
	}
}

func TestSet_ReadOne(t *testing.T) {
	row := newPersonTable()
	row.Name = "ABC"
	row.Id = 56
	utCTX.People.Create(row)

	record, err := utCTX.People.ReadOne(&personTable{Record: db.Record{Id: 56}})

	if err != nil {
		t.Error(err)
	}

	if record == nil {
		t.Error("Added record not found.")
	}
}

func TestSet_Update(t *testing.T) {
	row := newPersonTable()
	row.Name = "ABC"
	row.Id = 56
	utCTX.People.Create(&row)

	row.Age = 18

	utCTX.People.Update(&row)

	record, err := utCTX.People.ReadOne(&personTable{Record: db.Record{Id: 56}})

	if err != nil {
		t.Error(err)
	}

	tblRec := record.(*personTable)

	if tblRec.Age != row.Age {
		t.Errorf("Record didn't update. Expected %v got %v", row.Age, tblRec.Age)
	}
}

func TestSet_Delete(t *testing.T) {
	row := newPersonTable()
	row.Name = "ABC"
	row.Id = 99
	row.Deleted = false

	utCTX.People.Create(row)

	record, err := utCTX.People.ReadOne(&personTable{Record: db.Record{Id: 99}})

	if err != nil {
		t.Error(err)
	}

	if record == nil {
		t.Error("Record is nil")
	}

	record.Disable()

	if !record.IsDeleted() {
		t.Error("Record hasn't been set as Deleted")
	}
}
