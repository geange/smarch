package index

// FlushPolicy controls when segments are flushed from a RAM resident internal
// data-structure to the IndexWriters Directory.
//
// <p>Segments are traditionally flushed by:
//
// <ul>
//   <li>RAM consumption - configured via IndexWriterConfig#setRAMBufferSizeMB(double)
//   <li>Number of RAM resident documents - configured via
//       IndexWriterConfig#setMaxBufferedDocs(int)
// </ul>
//
// IndexWriter consults the provided FlushPolicy to control the flushing process.
// The policy is informed for each added or updated document as well as for each delete term. Based
// on the FlushPolicy, the information provided via DocumentsWriterPerThread and
// DocumentsWriterFlushControl, the FlushPolicy decides if a
// DocumentsWriterPerThread needs flushing and mark it as flush-pending via
// DocumentsWriterFlushControl#setFlushPending, or if deletes need to be applied.
type FlushPolicy interface {

	// OnDelete Called for each delete term. If this is a delete triggered due to an update the given
	// DocumentsWriterPerThread is non-null.
	//
	// Note: This method is called synchronized on the given DocumentsWriterFlushControl
	// and it is guaranteed that the calling thread holds the lock on the given
	// DocumentsWriterPerThread
	OnDelete(control *DocumentsWriterFlushControl, perThread *DocumentsWriterPerThread)

	// OnUpdate Called for each document update on the given DocumentsWriterPerThread's
	// DocumentsWriterPerThread.
	//
	// Note: This method is called synchronized on the given  DocumentsWriterFlushControl
	// and it is guaranteed that the calling thread holds the lock on the given
	// DocumentsWriterPerThread
	OnUpdate(control *DocumentsWriterFlushControl, perThread *DocumentsWriterPerThread)

	// OnInsert Called for each document addition on the given DocumentsWriterPerThreads
	// DocumentsWriterPerThread.
	//
	// Note: This method is synchronized by the given DocumentsWriterFlushControl and it is
	// guaranteed that the calling thread holds the lock on the given DocumentsWriterPerThread
	OnInsert(control *DocumentsWriterFlushControl, perThread *DocumentsWriterPerThread)

	// Init Called by DocumentsWriter to initialize the FlushPolicy
	Init(indexWriterConfig *LiveIndexWriterConfig)

	// FindLargestNonPendingWriter Returns the current most RAM consuming non-pending
	// DocumentsWriterPerThread with at least one indexed document.
	//
	// This method will never return null
	FindLargestNonPendingWriter(control *DocumentsWriterFlushControl, perThread *DocumentsWriterPerThread) *DocumentsWriterPerThread
}
