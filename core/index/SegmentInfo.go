package index

import (
	. "github.com/geange/smarch/core/codecs"
	. "github.com/geange/smarch/core/search"
	. "github.com/geange/smarch/core/store"
)

// SegmentInfo Information about a segment such as its name, directory, and files related to the segment.
type SegmentInfo struct {

	// TODO: remove these from this class, for now this is the representation
	// Used by some member fields to mean not present (e.g., norms, deletions).
	// default -1
	NO int

	// Used by some member fields to mean present (e.g., norms, deletions).
	// e.g. have norms; have deletes;
	// default -1
	YES int

	//Unique segment name in the directory.
	name string

	// number of docs in seg
	maxDoc int

	// Where this segment resides.
	dir Directory

	isCompoundFile bool

	// Id that uniquely identifies this segment.
	id []byte

	codec Codec

	diagnostics map[string]string

	attributes map[string]string

	indexSort Sort

	// Tracks the Lucene version this segment was created with, since 3.1. Null
	// indicates an older than 3.0 index, and it's used to detect a too old index.
	// The format expected is "x.y" - "2.x" for pre-3.0 indexes (or null), and
	// specific versions afterwards ("3.0.0", "3.1.0" etc.).
	// see o.a.l.util.Version.
	version string

	// Tracks the minimum version that contributed documents to a segment. For
	// flush segments, that is the version that wrote it. For merged segments,
	// this is the minimum minVersion of all the segments that have been merged
	// into this segment
	minVersion string

	setFiles map[string]struct{}
}

func NewSegmentInfo(dir Directory, version, minVersion, name string,
	maxDoc int, isCompoundFile bool, codec Codec, diagnostics map[string]string,
	id []byte, attributes map[string]string, indexSort Sort) *SegmentInfo {

	return &SegmentInfo{
		NO:             -1,
		YES:            -1,
		name:           name,
		maxDoc:         maxDoc,
		dir:            dir,
		isCompoundFile: isCompoundFile,
		id:             id,
		codec:          codec,
		diagnostics:    diagnostics,
		attributes:     attributes,
		indexSort:      indexSort,
		version:        version,
		minVersion:     minVersion,
		setFiles:       make(map[string]struct{}),
	}
}

func (s *SegmentInfo) Name() string {
	return s.name
}

func (s *SegmentInfo) Dir() Directory {
	return s.dir
}

// SetUseCompoundFile Mark whether this segment is stored as a compound file.
//isCompoundFile true if this is a compound file; else, false
func (s *SegmentInfo) SetUseCompoundFile(isCompoundFile bool) {
	s.isCompoundFile = isCompoundFile
}

// GetUseCompoundFile Returns true if this segment is stored as a compound file; else, false.
func (s *SegmentInfo) GetUseCompoundFile() bool {
	return s.isCompoundFile
}

// SetCodec Can only be called once.
func (s *SegmentInfo) SetCodec(codec Codec) {
	s.codec = codec
}

// GetCodec Return Codec that wrote this segment.
func (s *SegmentInfo) GetCodec() Codec {
	return s.codec
}

// MaxDoc Returns number of documents in this segment (deletions are not taken into account).
func (s *SegmentInfo) MaxDoc() int {
	return s.maxDoc
}

func (s *SegmentInfo) SetMaxDoc(maxDoc int) {
	s.maxDoc = maxDoc
}

// Files Return all files referenced by this SegmentInfo.
func (s *SegmentInfo) Files() map[string]struct{} {
	return s.setFiles
}

// GetVersion Returns the version of the code which wrote the segment.
func (s *SegmentInfo) GetVersion() string {
	return s.version
}

// GetMinVersion Return the minimum Lucene version that contributed documents to this segment,
// or null if it is unknown.
func (s *SegmentInfo) GetMinVersion() string {
	return s.minVersion
}

// GetId Return the id that uniquely identifies this segment.
func (s *SegmentInfo) GetId() []byte {
	bs := make([]byte, len(s.id))
	copy(bs, s.id)
	return bs
}

// SetFiles Sets the files written for this segment.
func (s *SegmentInfo) SetFiles(files map[string]struct{}) {
	s.setFiles = files
}

// AddFiles Add these files to the set of files written for this segment.
func (s *SegmentInfo) AddFiles(files map[string]struct{}) {
	for file, _ := range files {
		s.setFiles[s.namedForThisSegment(file)] = struct{}{}
	}
}

// AddFile Add this file to the set of files written for this segment.
func (s *SegmentInfo) AddFile(file string) {
	s.setFiles[s.namedForThisSegment(file)] = struct{}{}
}

func (s *SegmentInfo) checkFileNames(files map[string]struct{}) {

}

// strips any segment name from the file, naming it with this segment this is because "segment
// names" can change, e.g. by addIndexes(Dir)
func (s *SegmentInfo) namedForThisSegment(file string) string {
	return s.name + IndexFileNamesInstance.StripSegmentName(file)
}

// GetAttribute Get a codec attribute value, or null if it does not exist
func (s *SegmentInfo) GetAttribute(key string) string {
	return s.attributes[key]
}

// PutAttribute Puts a codec attribute value.
//
// This is a key-value mapping for the field that the codec can use to store additional
// metadata, and will be available to the codec when reading the segment via
// #getAttribute(String)
//
// If a value already exists for the field, it will be replaced with the new value. This method
// make a copy on write for every attribute change.
func (s *SegmentInfo) PutAttribute(key, value string) string {
	oldValue := s.attributes[key]
	s.attributes[key] = value
	return oldValue
}

// GetAttributes Returns the internal codec attributes map. internal codec attributes map.
func (s *SegmentInfo) GetAttributes() map[string]string {
	return s.attributes
}

// GetIndexSort Return the sort order of this segment, or null if the index has no sort.
func (s *SegmentInfo) GetIndexSort() Sort {
	return s.indexSort
}
