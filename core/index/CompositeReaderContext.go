package index

// CompositeReaderContext : IndexReaderContext for CompositeReader instance.
type CompositeReaderContext struct {

	//The reader context for this reader's immediate parent, or null if none
	parent *CompositeReaderContext

	// true if this context struct represents the top level reader within the hierarchical context
	isTopLevel bool

	//the doc base for this reader in the parent, 0 if parent is null
	docBaseInParent int

	//the ord for this reader in the parent, 0 if parent is null
	ordInParent int

	// An object that uniquely identifies this context without referencing
	// segments. The goal is to make it fine to have references to this
	// identity object, even after the index reader has been closed
	identity string

	children []IndexReaderContext
	leaves   []LeafReaderContext
	reader   CompositeReader
}

func (c *CompositeReaderContext) ID() string {
	return c.identity
}

func (c *CompositeReaderContext) Reader() IndexReader {
	return c.reader
}

func (c *CompositeReaderContext) Leaves() ([]LeafReaderContext, error) {
	return c.leaves, nil
}

func (c *CompositeReaderContext) Children() []IndexReaderContext {
	return c.children
}
