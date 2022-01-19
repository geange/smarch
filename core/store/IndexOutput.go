package store

import "io"

type IndexOutput interface {
	io.Closer

	DataOutput

	// GetName Returns the name used to create this IndexOutput. This is especially useful when using
	// Directory#createTempOutput.
	GetName() string

	// GetFilePointer Returns the current position in this file, where the next write will occur.
	GetFilePointer() int64

	// GetChecksum Returns the current checksum of bytes written so far
	GetChecksum() (int64, error)

	// AlignFilePointer Aligns the current file pointer to multiples of alignmentBytes bytes to improve reads
	// with mmap. This will write between 0 and (alignmentBytes-1) zero bytes using #writeByte(byte).
	AlignFilePointer(alignmentBytes int) (int64, error)

	// AlignOffset Aligns the given offset to multiples of alignmentBytes bytes by rounding up.
	// The alignment must be a power of 2.
	AlignOffset(offset int64, alignmentBytes int) (int64, error)
}
