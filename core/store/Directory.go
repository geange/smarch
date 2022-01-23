package store

import "io"

type Directory interface {
	io.Closer

	// ListAll Returns names of all files stored in this directory. The output must be in sorted (UTF-16,
	// java's String#compareTo) order.
	ListAll() ([]string, error)

	// DeleteFile Removes an existing file in the directory.
	// This method must throw either NoSuchFileException or FileNotFoundException
	// if name points to a non-existing file.
	DeleteFile(name string) error

	// FileLength Returns the byte length of a file in the directory.
	// This method must throw either NoSuchFileException or FileNotFoundException
	// if code name points to a non-existing file.
	// name: the name of an existing file.
	FileLength(name string) (int64, error)

	// CreateOutput Creates a new, empty file in the directory and returns an IndexOutput instance for
	// appending data to this file.
	// This method must throw java.nio.file.FileAlreadyExistsException if the file already
	// exists.
	CreateOutput(name string, context *IOContext) (IndexOutput, error)

	// CreateTempOutput Creates a new, empty, temporary file in the directory and returns an IndexOutput
	// instance for appending data to this file.
	// The temporary file name (accessible via IndexOutput#getName()) will start with
	// prefix, end with suffix and have a reserved file extension .tmp.
	CreateTempOutput(name string, context *IOContext) (IndexOutput, error)

	// Sync Ensures that any writes to these files are moved to stable storage (made durable).
	// Lucene uses this to properly commit changes to the index, to prevent a machine/OS crash from
	// corrupting the index.
	Sync(names []string) error

	// SyncMetaData Ensures that directory metadata, such as recent file renames, are moved to stable storage.
	SyncMetaData() error

	// Rename Renames source file to dest file where dest must not already exist in
	// the directory.
	//
	// It is permitted for this operation to not be truly atomic, for example both source
	// and dest can be visible temporarily in #listAll(). However, the implementation
	// of this method must ensure the content of dest appears as the entire source
	// atomically. So once dest is visible for readers, the entire content of previous
	// source is visible.
	//
	// This method is used by IndexWriter to publish commits.
	Rename(source, dest string) error

	// OpenInput Opens a stream for reading an existing file.
	//
	// This method must throw either NoSuchFileException or FileNotFoundException
	// if name points to a non-existing file.
	OpenInput(name string, context *IOContext) (IndexOutput, error)

	// OpenChecksumInput Opens a checksum-computing stream for reading an existing file.
	//
	// This method must throw either NoSuchFileException or FileNotFoundException
	// if name points to a non-existing file.
	OpenChecksumInput(name string, context *IOContext) (ChecksumIndexInput, error)

	// ObtainLock Acquires and returns a Lock for a file with the given name.
	ObtainLock(name string) (Lock, error)

	// CopyFrom Copies an existing src file from directory from to a non-existent file
	// dest in this directory.
	CopyFrom(from Directory, src, dest string, context *IOContext) error

	// EnsureOpen Ensures this directory is still open.
	EnsureOpen() error

	// GetPendingDeletions Returns a set of files currently pending deletion in this directory.
	GetPendingDeletions() (map[string]struct{}, error)

	// GetTempFileName Creates a file name for a temporary file. The name will start with prefix, end with
	// suffix and have a reserved file extension .tmp.
	GetTempFileName(prefix, suffix string, counter int64) (string, error)
}
