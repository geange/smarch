package index

import (
	. "github.com/geange/smarch/core/analysis"
	. "github.com/geange/smarch/core/codecs"
	. "github.com/geange/smarch/core/util"
)

// LiveIndexWriterConfig Holds all the configuration used by IndexWriter with few setters for settings that can be
// changed on an IndexWriter instance "live".
//
// since 4.0
type LiveIndexWriterConfig struct {
	analyzer Analyzer

	maxBufferedDocs     int
	mergedSegmentWarmer IndexReaderWarmer

	// modified by IndexWriterConfig

	// IndexDeletionPolicy controlling when commit points are deleted.
	delPolicy IndexDeletionPolicy

	// IndexCommit that IndexWriter is opened on.
	commit IndexCommit

	// OpenMode that IndexWriter is opened with.
	openMode OpenModeType

	//Compatibility version to use for this index.
	createdVersionMajor int

	// Similarity to use when encoding norms.
	similarity Similarity

	//MergeScheduler to use for running merges.
	mergeScheduler MergeScheduler

	//Codec used to write new segments.
	codec Codec

	//InfoStream for debugging messages.
	infoStream InfoStream

	//MergePolicy for selecting merges.
	mergePolicy MergePolicy

	//True if readers should be pooled.
	readerPooling bool

	//FlushPolicy to control when segments are flushed.
	flushPolicy FlushPolicy
}

// GetAnalyzer Returns the default analyzer to use for indexing documents.
func (l *LiveIndexWriterConfig) GetAnalyzer() Analyzer {
	return l.analyzer
}
