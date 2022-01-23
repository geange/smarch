package index

import (
	. "github.com/geange/smarch/core/search"
)

// PostingsEnum Iterates through the postings. NOTE: you must first call #nextDoc before using any of the
// per-doc methods.
type PostingsEnum interface {
	DocIdSetIterator

	// Freq Returns term frequency in the current document, or 1 if the field was indexed with
	// IndexOptions#DOCS. Do not call this before #nextDoc is first called, nor after
	// #nextDoc returns DocIdSetIterator#NO_MORE_DOCS.
	//   *
	// NOTE: if the PostingsEnum was obtain with #NONE, the result of this
	// method is undefined.
	Freq() (int, error)

	// NextPosition Returns the next position, or -1 if positions were not indexed. Calling this more than
	// #freq() times is undefined.
	NextPosition() (int, error)

	// StartOffset Returns start offset for the current position, or -1 if offsets were not indexed.
	StartOffset() (int, error)

	// EndOffset Returns end offset for the current position, or -1 if offsets were not indexed.
	EndOffset() (int, error)

	// GetPayload Returns the payload at this position, or null if no payload was indexed. You should not modify
	// anything (neither members of the returned BytesRef nor bytes in the byte[]).
	GetPayload() ([]byte, error)
}
