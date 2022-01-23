package analysis

import (
	"bytes"
	. "github.com/geange/smarch/core/util"
	"io"
)

// An Analyzer builds TokenStreams, which analyze text. It thus represents a policy for extracting
// index terms from text.
//
// In order to define what analysis is done, subclasses must define their
// TokenStreamComponents TokenStreamComponents in #createComponents(String). The components
// are then reused in each call to #tokenStream(String, Reader).
type Analyzer interface {
	io.Closer

	// CreateComponents Creates a new TokenStreamComponents instance for this analyzer.
	CreateComponents(fieldName string) *TokenStreamComponents

	// Normalize Wrap the given TokenStream in order to apply normalization filters. The default
	// implementation returns the TokenStream as-is. This is used by #normalize(String, String).
	Normalize(fieldName string, in TokenStream) (TokenStream, error)

	// TokenStreamV1 Returns a TokenStream suitable for fieldName, tokenizing the contents of
	// reader.
	//
	// This method uses #createComponents(String) to obtain an instance of
	// TokenStreamComponents. It returns the sink of the components and stores the components
	// internally. Subsequent calls to this method will reuse the previously stored components after
	// resetting them through TokenStreamComponents#setReader(Reader).
	//
	// NOTE: After calling this method, the consumer must follow the workflow described in
	// TokenStream to properly consume its contents. See the org.apache.lucene.analysis
	// Analysis package documentation for some examples demonstrating this.
	//
	// NOTE: If your data is available as a String, use #tokenStream(String,
	// String) which reuses a StringReader-like instance internally.
	TokenStreamV1(fieldName string, reader io.Reader) (TokenStream, error)

	// TokenStreamV2 Returns a TokenStream suitable for fieldName, tokenizing the contents of text.
	//
	// This method uses #createComponents(String) to obtain an instance of
	// TokenStreamComponents. It returns the sink of the components and stores the components
	// internally. Subsequent calls to this method will reuse the previously stored components after
	// resetting them through TokenStreamComponents#setReader(Reader).
	//
	// NOTE: After calling this method, the consumer must follow the workflow described in
	// TokenStream to properly consume its contents. See the
	// Analysis package documentation for some examples demonstrating this.
	TokenStreamV2(fieldName string, text string) (TokenStream, error)

	// NormalizeV1 Normalize a string down to the representation that it would have in the index.
	//
	// This is typically used by query parsers in order to generate a query on a given term,
	// without tokenizing or stemming, which are undesirable if the string to analyze is a partial
	// word (eg. in case of a wildcard or fuzzy query).
	//
	// This method uses #initReaderForNormalization(String, Reader) in order to apply
	// necessary character-level normalization and then #normalize(String, TokenStream) in
	// order to apply the normalizing token filters.
	NormalizeV1(fieldName string, text string) ([]byte, error)

	// InitReader Override this if you want to add a CharFilter chain.
	// The default implementation returns reader unchanged.
	InitReader(fieldName string, reader io.Reader) io.Reader

	// InitReaderForNormalization Wrap the given Reader with CharFilters that make sense for normalization. This
	// is typically a subset of the CharFilters that are applied in #initReader(String,
	// Reader). This is used by #normalize(String, String).
	InitReaderForNormalization(fieldName string, reader io.Reader) io.Reader

	// AttributeFactory Return the AttributeFactory to be used for #tokenStream analysis and
	// #normalize(String, String) normalization on the given FieldName. The default
	// implementation returns TokenStream#DEFAULT_TOKEN_ATTRIBUTE_FACTORY.
	AttributeFactory(fieldName string) AttributeFactory

	// GetPositionIncrementGap Invoked before indexing a IndexableField instance if terms have already been added to that
	// field. This allows custom analyzers to place an automatic position increment gap between
	// IndexbleField instances using the same field name. The default value position increment gap is
	// 0. With a 0 position increment gap and the typical default token position increment of 1, all
	// terms in a field, including across IndexableField instances, are in successive positions,
	// allowing exact PhraseQuery matches, for instance, across IndexableField instance boundaries.
	GetPositionIncrementGap(fieldName string) int

	// GetOffsetGap Just like #getPositionIncrementGap, except for Token offsets instead. By default this
	// returns 1. This method is only called if the field produced at least one token for indexing.
	GetOffsetGap(fieldName string) int

	// GetReuseStrategy Returns the used ReuseStrategy.
	GetReuseStrategy() ReuseStrategy
}

// TokenStreamComponents This class encapsulates the outer components of a token stream.
// It provides access to the source (a Reader Consumer and the outer end (sink),
// an instance of TokenFilter which also serves as the TokenStream returned by
// Analyzer#tokenStream(String, Reader).
type TokenStreamComponents struct {

	//Original source of the tokens.
	source func(reader io.Reader) error

	// Sink tokenstream, such as the outer tokenfilter decorating the chain.
	// This can be the source if there are no filters.
	sink TokenStream

	//Internal cache only used by Analyzer#tokenStream(String, String).
	reusableStringReader bytes.Reader
}

// ReuseStrategy Strategy defining how TokenStreamComponents are reused per call to
// Analyzer#tokenStream(String, java.io.Reader).
type ReuseStrategy interface {

	// GetReusableComponents Gets the reusable TokenStreamComponents for the field with the given name.
	GetReusableComponents(analyzer Analyzer, fieldName string) *TokenStreamComponents

	// SetReusableComponents Stores the given TokenStreamComponents as the reusable components for the field with the give name.
	SetReusableComponents(analyzer Analyzer, fieldName string, components *TokenStreamComponents)

	// GetStoredValue Returns the currently stored value.
	GetStoredValue(analyzer Analyzer) interface{}

	// SetStoredValue Sets the stored value.
	SetStoredValue(analyzer Analyzer, storedValue interface{})
}
