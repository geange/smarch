package analysis

import "io"

// A TokenStream enumerates the sequence of tokens, either from Fields of a
// Document or from query text.
//
// To make sure that filters and consumers know which attributes are available, the attributes must
// be added during instantiation. Filters and consumers are not required to check for availability
// of attributes in #incrementToken().
//
// You can find some example code for the new API in the analysis package level Javadoc.
//
// Sometimes it is desirable to capture a current state of a TokenStream, e.g., for
// buffering purposes (see CachingTokenFilter, TeeSinkTokenFilter). For this usecase
// AttributeSource#captureState and  AttributeSource#restoreState can be used.
//
// The TokenStream-API in Lucene is based on the decorator pattern. Therefore all
// non-abstract subclasses must be final or have at least a final implementation of
// #incrementToken}! This is checked when Java assertions are enabled.
type TokenStream interface {
	io.Closer

	// IncrementToken Consumers (i.e., IndexWriter) use this method to advance the stream to the next token.
	// Implementing classes must implement this method and update the appropriate
	// AttributeImpls with the attributes of the next token.
	//
	// The producer must make no assumptions about the attributes after the method has been
	// returned: the caller may arbitrarily change it. If the producer needs to preserve the state for
	// subsequent calls, it can use #captureState to create a copy of the current attribute
	// state.
	//
	// This method is called for every token of a document, so an efficient implementation is
	// crucial for good performance. To avoid calls to #addAttribute(Class) and
	// #getAttribute(Class), references to all AttributeImpls that this stream uses should be
	// retrieved during instantiation.
	//
	// To ensure that filters and consumers know which attributes are available, the attributes
	// must be added during instantiation. Filters and consumers are not required to check for
	// availability of attributes in #incrementToken().
	IncrementToken() (bool, error)

	// End This method is called by the consumer after the last token has been consumed, after
	// #incrementToken() returned false (using the new TokenStream API).
	// Streams implementing the old API should upgrade to use this feature.
	//
	// This method can be used to perform any end-of-stream operations, such as setting the final
	// offset of a stream. The final offset of a stream might differ from the offset of the last token
	// eg in case one or more whitespaces followed after the last token, but a WhitespaceTokenizer was
	// used.
	//
	// Additionally any skipped positions (such as those removed by a stopfilter) can be applied to
	// the position increment, or any adjustment of other attributes where the end-of-stream value may
	// be important.
	//
	// If you override this method, always call {@code super.end()}.
	End() error

	// Reset This method is called by a consumer before it begins consumption using IncrementToken().
	//
	// Resets this stream to a clean state. Stateful implementations must implement this method so
	// that they can be reused, just as if they had been created fresh.
	//
	// If you override this method, always call super.reset(), otherwise some internal
	// state will not be correctly reset (e.g., Tokenizer will throw
	// IllegalStateException on further usage).
	Reset() error
}
