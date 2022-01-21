package analysis

import "io"

// A Tokenizer is a TokenStream whose input is a Reader.
//
// This is an abstract class; subclasses must override {@link #incrementToken()}
//
// NOTE: Subclasses overriding IncrementToken() must call {@link
// AttributeSource#clearAttributes()} before setting attributes.
type Tokenizer interface {
	TokenStream

	// SetReader Expert: Set a new reader on the Tokenizer. Typically,
	// an analyzer (in its tokenStream method) will use this to re-use
	// a previously created tokenizer.
	SetReader(input io.Reader) error
}
