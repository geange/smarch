package index

type IndexWriterConfig struct {
}

// OpenModeType Specifies the open mode for IndexWriter.
type OpenModeType int

const (
	// OpenModeCreate Creates a new index or overwrites an existing one.
	OpenModeCreate = OpenModeType(iota)

	// OpenModeAppend Opens an existing index.
	OpenModeAppend

	// OpenModeCreateOrAppend Creates a new index if one does not exist,
	// otherwise it opens the index and documents will be appended.
	OpenModeCreateOrAppend
)
