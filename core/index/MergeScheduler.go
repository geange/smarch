package index

import (
	. "github.com/geange/smarch/core/store"
	. "github.com/geange/smarch/core/util"
	"io"
)

// MergeScheduler Expert: IndexWriter uses an instance implementing this interface to execute the merges
// selected by a MergePolicy. The default MergeScheduler is ConcurrentMergeScheduler.
type MergeScheduler interface {
	io.Closer

	// Merge Run the merges provided by MergeSource#getNextMerge().
	Merge(mergeSource MergeSource, trigger MergeTriggerType) error

	// WrapForMerge Wraps the incoming Directory so that we can merge-throttle it using
	// RateLimitedIndexOutput.
	WrapForMerge(merge OneMerge, in Directory)

	//IndexWriter calls this on init.
	initialize(infoStream InfoStream, directory Directory) error

	// Verbose Returns true if infoStream messages are enabled. This method is usually used in conjunction
	// with #message(String):
	Verbose() bool

	// Message Outputs the given message - this method assumes #verbose()} was called and returned true.
	Message(message string)
}

// MergeSource Provides access to new merges and executes the actual merge
type MergeSource interface {
	// GetNextMerge The MergeScheduler calls this method to retrieve the next merge requested by the MergePolicy
	GetNextMerge() OneMerge

	// OnMergeFinished Does finishing for a merge.
	OnMergeFinished(merge OneMerge)

	// HasPendingMerges Expert: returns true if there are merges waiting to be scheduled.
	HasPendingMerges() bool

	// Merge Merges the indicated segments, replacing them in the stack with a single segment.
	Merge(merge OneMerge) error
}
