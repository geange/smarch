package document

/**

import (
	. "github.com/geange/smarch/core/store"
)

type IndexCommit interface {
	// GetSegmentsFileName Get the segments file (segments_N) associated with this commit point.
	GetSegmentsFileName() string

	// GetFileNames Returns all index files referenced by this commit point.
	GetFileNames() ([]string, error)

	// GetDirectory Returns the Directory for the index.
	GetDirectory() Directory

	//Delete this commit point. This only applies when using the commit point in the context of
	// IndexWriter's IndexDeletionPolicy.
	//
	// Upon calling this, the writer is notified that this commit point should be deleted.
	//
	// Decision that a commit-point should be deleted is taken by the IndexDeletionPolicy
	// in effect and therefore this should only be called by its IndexDeletionPolicy#onInit
	// onInit() or IndexDeletionPolicy#onCommit onCommit() methods.
	Delete()

	// IsDeleted Returns true if this commit should be deleted; this is only used by IndexWriter after
	// invoking the IndexDeletionPolicy.
	IsDeleted() bool

	// GetSegmentCount Returns number of segments referenced by this commit.
	GetSegmentCount() interface{}

	//Returns the generation (the _N in segments_N) for this IndexCommit
	getGeneration() int

	// Returns userData, previously passed to IndexWriter#setLiveCommitData(Iterable) for this
	// commit. Map is map[string]string.
	getUserData() (map[string]string, error)

	// Package-private API for IndexWriter to init from a commit-point pulled from an NRT or non-NRT
	// reader.
	getReader()
}
*/
