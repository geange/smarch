package index

// IndexingChain Default general purpose indexing chain,
// which handles indexing all types of fields.
type IndexingChain struct {
	bytesUsed int64
}

// PerField NOTE: not static: accesses at least docState, termsHash.
type PerField struct {
	fieldName                string
	indexCreatedVersionMajor int
	schema                   *FieldSchema
	fieldInfo                *FieldInfo
	similarity               Similarity
}

// FieldSchema A schema of the field in the current document. With every new document this schema is reset. As
// the document fields are processed, we update the schema with options encountered in this
// document. Once the processing for the document is done, we compare the built schema of the
// current document with the corresponding FieldInfo (FieldInfo is built on a first document in
// the segment where we encounter this field). If there is inconsistency, we raise an error. This
// ensures that a field has the same data structures across all documents.
type FieldSchema struct {
}
