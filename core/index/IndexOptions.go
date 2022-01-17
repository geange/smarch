package index

// IndexOptions Controls how much information is stored in the postings lists.
type IndexOptions int

const (
	// IdxOptNone Not indexed
	IdxOptNone = IndexOptions(iota)

	// IdxOptDocs Only documents are indexed: term frequencies and positions are omitted. Phrase and other
	// positional queries on the field will throw an exception, and scoring will behave as if any term
	// in the document appears only once.
	IdxOptDocs

	// IdxOptDocsAndFreqs Only documents and term frequencies are indexed: positions are omitted. This enables normal
	// scoring, except Phrase and other positional queries will throw an exception.
	IdxOptDocsAndFreqs

	// IdxOptDocsAndFreqsAndPositions Indexes documents, frequencies and positions.
	// This is a typical default for full-text search:
	// full scoring is enabled and positional queries are supported.
	IdxOptDocsAndFreqsAndPositions

	// IdxOptDocsAndFreqsAndPositionsAndOffsets Indexes documents, frequencies, positions and offsets.
	// Character offsets are encoded alongside the positions.
	IdxOptDocsAndFreqsAndPositionsAndOffsets
)
