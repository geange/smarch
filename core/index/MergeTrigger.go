package index

type MergeTriggerType int

const (
	// MTypeSegmentFlush Merge was triggered by a segment flush.
	MTypeSegmentFlush = MergeTriggerType(iota)

	// MTypeFullFlush Merge was triggered by a full flush. Full flushes can be caused by a commit,
	// NRT reader reopen or a close call on the index writer.
	MTypeFullFlush

	// MTypeExplicit Merge has been triggered explicitly by the user.
	MTypeExplicit

	// MergeFinished Merge was triggered by a successfully finished merge.
	MergeFinished

	// MTypeClosing Merge was triggered by a closing IndexWriter.
	MTypeClosing

	// MTypeCommit Merge was triggered on commit.
	MTypeCommit

	// MTypeGetReader Merge was triggered on opening NRT readers.
	MTypeGetReader
)
