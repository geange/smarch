package store

type DataOutput interface {
	WriteByte(b byte) error
	Write(b []byte) error
	WriteString(s string) error
	WriteUint16(i uint16) error
	WriteUint32(i uint32) error
	WriteUint64(i uint64) error
	WriteVInt(i uint64) error
	WriteZInt32(i uint64) error
	WriteZInt64(i uint64) error
	WriteMapOfStrings(values map[string]string) error
	WriteSetOfStrings(values map[string]struct{}) error
	CopyBytes(input DataInput, size int64) error
}
