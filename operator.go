package db

// Operator provides an interface to the datastore
type Operator interface {
	//Insert adds a new record and returns the ID created
	Insert(obj interface{}) (int64, error)

	// InsertMulti adds multiple records, returns the number of rows saved
	InsertMulti(batchCount int, objs interface{}) (int64, error)

	// ReadOne reads one record filtered by the values already specified in obj
	ReadOne(obj interface{}, related ...string) error

	// Read reads all records filtered by the values specified in the filter
	Read(filter interface{}, container interface{}) error

	// Update will persist any changes made to a record
	Update(obj interface{}) (int64, error)
}

// NewOperator returns an implementation of an Operator
var NewOperator func() Operator
