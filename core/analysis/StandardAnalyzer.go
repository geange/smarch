package analysis

import (
	"io"

	. "github.com/geange/smarch/core/util"
)

// StandardAnalyzer Filters StandardTokenizer with LowerCaseFilter and StopFilter, using a
// configurable list of stop words.
type StandardAnalyzer struct {
	maxTokenLength int
}

func (s *StandardAnalyzer) Close() error {
	//TODO implement me
	panic("implement me")
}

func (s *StandardAnalyzer) CreateComponents(fieldName string) *TokenStreamComponents {
	//TODO implement me
	panic("implement me")
}

func (s *StandardAnalyzer) Normalize(fieldName string, in TokenStream) (TokenStream, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StandardAnalyzer) TokenStreamV1(fieldName string, reader io.Reader) (TokenStream, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StandardAnalyzer) TokenStreamV2(fieldName string, text string) (TokenStream, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StandardAnalyzer) NormalizeV1(fieldName string, text string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StandardAnalyzer) InitReader(fieldName string, reader io.Reader) io.Reader {
	//TODO implement me
	panic("implement me")
}

func (s *StandardAnalyzer) InitReaderForNormalization(fieldName string, reader io.Reader) io.Reader {
	//TODO implement me
	panic("implement me")
}

func (s *StandardAnalyzer) AttributeFactory(fieldName string) AttributeFactory {
	//TODO implement me
	panic("implement me")
}

func (s *StandardAnalyzer) GetPositionIncrementGap(fieldName string) int {
	//TODO implement me
	panic("implement me")
}

func (s *StandardAnalyzer) GetOffsetGap(fieldName string) int {
	//TODO implement me
	panic("implement me")
}

func (s *StandardAnalyzer) GetReuseStrategy() ReuseStrategy {
	//TODO implement me
	panic("implement me")
}
