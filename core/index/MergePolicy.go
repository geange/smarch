package index

import . "github.com/geange/smarch/core/util"

// MergePolicy Expert: a MergePolicy determines the sequence of primitive merge operations.
//
// Whenever the segments in an index have been altered by IndexWriter, either the
// addition of a newly flushed segment, addition of many segments from addIndexescalls, or a
// previous merge that may now need to cascade, IndexWriter invokes  #findMerges to
// give the MergePolicy a chance to pick merges that are now required. This method returns a
// MergeSpecification instance describing the set of merges that should be done, or null if no
// merges are necessary. When IndexWriter.forceMerge is called, it calls
// #findForcedMerges(SegmentInfos, int, Map, MergeContext) and the MergePolicy should then return
// the necessary merges.
//
// Note that the policy can return more than one merge at a time. In this case, if the writer is
// using  SerialMergeScheduler, the merges will be run sequentially but if it is using
// ConcurrentMergeScheduler they will be run concurrently.
//
// The default MergePolicy is  TieredMergePolicy.
type MergePolicy interface {
}

// OneMerge provides the information necessary to perform an individual primitive merge operation,
// resulting in a single new segment. The merge spec includes the subset of segments to be merged
// as well as whether the new segment should use the compound file format.
type OneMerge interface {
	// MergeInit Called by IndexWriter after the merge started and from the thread that will be
	// executing the merge.
	MergeInit() error

	// MergeFinished Called by IndexWriter after the merge is done and all readers have been closed.
	MergeFinished(success, segmentDropped bool) error

	//Closes this merge and releases all merge readers
	Close(success, segmentDropped bool)
}

// MergeContext This interface represents the current context of the merge selection process. It allows to
// access real-time information like the currently merging segments or how many deletes a segment
// would claim back if merged. This context might be stateful and change during the execution of a
// merge policy's selection processes.
type MergeContext interface {
	// NumDeletesToMerge Returns the number of deleted documents in the given segments.
	NumDeletesToMerge(info *SegmentCommitInfo) (int, error)

	// GetInfoStream Returns the info stream that can be used to log messages
	GetInfoStream() InfoStream

	// GetMergingSegments Returns an unmodifiable set of segments that are currently merging.
	GetMergingSegments() map[string]*SegmentCommitInfo
}

type MergeReader struct {
	reader       SegmentReader
	hardLiveDocs Bits
}
