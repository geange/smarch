package index

import (
	. "github.com/geange/smarch/core/store"
)

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
	// ListCommits Returns all commit points that exist in the Directory. Normally, because the default is
	// KeepOnlyLastCommitDeletionPolicy, there would be only one commit point. But if you're using a
	// custom IndexDeletionPolicy then there could be many commits. Once you have a given
	// commit, you can open a reader on it by calling DirectoryReader#open(IndexCommit) There
	// must be at least one commit in the Directory, else this method throws
	// IndexNotFoundException. Note that if a commit is in progress while this method is running,
	// that commit may or may not be returned.
	//
	// return a sorted list of IndexCommits, from oldest to latest.
	ListCommits(dir Directory) ([]IndexCommit, error)

	// IndexExists Returns true if an index likely exists at the specified directory. Note that if a
	// corrupt index exists, or if an index in the process of committing
	IndexExists(dir Directory) (bool, error)

	// Directory Returns the directory this index resides in.
	Directory() Directory

	// DoOpenIfChangedV1 Implement this method to support #openIfChanged(DirectoryReader). If this reader does
	// not support reopen, return null, so client code is happy. This should be consistent
	// with #isCurrent (should always return true) if reopen is not supported.
	DoOpenIfChangedV1() (DirectoryReader, error)

	// DoOpenIfChangedV2 Implement this method to support #openIfChanged(DirectoryReader,IndexCommit). If this
	// reader does not support reopen from a specific IndexCommit, throw
	// UnsupportedOperationException.
	DoOpenIfChangedV2(commit IndexCommit) (DirectoryReader, error)

	// DoOpenIfChanged Implement this method to support #openIfChanged(DirectoryReader,IndexWriter,boolean).
	// If this reader does not support reopen from IndexWriter, throw UnsupportedOperationException.
	DoOpenIfChanged(writer *IndexWriter, applyAllDeletes bool) (DirectoryReader, error)

	// GetVersion Version number when this IndexReader was opened.
	GetVersion() int64

	// IsCurrent Check whether any new changes have occurred to the index since this reader was opened.
	//
	// If this reader was created by calling #open, then this method checks if any further
	// commits (see IndexWriter#commit) have occurred in the directory.
	//
	// If instead this reader is a near real-time reader (ie, obtained by a call to
	// DirectoryReader#open(IndexWriter), or by calling #openIfChanged on a near real-time
	// reader), then this method checks if either a new commit has occurred, or any new uncommitted
	// changes have taken place via the writer. Note that even if the writer has only performed
	// merging, this method will still return false.
	//
	// In any event, if this returns false, you should call #openIfChanged to get a new
	// reader that sees the changes.
	IsCurrent() (bool, error)

	// GetIndexCommit Expert: return the IndexCommit that this reader has opened.
	GetIndexCommit() (IndexCommit, error)
}
