package index

// DirectoryReader is an implementation of CompositeReader that can read indexes in a Directory.
//
// DirectoryReader instances are usually constructed with a call to one of the static
// open() methods, e.g. #open(Directory).
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
// synchronization, you should not synchronize on the IndexReader instance; use
// your own (non-Lucene) objects instead.
type DirectoryReader interface {
}
