package index

// StoredFieldVisitor Expert: provides a low-level means of accessing the stored field values in an index. See
// IndexReader#document(int, StoredFieldVisitor).
//
// NOTE: a StoredFieldVisitor implementation should not try to load or visit other
// stored documents in the same reader because the implementation of stored fields for most codecs
// is not reentrant and you will see strange exceptions as a result.
//
// See DocumentStoredFieldVisitor, which is a StoredFieldVisitor that builds
// the Document containing all stored fields. This is used by IndexReader#document(int)}.
type StoredFieldVisitor interface {

	// BinaryField Process a binary field.
	BinaryField(fieldInfo *FieldInfo, value []byte) error

	// StringField Process a string field.
	StringField(fieldInfo *FieldInfo, value string) error

	// Int32Field Process a int numeric field.
	Int32Field(fieldInfo *FieldInfo, value int32) error

	// Int64Field Process a long numeric field.
	Int64Field(fieldInfo *FieldInfo, value int64) error

	// Float32Field Process a float numeric field.
	Float32Field(fieldInfo *FieldInfo, value float32) error

	// Float64Field Process a double numeric field.
	Float64Field(fieldInfo *FieldInfo, value float64) error

	// Hook before processing a field. Before a field is processed, this method is invoked so that
	// subclasses can return a Status representing whether they need that particular field or
	// not, or to stop processing entirely.
}

// NeedsFieldStatus Enumeration of possible return values for #needsField.
type NeedsFieldStatus int

const (
	// NeedsFieldStatusYes : the field should be visited.
	NeedsFieldStatusYes = NeedsFieldStatus(iota)

	// NeedsFieldStatusNo : don't visit this field, but continue processing fields for this document.
	NeedsFieldStatusNo

	// NeedsFieldStatusStop : don't visit this field and stop processing any other fields for this document.
	NeedsFieldStatusStop
)
