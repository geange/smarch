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
}
