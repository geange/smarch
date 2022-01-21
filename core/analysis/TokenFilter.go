package analysis

type TokenFilter struct {
	//The source of tokens for this filter.
	input TokenStream
}

func (t *TokenFilter) Close() error {
	return t.input.Close()
}

func (t *TokenFilter) IncrementToken() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TokenFilter) End() error {
	return t.input.End()
}

func (t *TokenFilter) Reset() error {
	return t.input.Reset()
}
