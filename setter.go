package db

import "reflect"

// Setter is an interface that provides functionality to add, update, delete & find records
type Setter interface {
	// GetRecordType returns the true type of the IRecord
	GetRecordType() reflect.Type

	// Create validates the model and then saves that record to the database.
	// # id, err := folio.Ctx.Profile.Create(&profile)
	Create(item IRecord) (id int64, err error)

	// CreateMulti inserts multiple records, without running validation
	// # manufacturers := []Manufacturer{}
	// # count, err := folio.Ctx.Manufacturer.CreateMulti(manufacturers)
	CreateMulti(items IRecordCollection) (insertCount int64, err error)

	// ReadOne reads a single record from the database
	// filter: An object that has the fields populated that you want to filter on (Filters will always be 'AND')
	// related: Relationships are lazy-loaded, to include nested items you must specify them.
	// # record, err := testCtx.Profile.ReadOne(&Profile{ID: 56}, "User")
	ReadOne(filter IRecord, related ...string) (IRecord, error)

	// Read reads all records that fit the filter provided
	// filter: An object that has the fields populated that you want to filter on (Filters will always be 'AND')
	// container: The result set will populate the container.
	// # var results []*artifact.Upload
	// # upl := artifact.Upload{Type: "JPEG"}
	// # err := artifact.Ctx.Upload.Read(&upl, &results)
	Read(filter IRecord, container IRecordCollection) error

	// Fetch reads all records that fit the filter provided
	Fetch(page, pageSize int, expr func(obj IRecord) bool) (IRecordCollection, error)

	// Update saves the provided record to the database.
	// The record must exist in the database.
	// item: The record you want to update.
	// # row := tableA{}
	// #  testCtx.TableA.Update(&row)
	Update(item IRecord) error

	// Delete will delete the record from the database.
	// This function currently only deletes a record based on the provided ID
	// item: The record containing the ID you want to delete.
	// # row := TableA{ID: 99}
	// # testCtx.TableA.Delete(row)
	Delete(item IRecord) error
}

// New creates an instance of the selected Setter implementation
var New func(t IRecord) Setter
