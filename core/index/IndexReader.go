package index

import (
	. "github.com/geange/smarch/core/document"
	"io"
)

type IndexReader interface {
	io.Closer

	// GetTermVectors Retrieve term vectors for this document, or null if term vectors were not indexed.
	// The returned Fields instance acts like a single-document inverted index (the docID will be 0).
	GetTermVectors(docID int) (*Fields, error)

	// GetTermVector Retrieve term vector for this document and field, or null if term vectors were not indexed.
	// The returned Fields instance acts like a single-document inverted index (the docID will be 0).
	GetTermVector(docID int, field string) (Terms, error)

	// NumDocs Returns the number of documents in this index.
	//
	// NOTE: This operation may run in O(maxDoc). Implementations that can't return this
	// number in constant-time should cache it.
	NumDocs() int

	// MaxDoc Returns one greater than the largest possible document number. This may be used to, e.g.,
	// determine how big to allocate an array which will have an element for every document number in
	// an index.
	MaxDoc() int

	// NumDeletedDocs Returns the number of deleted documents.
	// NOTE: This operation may run in O(maxDoc).
	NumDeletedDocs() int

	// Document Expert: visits the fields of a stored document, for custom processing/loading of each field.
	// If you simply want to load all fields, use #document(int).
	// If you want to load a subset, use DocumentStoredFieldVisitor.
	Document(docID int, visitor StoredFieldVisitor) error

	// DocumentV1 Returns the stored fields of the nth Document in this
	// index. This is just sugar for using DocumentStoredFieldVisitor.
	//
	// NOTE: for performance reasons, this method does not check if the requested document
	// is deleted, and therefore asking for a deleted document may yield unspecified results. Usually
	// this is not required, however you can test if the doc is deleted by checking the Bits
	// returned from MultiBits#getLiveDocs.
	//
	// NOTE: only the content of a field is returned, if that field was stored during
	// indexing. Metadata like boost, omitNorm, IndexOptions, tokenized, etc., are not preserved.
	DocumentV1(docID int64) (Document, error)

	// DocumentV2 Like #document(int) but only loads the specified fields. Note that this is simply sugar
	// for DocumentStoredFieldVisitor#DocumentStoredFieldVisitor(Set).
	DocumentV2(docID int64, fieldsToLoad map[string]struct{}) (Document, error)

	// HasDeletions Returns true if any documents have been deleted. Implementers should consider overriding this
	// method if #maxDoc() or #numDocs() are not constant-time operations.
	HasDeletions() bool

	// DoClose Implements close.
	DoClose() error

	// GetContext Expert: Returns the root IndexReaderContext for this IndexReader's sub-reader
	// tree.
	//
	// Iff this reader is composed of sub readers, i.e. this reader being a composite reader, this
	// method returns a CompositeReaderContext holding the reader's direct children as well as
	// a view of the reader tree's atomic leaf contexts. All sub- IndexReaderContext instances
	// referenced from this readers top-level context are private to this reader and are not shared
	// with another context tree. For example, IndexSearcher uses this API to drive searching by one
	// atomic leaf reader at a time. If this reader is not composed of child readers, this method
	// returns an LeafReaderContext.
	//
	// Note: Any of the sub-CompositeReaderContext instances referenced from this top-level
	// context do not support CompositeReaderContext#leaves(). Only the top-level context
	// maintains the convenience leaf-view for performance reasons.
	GetContext() IndexReaderContext

	// Leaves Returns the reader's leaves, or itself if this reader is atomic.
	// This is a convenience method calling this.getContext().leaves().
	Leaves() []LeafReaderContext

	// Optional method: Return a CacheHelper that can be used to cache based on the content of
	// this reader. Two readers that have different data or different sets of deleted documents will
	// be considered different.
	//
	// A return value of null indicates that this reader is not suited for caching, which
	// is typically the case for short-lived wrappers that alter the content of the wrapped reader.
	getReaderCacheHelper() CacheHelper

	// DocFreq Returns the number of documents containing the term. This method returns 0 if the
	// term or field does not exists. This method does not take into account deleted documents that
	// have not yet been merged away.
	DocFreq(term *Term) (int, error)

	// TotalTermFreq Returns the total number of occurrences of term across all documents (the sum of the
	// freq() for each doc that has this term). Note that, like other term measures, this measure does
	// not take deleted documents into account.
	TotalTermFreq(term *Term) (int64, error)

	// GetSumDocFreq Returns the sum of TermsEnum#docFreq() for all terms in this field. Note that, just
	// like other term measures, this measure does not take deleted documents into account.
	GetSumDocFreq(field string) (int64, error)

	// GetDocCount Returns the number of documents that have at least one term for this field. Note that, just
	// like other term measures, this measure does not take deleted documents into account.
	GetDocCount(field string) (int, error)

	// GetSumTotalTermFreq Returns the sum of TermsEnum#totalTermFreq for all terms in this field. Note that, just
	// like other term measures, this measure does not take deleted documents into account.
	GetSumTotalTermFreq(field string) (int64, error)
}

type CacheHelper interface {
	// GetKey Get a key that the resource can be cached on. The given entry can be compared using identity,
	// ie. Object#equals is implemented as == and Object#hashCode is implemented as System#identityHashCode.
	GetKey() CacheKey

	// Add a ClosedListener which will be called when the resource guarded by #getKey() is closed.
	addClosedListener(listener ClosedListener)
}

type CacheKey struct {
}

// ClosedListener A listener that is called when a resource gets closed.
type ClosedListener interface {
	// OnClose Invoked when the resource (segment core, or index reader) that is being cached on is closed.
	OnClose(key *CacheKey) error
}
