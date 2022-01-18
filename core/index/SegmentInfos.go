package index

// SegmentInfos A collection of segmentInfo objects with methods for operating on those segments in relation to
// the file system.
//
// The active segments in the index are stored in the segment info file, <code>segments_N</code>.
// There may be one or more <code>segments_N</code> files in the index; however, the one with the
// largest generation is the active one (when older segments_N files are present it's because they
// temporarily cannot be deleted, or a custom {@link IndexDeletionPolicy} is in use). This file
// lists each segment by name and has details about the codec and generation of deletes.
type SegmentInfos struct {
	//Used to name new segments.
	counter int64

	//Counts how often the index has been changed.
	version int64

	// generation of the "segments_N" for the next commit
	generation int64

	// generation of the "segments_N" file we last successfully read
	// or wrote; this is normally the same as generation except if
	// there was an IOException that had interrupted a commit
	lastGeneration int64

	//Opaque map[string]string; that user can specify during IndexWriter.commit
	userData map[string]string

	segments []*SegmentCommitInfo
}

var (
	VersionCurrent = VERSION86
)

const (
	// VERSION70 The version that added information about the Lucene version at the time when the index has been created.
	VERSION70 = 7

	// VERSION72 The version that updated segment name counter to be long instead of int.
	VERSION72 = 8

	// VERSION74 The version that recorded softDelCount
	VERSION74 = 9

	// VERSION86 The version that recorded SegmentCommitInfo IDs
	VERSION86 = 10

	// OldSegmentsGen Name of the generation reference file name
	OldSegmentsGen = "segments.gen"
)
