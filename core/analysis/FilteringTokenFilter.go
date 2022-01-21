package analysis

import . "github.com/geange/smarch/core/util"

// FilteringTokenFilter Abstract base class for TokenFilters that may remove tokens. You have to implement
// #accept and return a boolean if the current token should be preserved. #incrementToken
// uses this method to decide if a token should be passed to the caller.
type FilteringTokenFilter struct {
	TokenFilter

	posIncrAtt       PositionIncrementAttribute
	skippedPositions int
	accept           func() bool
}

func (f *FilteringTokenFilter) IncrementToken() (bool, error) {
	f.skippedPositions = 0
	for {
		ok, err := f.input.IncrementToken()
		if err != nil {
			return false, err
		}
		if !ok {
			break
		}

		if f.accept() {
			if f.skippedPositions != 0 {
				f.posIncrAtt.SetPositionIncrement(f.posIncrAtt.GetPositionIncrement() + f.skippedPositions)
			}
			return true, nil
		}

		f.skippedPositions += f.posIncrAtt.GetPositionIncrement()
	}

	return false, nil
}
