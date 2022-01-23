package index

import . "github.com/geange/smarch/core/util"

// FieldInvertState This class tracks the number and position / offset parameters of terms being added to the index.
// The information collected in this class is also used to calculate the normalization factor for a
// field.
type FieldInvertState struct {
	indexCreatedVersionMajor int
	name                     string
	indexOptions             IndexOptions
	position                 int
	length                   int
	numOverlap               int
	offset                   int
	maxTermFrequency         int
	uniqueTermCount          int

	// we must track these across field instances (multi-valued case)
	lastStartOffset int
	lastPosition    int

	attributeSource *AttributeSource

	offsetAttribute OffsetAttribute
}
