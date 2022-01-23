package index

import "sync"

// FieldInfos Collection of FieldInfos (accessible by number or by name).
type FieldInfos struct {
	fieldInfos []FieldInfo

	hasFreq         bool
	hasProx         bool
	hasPayloads     bool
	hasOffsets      bool
	hasVectors      bool
	hasNorms        bool
	hasDocValues    bool
	hasPointValues  bool
	hasVectorValues bool

	softDeletesField string

	// used only by fieldInfo(int)
	byNumber []FieldInfo
	byName   map[string]FieldInfo

	// for an unmodifiable iterator
	values []FieldInfo
}

type FieldNumbers struct {
	numberToName map[int]string
	nameToNumber map[string]int
	indexOptions map[string]IndexOptions

	// We use this to enforce that a given field never
	// changes DV type, even across segments / IndexWriter
	// sessions:
	docValuesType map[string]DocValuesType

	dimensions map[string]FieldDimensions

	vectorProps      map[string]FieldVectorProperties
	omitNorms        map[string]bool
	storeTermVectors map[string]bool

	// TODO: we should similarly catch an attempt to turn
	// norms back on after they were already committed; today
	// we silently discard the norm but this is badly trappy
	// deault -1
	lowestUnassignedFieldNumber int

	// The soft-deletes field from IWC to enforce a single soft-deletes field
	softDeletesFieldName string

	sync.RWMutex
}

func NewFieldNumbers(softDeletesFieldName string) *FieldNumbers {
	return &FieldNumbers{
		numberToName:                map[int]string{},
		nameToNumber:                map[string]int{},
		indexOptions:                map[string]IndexOptions{},
		docValuesType:               map[string]DocValuesType{},
		dimensions:                  map[string]FieldDimensions{},
		vectorProps:                 map[string]FieldVectorProperties{},
		omitNorms:                   map[string]bool{},
		storeTermVectors:            map[string]bool{},
		lowestUnassignedFieldNumber: -1,
		softDeletesFieldName:        softDeletesFieldName,
		RWMutex:                     sync.RWMutex{},
	}
}

// AddOrGet Returns the global field number for the given field name. If the name does not exist yet it
// tries to add it with the given preferred field number assigned if possible otherwise the
// first unassigned field number is used as the field number.
func (f *FieldNumbers) AddOrGet(fi *FieldInfo) int {
	panic("")
}

// VerifyOrCreateDvOnlyField This function is called from IndexWriter to verify if doc values of the field can be
// updated. If the field with this name already exists, we verify that it is doc values only
// field. If the field doesn't exists and the parameter fieldMustExist is false, we create a new
// field in the global field numbers.
func (f *FieldNumbers) VerifyOrCreateDvOnlyField(fieldName string, dvType DocValuesType, fieldMustExist bool) error {
	panic("")
}

// ConstructFieldInfo Construct a new FieldInfo based on the options in global field numbers. This method is not
// synchronized as all the options it uses are not modifiable.
func (f *FieldNumbers) ConstructFieldInfo(fieldName string, dvType DocValuesType, newFieldNumber int) *FieldInfo {
	panic("")
}

func (f *FieldNumbers) getFieldNames() map[string]struct{} {
	result := make(map[string]struct{}, len(f.nameToNumber))
	for key, _ := range result {
		result[key] = struct{}{}
	}
	return result
}

func (f *FieldNumbers) Clear() {
	f.numberToName = map[int]string{}
	f.nameToNumber = map[string]int{}
	f.indexOptions = map[string]IndexOptions{}
	f.docValuesType = map[string]DocValuesType{}
	f.dimensions = map[string]FieldDimensions{}
	f.lowestUnassignedFieldNumber = -1
}

type FieldDimensions struct {
	DimensionCount      int
	IndexDimensionCount int
	DimensionNumBytes   int
}

type FieldVectorProperties struct {
	numDimensions      int
	similarityFunction VectorSimilarityFunction
}
