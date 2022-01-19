package store

type DataInput interface {
	ReadByte() (byte, error)
	ReadBytes(n int) ([]byte, error)
	ReadUint16() (uint16, error)
	ReadUint32() (uint32, error)
	ReadUint64() (uint64, error)
	ReadVInt() (uint64, error)
	ReadZInt32() (uint64, error)
	ReadZInt64() (uint64, error)
	ReadString() (string, error)
	ReadMapOfStrings() (map[string]string, error)
	ReadSetOfStrings() (map[string]struct{}, error)
	SkipBytes(size int) error
}
