package store

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"

	"github.com/gogo/protobuf/proto"
)

type FSIndexOutput struct {
	name           string
	file           *os.File
	byteOrder      binary.ByteOrder
	bytesWritten   int64
	flushedOnClose bool
}

func (f *FSIndexOutput) Close() error {
	return f.file.Close()
}

func (f *FSIndexOutput) WriteByte(b byte) error {
	return f.Write([]byte{b})
}

func (f *FSIndexOutput) Write(b []byte) error {
	_, err := f.file.Write(b)
	f.bytesWritten += int64(len(b))
	return err
}

func (f *FSIndexOutput) WriteString(s string) error {
	size := uint64(len(s))
	if err := f.WriteVInt(size); err != nil {
		return err
	}

	return f.Write([]byte(s))
}

func (f *FSIndexOutput) WriteUint16(i uint16) error {
	bs := make([]byte, 2)
	f.byteOrder.PutUint16(bs, i)
	return f.Write(bs)
}

func (f *FSIndexOutput) WriteUint32(i uint32) error {
	bs := make([]byte, 4)
	f.byteOrder.PutUint32(bs, i)
	return f.Write(bs)
}

func (f *FSIndexOutput) WriteUint64(i uint64) error {
	bs := make([]byte, 8)
	f.byteOrder.PutUint64(bs, i)
	return f.Write(bs)
}

func (f *FSIndexOutput) WriteVInt(i uint64) error {
	bs := proto.EncodeVarint(i)
	return f.Write(bs)
}

func (f *FSIndexOutput) WriteZInt32(i uint64) error {
	return f.WriteVInt(zigZagEncode32(i))
}

func (f *FSIndexOutput) WriteZInt64(i uint64) error {
	return f.WriteVInt(zigZagEncode64(i))
}

func (f *FSIndexOutput) WriteMapOfStrings(values map[string]string) error {
	if err := f.WriteVInt(uint64(len(values))); err != nil {
		return err
	}

	for k, v := range values {
		if err := f.WriteString(k); err != nil {
			return err
		}
		if err := f.WriteString(v); err != nil {
			return err
		}
	}
	return nil
}

func (f *FSIndexOutput) WriteSetOfStrings(values map[string]struct{}) error {
	if err := f.WriteVInt(uint64(len(values))); err != nil {
		return err
	}

	for k, _ := range values {
		if err := f.WriteString(k); err != nil {
			return err
		}
	}
	return nil
}

func (f *FSIndexOutput) CopyBytes(input DataInput, size int64) error {
	if size <= 0 {
		return fmt.Errorf("size is negative")
	}

	bs, err := input.ReadBytes(int(size))
	if err != nil {
		return err
	}
	return f.Write(bs)
}

func (f *FSIndexOutput) GetName() string {
	return f.name
}

func (f *FSIndexOutput) GetFilePointer() int64 {
	return f.bytesWritten
}

func (f *FSIndexOutput) GetChecksum() (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FSIndexOutput) AlignFilePointer(alignmentBytes int) (int64, error) {
	offset := f.GetFilePointer()
	alignedOffset, err := f.AlignOffset(offset, alignmentBytes)
	if err != nil {
		return 0, err
	}

	count := int(alignedOffset - offset)
	bs := make([]byte, count)
	for i := 0; i < count; i++ {
		bs[i] = 0
	}

	if err := f.Write(bs); err != nil {
		return 0, err
	}
	return alignedOffset, nil
}

func (f *FSIndexOutput) AlignOffset(offset int64, alignmentBytes int) (int64, error) {
	if offset < 0 {
		return 0, errors.New("offset must be positive")
	}

	if bitCount(alignmentBytes) != 1 || alignmentBytes < 0 {
		return 0, errors.New("alignment must be a power of 2")
	}

	n, err := addExact(offset-1, int64(alignmentBytes))
	if err != nil {
		return 0, err
	}

	return n & int64(-alignmentBytes), nil
}

func zigZagEncode32(x uint64) uint64 {
	return uint64((uint32(x) << 1) ^ uint32(int32(x)>>31))
}

func zigZagEncode64(x uint64) uint64 {
	return (x << 1) ^ uint64(int64(x)>>63)
}

func bitCount(num int) int {
	count := 0
	for num != 0 {
		if num&1 > 0 {
			count++
		}
		num = num >> 1
	}
	return count
}

func addExact(x, y int64) (int64, error) {
	r := x + y
	if ((x ^ r) & (y ^ r)) < 0 {
		return 0, errors.New("long overflow")
	}
	return r, nil
}
