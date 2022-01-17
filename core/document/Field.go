package document

import . "github.com/geange/smarch/core/types"

type Field struct {
	//Field's type
	_type *IndexableFieldType

	//Field's name
	name string

	//Field's value
	fDataType  FieldDataType
	fieldsData interface{}

	//Pre-analyzed tokenStream for indexed fields; this is separate from fieldsData because you are
	//allowed to have both; eg maybe field has a String value but you customize how it's tokenized
	tokenStream TokenStream
}

func NewField(name string, data interface{}, _type *IndexableFieldType) *Field {
	return &Field{
		_type:       _type,
		name:        name,
		fDataType:   GetFieldDataType(data),
		fieldsData:  data,
		tokenStream: nil,
	}
}

func (f *Field) Name() string {
	return f.name
}

func (f *Field) Value() interface{} {
	return f.fieldsData
}

func (f *Field) ValueBytes() []byte {
	switch f.fDataType {
	case FieldDataTypeBytes:
		return f.fieldsData.([]byte)
	default:
		return nil
	}
}

func (f *Field) ValueString() string {
	switch f.fDataType {
	case FieldDataTypeString:
		return f.fieldsData.(string)
	default:
		return ""
	}
}

func (f *Field) ValueInt32() int32 {
	switch f.fDataType {
	case FieldDataTypeInt32:
		return f.fieldsData.(int32)
	default:
		return 0
	}
}

func (f *Field) ValueInt64() int64 {
	switch f.fDataType {
	case FieldDataTypeInt64:
		return f.fieldsData.(int64)
	default:
		return 0
	}
}

func (f *Field) ValueFloat32() float32 {
	switch f.fDataType {
	case FieldDataTypeFloat32:
		return f.fieldsData.(float32)
	default:
		return 0
	}
}

func (f *Field) ValueFloat64() int64 {
	switch f.fDataType {
	case FieldDataTypeFloat64:
		return f.fieldsData.(int64)
	default:
		return 0
	}
}
