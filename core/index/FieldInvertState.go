package index

import . "github.com/geange/smarch/core/util"

// FieldInvertState This class tracks the number and position / offset parameters of terms being added to the index.
// The information collected in this class is also used to calculate the normalization factor for a field.
// 追踪添加到索引的 term 的数量、位置、偏移量。收集的信息用于计算 field 的归一化因子
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
