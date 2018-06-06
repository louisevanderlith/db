package db

// IRecordCollection provides a way to work with IRecord collections
type IRecordCollection interface {
	Each(handler func(IRecord))
	Length() int
	At(index int) IRecord
	Add(obj IRecord)
}
