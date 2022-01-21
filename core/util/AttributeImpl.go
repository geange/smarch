package util

// AttributeImpl Base class for Attributes that can be added to a AttributeSource.
//
// Attributes are used to add data in a dynamic, yet type-safe way to a source of usually
// streamed objects, e. g. a org.apache.lucene.analysis.TokenStream.
//
// All implementations must list all implemented Attribute interfaces in their
// implements clause. AttributeSource reflectively identifies all attributes and makes them
// available to consumers like TokenStreams.
type AttributeImpl interface {

	// Clear Clears the values in this AttributeImpl and resets it to its default value. If this
	// implementation implements more than one Attribute interface it clears all.
	Clear()

	// End Clears the values in this AttributeImpl and resets it to its value at the end of the field. If
	// this implementation implements more than one Attribute interface it clears all.
	End()

	// CopyTo Copies the values from this Attribute into the passed-in target attribute. The target
	// implementation must support all the Attributes this implementation supports.
	CopyTo(target AttributeImpl)

	// Clone In most cases the clone is, and should be, deep in order to be able to properly capture the
	// state of all attributes.
	Clone() AttributeImpl
}
