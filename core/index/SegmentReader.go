package index

// SegmentReader IndexReader implementation over a single segment.
// Instances pointing to the same segment (but with different deletes, etc) may share the same core data.
type SegmentReader struct {
}
