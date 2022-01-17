package analysis

import "io"

// An Analyzer builds TokenStreams, which analyze text. It thus represents a policy for extracting
// index terms from text.
type Analyzer interface {
	io.Closer
}
