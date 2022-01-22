package index

import . "github.com/geange/smarch/core/store"

// IndexCommit Expert: represents a single commit into an index as seen by the IndexDeletionPolicy or IndexReader.
//
// Changes to the content of an index are made visible only after the writer who made that change
// commits by writing a new segments file (segments_N). This point in time, when the
// action of writing of a new segments file to the directory is completed, is an index commit.
//
// Each index commit point has a unique segments file associated with it. The segments file
// associated with a later index commit point would have a larger N.
type IndexCommit interface {

	// GetSegmentsFileName Get the segments file (segments_N) associated with this commit point.
	GetSegmentsFileName() string

	// GetFileNames Returns all index files referenced by this commit point.
	GetFileNames() ([]string, error)

	// GetDirectory Returns the Directory for the index.
	GetDirectory() Directory

	// Delete this commit point. This only applies when using the commit point in the context of
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
	GetSegmentCount() int

	// GetGeneration Returns the generation (the _N in segments_N) for this IndexCommit
	GetGeneration() int

	// GetUserData Returns userData, previously passed to IndexWriter#setLiveCommitData(Iterable) for this commit.
	GetUserData() (map[string]string, error)
}
