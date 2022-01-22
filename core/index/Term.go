package index

import (
	"bytes"
	"strings"
)

// A Term represents a word from text. This is the unit of search. It is composed of two elements,
// the text of the word, as a string, and the name of the field that the text occurred in.
//
// Note that terms may represent more than words from text fields, but also things like dates,
// email addresses, urls, etc.
type Term struct {
	field string
	bytes []byte
}

func NewTerm(fld string, v interface{}) *Term {
	term := &Term{
		field: fld,
		bytes: nil,
	}

	switch v.(type) {
	case string:
		term.bytes = []byte(v.(string))
	case []byte:
		obj := v.([]byte)
		buf := make([]byte, len(obj))
		copy(buf, obj)
		term.bytes = buf
	case *bytes.Buffer:
		obj := v.(*bytes.Buffer)
		term.bytes = obj.Bytes()
	default:
		term.bytes = make([]byte, 0)
	}

	return term
}

// Field Returns the field of this term. The field indicates the part of a document which this term came from.
func (t *Term) Field() string {
	return t.field
}

// Text Returns the text of this term. In the case of words, this is simply the text of the word. In
// the case of dates and other types, this is an encoding of the object as a string.
func (t *Term) Text() string {
	return string(t.bytes)
}

// Bytes Returns the bytes of this term, these should not be modified.
func (t *Term) Bytes() []byte {
	return t.bytes
}

// Set Resets the field and text of a Term.
//
// WARNING: the provided BytesRef is not copied, but used directly. Therefore the bytes should
// not be modified after construction, for example, you should clone a copy rather than pass
// reused bytes from a TermsEnum.
func (t *Term) Set(fld string, bytes []byte) {
	t.field = fld
	t.bytes = bytes
}

// CompareTo Compares two terms, returning a negative integer if this term belongs before the argument,
// zero if this term is equal to the argument, and a positive integer if this term belongs after the argument.
// The ordering of terms is first by field, then by text.
func (t *Term) CompareTo(other *Term) int {
	if t.field == other.field {
		return bytes.Compare(t.bytes, other.bytes)
	}

	return strings.Compare(t.field, other.field)
}
