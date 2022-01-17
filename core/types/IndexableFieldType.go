package types

type IndexableFieldType struct {
}

type (
	FieldDataType int
)

func GetFieldDataType(data interface{}) FieldDataType {
	panic("impl it")
}

const (
	FieldDataTypeBytes = FieldDataType(iota)
	FieldDataTypeString
	FieldDataTypeInt32
	FieldDataTypeInt64
	FieldDataTypeFloat32
	FieldDataTypeFloat64
)
