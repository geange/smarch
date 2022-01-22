package index

// BaseCompositeReader Base class for implementing CompositeReaders based on an array of sub-readers. The
// implementing class has to add code for correctly refcounting and closing the sub-readers.
//
// User code will most likely use MultiReader to build a composite reader on a set of
// sub-readers (like several DirectoryReaders).
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
type BaseCompositeReader struct {
	subReaders []IndexReader

	/** A comparator for sorting sub-readers */
	//protected final Comparator<R> subReadersSorter;

	// 1st docno for each reader
	starts []int

	maxDoc int

	// computed lazily default -1
	numDocs int64

	// List view solely for #getSequentialSubReaders(), for effectiveness the array is used internally.
	subReadersList []IndexReader
}

func (b BaseCompositeReader) Close() error {
	//TODO implement me
	panic("implement me")
}
