package index

import (
	"errors"
	"fmt"
)

// FieldInfo Access to the Field Info file that describes document fields and whether or not they are indexed.
// Each segment has a separate Field Info file. Objects of this class are thread-safe for multiple
// readers, but only one thread can be adding documents at a time, with no other reader or writer
// threads accessing this object.
type FieldInfo struct {
	//Field's name
	name string

	//Internal field number
	number int

	// default DocValuesType.NONE
	docValuesType DocValuesType

	// True if any document indexed term vectors
	storeTermVector bool

	// omit norms associated with indexed fields
	omitNorms bool

	indexOptions IndexOptions

	// whether this field stores payloads together with term positions
	storePayloads bool

	attributes map[string]string

	dvGen int64

	//If both of these are positive it means this field indexed points (see
	// org.apache.lucene.codecs.PointsFormat)
	pointDimensionCount int

	pointIndexDimensionCount int
	pointNumBytes            int

	// if it is a positive value, it means this field indexes vectors
	vectorDimension          int
	vectorSimilarityFunction VectorSimilarityFunction

	// whether this field is used as the soft-deletes field
	softDeletesField bool
}

// CheckConsistency Check correctness of the FieldInfo options
func (f *FieldInfo) CheckConsistency() error {
	panic("")
}

// VerifySameSchema Verify that the provided FieldInfo has the same schema as this FieldInfo
func (f *FieldInfo) VerifySameSchema(o *FieldInfo) error {
	panic("")
}

// VerifySameIndexOptions Verify that the provided index options are the same
func (f *FieldInfo) VerifySameIndexOptions(fileName string, indexOptions1, indexOptions2 IndexOptions) error {
	panic("")
}

// VerifySameStoreTermVectors Verify that the provided store term vectors options are the same
func (f *FieldInfo) VerifySameStoreTermVectors(fieldName string, storeTermVector1, storeTermVector2 bool) error {
	panic("")
}

// VerifySameOmitNorms Verify that the provided omitNorms are the same
func (f *FieldInfo) VerifySameOmitNorms(fieldName string, omitNorms1, omitNorms2 string) error {
	panic("")
}

// VerifySamePointsOptions Verify that the provided points indexing options are the same
func (f *FieldInfo) VerifySamePointsOptions(fieldName string, pointDimensionCount1, indexDimensionCount1, numBytes1,
	pointDimensionCount2, indexDimensionCount2, numBytes2 int) {
	panic("")
}

// VerifySameVectorOptions Verify that the provided vector indexing options are the same
func (f *FieldInfo) VerifySameVectorOptions(fieldName string, vd1 int, vsf1 VectorSimilarityFunction,
	vd2 int, vsf2 VectorSimilarityFunction) {
	panic("")
}

// SetPointDimensions Record that this field is indexed with points, with the specified number
//of dimensions and bytes per dimension.
func (f *FieldInfo) SetPointDimensions(dimensionCount, indexDimensionCount, numBytes int) {
	panic("")
}

// GetPointDimensionCount Return point data dimension count
func (f *FieldInfo) GetPointDimensionCount() int {
	return f.pointDimensionCount
}

// GetPointIndexDimensionCount Return point data dimension count
func (f *FieldInfo) GetPointIndexDimensionCount() int {
	return f.pointIndexDimensionCount
}

// GetPointNumBytes Return number of bytes per dimension
func (f *FieldInfo) GetPointNumBytes() int {
	return f.pointNumBytes
}

// GetVectorDimension Returns the number of dimensions of the vector value
func (f *FieldInfo) GetVectorDimension() int {
	return f.vectorDimension
}

// GetVectorSimilarityFunction Returns VectorSimilarityFunction for the field
func (f *FieldInfo) GetVectorSimilarityFunction() VectorSimilarityFunction {
	return f.vectorSimilarityFunction
}

// SetDocValuesType Record that this field is indexed with docvalues, with the specified type
func (f *FieldInfo) SetDocValuesType(typ DocValuesType) error {
	if f.docValuesType != DocValuesTypeNone && typ != DocValuesTypeNone && f.docValuesType != typ {
		return fmt.Errorf("cannot change DocValues type from %s to %s",
			f.docValuesType.String(), typ.String())
	}
	f.docValuesType = typ
	return f.CheckConsistency()
}

// GetIndexOptions Returns IndexOptions for the field, or IndexOptions.NONE if the field is not indexed
func (f *FieldInfo) GetIndexOptions() IndexOptions {
	return f.indexOptions
}

// GetName Returns name of this field
func (f *FieldInfo) GetName() string {
	return f.name
}

// GetFieldNumber Returns the field number
func (f *FieldInfo) GetFieldNumber() int {
	return f.number
}

// GetDocValuesType Returns DocValuesType of the docValues; this is DocValuesTypeNone if the field has no docvalues.
func (f *FieldInfo) GetDocValuesType() DocValuesType {
	return f.docValuesType
}

// SetDocValuesGen Sets the docValues generation of this field.
func (f *FieldInfo) SetDocValuesGen(dvGen int64) error {
	f.dvGen = dvGen
	return f.CheckConsistency()
}

// GetDocValuesGen Returns the docValues generation of this field, or -1 if no docValues updates exist for it.
func (f *FieldInfo) GetDocValuesGen() int64 {
	return f.dvGen
}

func (f *FieldInfo) SetStoreTermVectors() error {
	f.storeTermVector = true
	return f.CheckConsistency()
}

func (f *FieldInfo) SetStorePayloads() error {
	if f.indexOptions > IdxOptDocsAndFreqsAndPositions {
		f.storePayloads = true
	}
	return f.CheckConsistency()
}

// OmitsNorms Returns true if norms are explicitly omitted for this field
func (f *FieldInfo) OmitsNorms() bool {
	return f.omitNorms
}

// SetOmitsNorms Omit norms for this field.
func (f *FieldInfo) SetOmitsNorms() error {
	if f.indexOptions == IdxOptNone {
		return errors.New("cannot omit norms: this field is not indexed")
	}
	f.omitNorms = true
	return f.CheckConsistency()
}

// HasNorms Returns true if this field actually has any norms.
func (f *FieldInfo) HasNorms() bool {
	return f.indexOptions != IdxOptNone && f.omitNorms == false
}

// HasPayloads Returns true if any payloads exist for this field.
func (f *FieldInfo) HasPayloads() bool {
	return f.storePayloads
}

// HasVectors Returns true if any term vectors exist for this field.
func (f *FieldInfo) HasVectors() bool {
	return f.storeTermVector
}

// HasVectorValues Returns whether any (numeric) vector values exist for this field
func (f *FieldInfo) HasVectorValues() bool {
	return f.vectorDimension > 0
}

// GetAttribute Get a codec attribute value, or null if it does not exist
func (f *FieldInfo) GetAttribute(key string) string {
	return f.attributes[key]
}

// PutAttribute Puts a codec attribute value.
//
// This is a key-value mapping for the field that the codec can use to store additional
// metadata, and will be available to the codec when reading the segment via
// #getAttribute(String)
//
// If a value already exists for the key in the field, it will be replaced with the new value.
// If the value of the attributes for a same field is changed between the documents, the behaviour
// after merge is undefined.
func (f *FieldInfo) PutAttribute(key, value string) {
	f.attributes[key] = value
}

// Attributes Returns internal codec attributes map.
func (f *FieldInfo) Attributes() map[string]string {
	return f.attributes
}

// IsSoftDeletesField Returns true if this field is configured and used as the soft-deletes field.
// See IndexWriterConfig#softDeletesField
func (f *FieldInfo) IsSoftDeletesField() bool {
	return f.softDeletesField
}
