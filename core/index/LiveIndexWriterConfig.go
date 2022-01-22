package index

import "github.com/geange/smarch/core/analysis"

// LiveIndexWriterConfig Holds all the configuration used by IndexWriter with few setters for settings that can be
// changed on an IndexWriter instance "live".
//
// since 4.0
type LiveIndexWriterConfig struct {
	analyzer analysis.Analyzer

	maxBufferedDocs     int
	mergedSegmentWarmer IndexReaderWarmer

	delPolicy IndexDeletionPolicy
}
