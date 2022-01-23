package index

import . "github.com/geange/smarch/core/util"

// DocumentsWriterFlushControl This class controls DocumentsWriterPerThread flushing during indexing.
// It tracks the memory consumption per DocumentsWriterPerThread and uses a configured FlushPolicy
// to decide if a DocumentsWriterPerThread} must flush.
//
// In addition to the FlushPolicy the flush control might set certain
// DocumentsWriterPerThread} as flush pending iff a DocumentsWriterPerThread} exceeds the
// IndexWriterConfig#getRAMPerThreadHardLimitMB() to prevent address space exhaustion.
type DocumentsWriterFlushControl struct {
	hardMaxBytesPerDWPT int64
	activeBytes         int64
	flushBytes          int64

	numPending          int
	numDocsSinceStalled int
	flushDeletes        bool
	fullFlush           bool

	// only for assertion that we don't get stale DWPTs from the pool
	fullFlushMarkDone bool

	// The flushQueue is used to concurrently distribute DWPTs that are ready to be flushed ie. when a
	// full flush is in
	// progress. This might be triggered by a commit or NRT refresh. The trigger will only walk all
	// eligible DWPTs and
	// mark them as flushable putting them in the flushQueue ready for other threads (ie. indexing
	// threads) to help flushing
	flushQueue interface{}

	// only for safety reasons if a DWPT is close to the RAM limit
	//private final Queue<DocumentsWriterPerThread> flushQueue = new LinkedList<>();
	//private final Queue<DocumentsWriterPerThread> blockedFlushes = new LinkedList<>();
	//private final List<DocumentsWriterPerThread> flushingWriters = new ArrayList<>();

	// flushingWriters holds all currently flushing writers. There might be writers in this list that
	// are also in the flushQueue which means that writers in the flushingWriters list are not
	// necessarily
	// already actively flushing. They are only in the state of flushing and might be picked up in the
	// future by
	// polling the flushQueue

	maxConfiguredRamBuffer float64

	peakActiveBytes       int64 // only with assert
	peakFlushBytes        int64 // only with assert
	peakNetBytes          int64 // only with assert
	peakDelta             int64 // only with assert
	flushByRAMWasDisabled bool  // only with assert

	stallControl    *DocumentsWriterStallControl
	perThreadPool   *DocumentsWriterPerThreadPool
	flushPolicy     FlushPolicy
	closed          bool
	documentsWriter *DocumentsWriter
	config          *LiveIndexWriterConfig
	infoStream      InfoStream
}
