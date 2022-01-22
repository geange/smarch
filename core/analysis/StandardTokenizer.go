package analysis

import "io"

type StandardTokenizer struct {
}

func (s *StandardTokenizer) Close() error {
	//TODO implement me
	panic("implement me")
}

func (s *StandardTokenizer) IncrementToken() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StandardTokenizer) End() error {
	//TODO implement me
	panic("implement me")
}

func (s *StandardTokenizer) Reset() error {
	//TODO implement me
	panic("implement me")
}

func (s *StandardTokenizer) SetReader(input io.Reader) error {
	//TODO implement me
	panic("implement me")
}

type STDTokenType int

const (
	STDTokenTypeAlphaNum       = STDTokenType(iota) // Alpha/numeric token type
	STDTokenTypeNum                                 // Numeric token type
	STDTokenTypeSoutheastAsian                      // Southeast Asian token type
	STDTokenTypeIdeographic                         // Ideographic token type
	STDTokenTypeHiragana                            // Hiragana token type
	STDTokenTypeKatakana                            // Katakana token type
	STDTokenTypeHangul                              // Hangul token type
	STDTokenTypeEmoji                               // Emoji token type.
)
