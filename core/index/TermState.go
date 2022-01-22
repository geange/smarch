package index

// TermState Encapsulates all required internal state to position the associated TermsEnum} without
// re-seeking.
type TermState interface {
	// CopyFrom Copies the content of the given TermState} to this instance
	CopyFrom(other TermState)

	Clone() TermState
}
