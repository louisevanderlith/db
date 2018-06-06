package tests

import (
	"time"

	"github.com/louisevanderlith/db"
)

type utContext struct {
	People       db.Setter
	Transactions db.Setter
}

var utCTX *utContext

func init() {
	utCTX = &utContext{
		People:       db.NewDBSet(personTable{}),
		Transactions: db.NewDBSet(transactionTable{}),
	}

	db.SyncDatabase(utCTX, "memory")
}

type People []*personTable

func (r People) Each(handler func(db.IRecord)) {
	for _, v := range r {
		handler(v)
	}
}

func (r People) Length() int {
	return len(r)
}

func (r People) At(index int) db.IRecord {
	return r[index]
}

func (r People) Add(obj db.IRecord) {
	item, ok := obj.(*personTable)

	if ok {
		r = append(r, item)
	}
}

type personTable struct {
	db.Record
	Name string
	Age  int
}

func (t personTable) Validate() (bool, error) {
	return true, nil
}

func newPersonTable() personTable {
	return personTable{
		Record: db.Record{
			Id:         0,
			Deleted:    false,
			CreateDate: time.Now(),
		},
	}
}

type transactionTable struct {
	db.Record
	Amount float32
	Person *personTable
	Paid   []*personTable
}

func (t transactionTable) Validate() (bool, error) {
	return true, nil
}

func newTransactionTable() transactionTable {
	return transactionTable{
		Record: db.Record{
			Id:         0,
			Deleted:    false,
			CreateDate: time.Now(),
		},
	}
}
