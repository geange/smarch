package util

type OffsetAttribute interface {
	Attribute

	// StartOffset Returns this Token's starting offset, the position of the first character corresponding to this
	// token in the source text.
	//
	// Note that the difference between #endOffset() and startOffset() may not
	// be equal to termText.length(), as the term text may have been altered by a stemmer or some
	// other filter.
	StartOffset() int

	// SetOffset Set the starting and ending offset.
	SetOffset(startOffset, endOffset int)

	// EndOffset Returns this Token's ending offset, one greater than the position of the last character
	// corresponding to this token in the source text. The length of the token in the source text is (
	// endOffset() - #startOffset()).
	EndOffset() int
}
