package index

// LeafReader is an abstract class, providing an interface for accessing an index. Search of
// an index is done entirely through this abstract interface, so that any subclass which implements
// it is searchable. IndexReaders implemented by this subclass do not consist of several
// sub-readers, they are atomic. They support retrieval of stored fields, doc values, terms, and
// postings.
//
// For efficiency, in this API documents are often referred to via document numbers,
// non-negative integers which each name a unique document in the index. These document numbers are
// ephemeral -- they may change as documents are added to and deleted from an index. Clients should
// thus not rely on a given document having the same number between sessions.
//
// thread-safety
//
// NOTE: IndexReader instances are completely thread safe, meaning multiple
// threads can call any of its methods, concurrently. If your application requires external
// synchronization, you should not synchronize on the <code>IndexReader</code> instance; use
// your own (non-Lucene) objects instead.
type LeafReader interface {
	IndexReader
}
