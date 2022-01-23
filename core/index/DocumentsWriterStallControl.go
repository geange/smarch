package index

// DocumentsWriterStallControl Controls the health status of a DocumentsWriter sessions. This class used to block
// incoming indexing threads if flushing significantly slower than indexing to ensure the
// DocumentsWriters healthiness. If flushing is significantly slower than indexing the net memory
// used within an IndexWriter session can increase very quickly and easily exceed the JVM's
// available memory.
// *
// <p>To prevent OOM Errors and ensure IndexWriter's stability this class blocks incoming threads
// from indexing once 2 x number of available DocumentsWriterPerThreads in
// DocumentsWriterPerThreadPool} is exceeded. Once flushing catches up and the number of flushing
// DWPT is equal or lower than the number of active DocumentsWriterPerThreads threads are
// released and can continue indexing.
type DocumentsWriterStallControl struct {
}
