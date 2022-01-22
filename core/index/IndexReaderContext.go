package index

// IndexReaderContext A struct like class that represents a hierarchical relationship between IndexReader instances.
type IndexReaderContext interface {

	// ID Expert: Return an Object that uniquely identifies this context. The returned object
	// does neither reference this IndexReaderContext nor the wrapped IndexReader.
	ID() string

	// Reader Returns the IndexReader, this context represents.
	Reader() IndexReader

	// Leaves Returns the context's leaves if this context is a top-level context. For convenience,
	// if this is an LeafReaderContext this returns itself as the only leaf.
	//
	// Note: this is convenience method since leaves can always be obtained by walking the context
	// tree using #children().
	Leaves() ([]LeafReaderContext, error)

	// Children Returns the context's children if this context is a composite context otherwise null
	Children() []IndexReaderContext
}
