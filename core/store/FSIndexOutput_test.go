package store

import (
	"encoding/binary"
	"fmt"
	"os"
	"testing"
)

func TestFSIndexOutput_WriteByte(t *testing.T) {
	file, err := os.CreateTemp(`D:\code\golang\smarch`, "demo-*.txt")
	if err != nil {
		t.Error(err)
	}
	output := FSIndexOutput{
		name:           "xxxxxxxx",
		file:           file,
		byteOrder:      binary.BigEndian,
		bytesWritten:   0,
		flushedOnClose: false,
	}
	if err := output.WriteString("Hello"); err != nil {
		t.Error(err)
	}

	stat, _ := file.Stat()
	fmt.Println(stat.Size())

	if err := output.Close(); err != nil {
		t.Error(err)
	}
}
