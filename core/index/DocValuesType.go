package index

type DocValuesType int

const (
	// DocValuesTypeNone No doc values for this field.
	DocValuesTypeNone = DocValuesType(iota)

	// DocValuesTypeNumeric A per-document Number
	DocValuesTypeNumeric

	// DocValuesTypeBinary A per-document byte[]. Values may be larger than 32766 bytes,
	//but different codecs may enforce their own limits.
	DocValuesTypeBinary

	// DocValuesTypeSorted A pre-sorted byte[]. Fields with this type only store distinct byte values and store an
	// additional offset pointer per document to dereference the shared byte[]. The stored byte[] is
	// presorted and allows access via document id, ordinal and by-value. Values must be <=
	// 32766} bytes.
	DocValuesTypeSorted

	// DocValuesTypeSortedNumeric A pre-sorted Number[]. Fields with this type store numeric values in sorted order according to
	// Long#compare(long, long)}.
	DocValuesTypeSortedNumeric

	// DocValuesTypeSortedSet A pre-sorted Set&lt;byte[]&gt;. Fields with this type only store distinct byte values and store
	// additional offset pointers per document to dereference the shared byte[]s. The stored byte[] is
	// presorted and allows access via document id, ordinal and by-value. Values must be <=
	// 32766} bytes.
	DocValuesTypeSortedSet
)

func (typ DocValuesType) String() string {
	switch typ {
	case DocValuesTypeNone:
		return "NONE"
	case DocValuesTypeNumeric:
		return "NUMERIC"
	case DocValuesTypeBinary:
		return "BINARY"
	case DocValuesTypeSorted:
		return "SORTED"
	case DocValuesTypeSortedNumeric:
		return "SORTED_NUMERIC"
	case DocValuesTypeSortedSet:
		return "SORTED_SET"
	default:
		return "NONE"
	}
}
