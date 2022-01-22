package index

import (
	. "github.com/geange/smarch/core/util"
)

// TermsEnum Iterator to seek (#seekCeil(BytesRef), #seekExact(BytesRef)) or step through
// (#next} terms to obtain frequency information (#docFreq), PostingsEnum or
// PostingsEnum for the current term (#postings.
//
// Term enumerations are always ordered by BytesRef.compareTo, which is Unicode sort order if the
// terms are UTF-8 bytes. Each term in the enumeration is greater than the one before it.
//
// The TermsEnum is unpositioned when you first obtain it and you must first successfully call
// #next} or one of the seek methods.
type TermsEnum interface {
	// Attributes Returns the related attributes.
	Attributes() *AttributeSource

	// SeekExact Attempts to seek to the exact term, returning true if the term is found. If this returns false,
	// the enum is unpositioned. For some codecs, seekExact may be substantially faster than #seekCeil.
	//
	// return true if the term is found; return false if the enum is unpositioned.
	SeekExact(text []byte) (bool, error)

	// SeekCeil Seeks to the specified term, if it exists, or to the next (ceiling) term. Returns SeekStatus to
	// indicate whether exact term was found, a different term was found, or EOF was hit. The target
	// term may be before or after the current term. If this returns SeekStatus.END, the enum is
	// unpositioned.
	SeekCeil(text []byte) (SeekStatusType, error)

	// SeekExactOrd Seeks to the specified term by ordinal (position) as previously returned by #ord}. The
	// target ord may be before or after the current ord, and must be within bounds.
	SeekExactOrd(ord int64) error

	// SeekExactTerm Expert: Seeks a specific position by TermState} previously obtained from
	// #termState()}. Callers should maintain the TermState} to use this method. Low-level
	// implementations may position the TermsEnum without re-seeking the term dictionary.
	//
	// Seeking by TermState} should only be used iff the state was obtained from the same
	// TermsEnum} instance.
	//
	// NOTE: Using this method with an incompatible TermState} might leave this
	// TermsEnum} in undefined state. On a segment level TermState} instances are compatible
	// only iff the source and the target TermsEnum} operate on the same field. If operating on
	// segment level, TermState instances must not be used across segments.
	//
	// NOTE: A seek by TermState} might not restore the AttributeSource}'s state.
	// AttributeSource} states must be maintained separately if this method is used.
	SeekExactTerm(term []byte, state SeekStatusType) error

	// Term Returns current term. Do not call this when the enum is unpositioned.
	Term() ([]byte, error)

	// Ord Returns ordinal position for current term. This is an optional method (the codec may throw
	// UnsupportedOperationException}). Do not call this when the enum is unpositioned.
	Ord() (int64, error)

	// DocFreq Returns the number of documents containing the current term. Do not call this when the enum is
	// unpositioned. SeekStatus#END}.
	DocFreq() (int, error)

	// TotalTermFreq Returns the total number of occurrences of this term across all documents (the sum of the
	// freq() for each doc that has this term). Note that, like other term measures, this measure does
	// not take deleted documents into account.
	TotalTermFreq() (int64, error)

	// Postings Get PostingsEnum for the current term. Do not call this when the enum is unpositioned.
	// This method will not return null.
	//
	// NOTE: the returned iterator may return deleted documents, so deleted documents have
	// to be checked on top of the PostingsEnum.
	//
	// Use this method if you only require documents and frequencies, and do not need any proximity
	// data. This method is equivalent to #postings(PostingsEnum, int) postings(reuse,
	// PostingsEnum.FREQS)
	Postings(reuse PostingsEnum) (PostingsEnum, error)

	// PostingsFlags Get PostingsEnum} for the current term, with control over whether freqs, positions,
	// offsets or payloads are required. Do not call this when the enum is unpositioned. This method
	// will not return null.
	//
	// NOTE: the returned iterator may return deleted documents, so deleted documents have
	// to be checked on top of the PostingsEnum}.
	PostingsFlags(reuse PostingsEnum, flags int) (PostingsEnum, error)

	// Impacts Return a ImpactsEnum.
	Impacts(flags int) (ImpactsEnum, error)

	// TermState Expert: Returns the TermsEnums internal state to position the TermsEnum without re-seeking the
	// term dictionary.
	//
	// NOTE: A seek by TermState} might not capture the AttributeSource}'s state.
	// Callers must maintain the AttributeSource} states separately
	TermState() (TermState, error)
}

type SeekStatusType int

const (
	// SeekStatusEnd The term was not found, and the end of iteration was hit.
	SeekStatusEnd = SeekStatusType(iota)

	// SeekStatusFound The precise term was found.
	SeekStatusFound

	// SeekStatusNotFound A different term was found after the requested term
	SeekStatusNotFound
)
