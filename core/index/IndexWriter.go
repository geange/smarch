package index

import (
	"github.com/geange/smarch/core/store"
	"math"
)

const (
	// MaxDocs Hard limit on maximum number of documents that may be added to the index. If you try to add
	// more than this you'll hit IllegalArgumentException.
	//
	// We defensively subtract 128 to be well below the lowest
	// ArrayUtil.MAX_ARRAY_LENGTH on "typical" JVMs.  We don't just use
	// ArrayUtil.MAX_ARRAY_LENGTH here because this can vary across JVMs:
	MaxDocs = math.MaxInt32 - 128

	// MaxPosition Maximum value of the token position in an indexed field.
	MaxPosition = math.MaxInt32 - 128
)

type IndexWriter struct {
	// Use package-private instance var to enforce the limit so testing
	// can use less electricity:
	actualMaxDocs int

	// original user directory
	directoryOrig store.Directory
	// wrapped with additional checks
	directory store.Directory

	// increments every time a change is completed
	changeCount int64
	// last changeCount that was committed
	lastCommitChangeCount int64

	// list of segmentInfo we will fallback to if the commit fails
	rollbackSegments []*SegmentCommitInfo

	// set when a commit is pending (after prepareCommit() & before commit())
	pendingCommit *SegmentInfos
}

// SetMaxDocs Used only for testing.
func (i *IndexWriter) SetMaxDocs(maxDocs int) {
	if maxDocs > MaxDocs {
		maxDocs = MaxDocs
	}
	i.actualMaxDocs = maxDocs
}

func (i *IndexWriter) GetActualMaxDocs() int {
	return i.actualMaxDocs
}
