package index

// DocumentsWriterDeleteQueue is a non-blocking linked pending deletes queue. In contrast to
// other queue implementation we only maintain the tail of the queue. A delete queue is always used
// in a context of a set of DWPTs and a global delete pool. Each of the DWPT and the global pool
// need to maintain their 'own' head of the queue (as a DeleteSlice instance per
// DocumentsWriterPerThread). The difference between the DWPT and the global pool is that the DWPT
// starts maintaining a head once it has added its first document since for its segments private
// deletes only the deletes after that document are relevant. The global pool instead starts
// maintaining the head once this instance is created by taking the sentinel instance as its initial
// head.
//
// Since each DeleteSlice maintains its own head and the list is only single linked the
// garbage collector takes care of pruning the list for us. All nodes in the list that are still
// relevant should be either directly or indirectly referenced by one of the DWPT's private
// DeleteSlice or by the global BufferedUpdates slice.
//
// Each DWPT as well as the global delete pool maintain their private DeleteSlice instance. In
// the DWPT case updating a slice is equivalent to atomically finishing the document. The slice
// update guarantees a "happens before" relationship to all other updates in the same indexing
// session. When a DWPT updates a document it:
//
// <ol>
//   <li>consumes a document and finishes its processing
//   <li>updates its private DeleteSlice either by calling #updateSlice(DeleteSlice)
//       or #add(Node, DeleteSlice) (if the document has a delTerm)
//   <li>applies all deletes in the slice to its private BufferedUpdates and resets it
//   <li>increments its internal document id
// </ol>
//
// The DWPT also doesn't apply its current documents delete term until it has updated its delete
// slice which ensures the consistency of the update. If the update fails before the DeleteSlice
// could have been updated the deleteTerm will also not be added to its private deletes neither to
// the global deletes.
type DocumentsWriterDeleteQueue struct {

	// the current end (latest delete operation) in the delete queue:
	tail interface{}

	closed bool

	// Used to record deletes against all prior (already written to disk) segments. Whenever any
	// segment flushes, we bundle up this set of deletes and insert into the buffered updates stream
	// before the newly flushed segment(s).
	// 用于记录对所有先前（已写入磁盘）段的删除。 每当任何段刷新时，我们都会捆绑这组删除，并在新刷新的段之前插入缓冲的更新流。
	//globalSlice
}

type DeleteSlice struct {
	// No need to be volatile, slices are thread captive (only accessed by one thread)!
	sliceHead *Node // we don't apply this one
	sliceTail *Node
}

type Node struct {
	Item interface{}
	Next *Node
}
