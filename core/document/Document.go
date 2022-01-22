package document

// Document Documents are the unit of indexing and search.
// A Document is a set of fields. Each field has a name and a textual value. A field may be
// org.apache.lucene.index.IndexableFieldType#stored() stored with the document, in which
// case it is returned with search hits on the document. Thus each document should typically contain
// one or more stored fields which uniquely identify it.
// Note that fields which are not
// org.apache.lucene.index.IndexableFieldType#stored() stored are not available in documents
// retrieved from the index, e.g. with ScoreDoc#doc or IndexReader#document(int).
type Document struct {
	fields []*Field
}

// Add Adds a field to a document. Several fields may be added with the same name. In this case, if
// the fields are indexed, their text is treated as though appended for the purposes of search.
// Note that add like the removeField(s) methods only makes sense prior to adding a document to
// an index. These methods cannot be used to change the content of an existing index! In order to
// achieve this, a document has to be deleted from an index and a new changed version of that
// document has to be added.
func (d *Document) Add(field *Field) {
	d.fields = append(d.fields, field)
}

// RemoveField Removes field with the specified name from the document. If multiple fields exist with this
// name, this method removes the first field that has been added. If there is no field with the
// specified name, the document remains unchanged.
// Note that the removeField(s) methods like the add method only make sense prior to adding a
// document to an index. These methods cannot be used to change the content of an existing index!
// In order to achieve this, a document has to be deleted from an index and a new changed version
// of that document has to be added.
func (d *Document) RemoveField(name string) {
	for i := 0; i < len(d.fields); i++ {
		if d.fields[i].Name() == name {
			d.fields = append(d.fields[:i], d.fields[i+1:]...)
			return
		}
	}
}

// RemoveFields Removes all fields with the given name from the document. If there is no field with the
// specified name, the document remains unchanged.
// Note that the removeField(s) methods like the add method only make sense prior to adding a
// document to an index. These methods cannot be used to change the content of an existing index!
// In order to achieve this, a document has to be deleted from an index and a new changed version
// of that document has to be added.
func (d *Document) RemoveFields(name string) {

}

// GetBinaryValues Returns an array of byte arrays for of the fields that have the name specified as the method
// parameter. This method returns an empty array when there are no matching fields. It never
// returns null.
func (d *Document) GetBinaryValues(name string) [][]byte {
	panic("")
}

// GetBinaryValue Returns an array of bytes for the first (or only) field that has the name specified as the
// method parameter. This method will return null if no binary fields with the
// specified name are available. There may be non-binary fields with the same name.
func (d *Document) GetBinaryValue(name string) []byte {
	panic("")
}

// GetField Returns a field with the given name if any exist in this document, or null. If multiple fields
// exists with this name, this method returns the first value added.
func (d *Document) GetField(name string) *Field {
	panic("")
}

// GetFieldsByName Returns an array of IndexableFields with the given name. This method returns an empty
// array when there are no matching fields. It never returns null.
func (d *Document) GetFieldsByName(name string) []*Field {
	panic("")
}

// GetFields Returns a List of all the fields in a document.
// Note that fields which are not stored are not available in documents retrieved
// from the index, e.g. IndexSearcher#doc(int) or IndexReader#document(int).
func (d *Document) GetFields() []*Field {
	return d.fields
}

// GetValues Returns an array of values of the field specified as the method parameter. This method returns
// an empty array when there are no matching fields. It never returns null. For a numeric
// StoredField} it returns the string value of the number. If you want the actual numeric field
// instances back, use #getFields}.
func (d *Document) GetValues() []string {
	panic("")
}

func (d *Document) Clear() {
	d.fields = d.fields[:0]
}

type Node struct {
	Value *Field
	Next  *Node
}
