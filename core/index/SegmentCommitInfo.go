package index

import (
	"fmt"
	"github.com/google/uuid"
)

// SegmentCommitInfo Embeds a [read-only] SegmentInfo and adds per-commit fields.
// 嵌入 [只读] SegmentInfo 并添加每个提交字段。
type SegmentCommitInfo struct {

	//The SegmentInfo that we wrap.
	info *SegmentInfo

	//Id that uniquely identifies this segment commit.
	id []byte

	// How many deleted docs in the segment:
	delCount int

	// How many soft-deleted docs in the segment that are not also hard-deleted:
	softDelCount int

	// Generation number of the live docs file (-1 if there
	// are no deletes yet):
	delGen int64

	// Normally 1+delGen, unless an exception was hit on last
	// attempt to write:
	nextWriteDelGen int64

	// Generation number of the FieldInfos (-1 if there are no updates)
	fieldInfosGen int64

	// Normally 1+fieldInfosGen, unless an exception was hit on last attempt to write
	nextWriteFieldInfosGen int64

	// Generation number of the DocValues (-1 if there are no updates)
	docValuesGen int64

	// Normally 1+dvGen, unless an exception was hit on last attempt to write
	nextWriteDocValuesGen int64

	// Track the per-field DocValues update files
	dvUpdatesFiles map[int]map[string]struct{}

	// TODO should we add .files() to FieldInfosFormat, like we have on
	// LiveDocsFormat?
	// track the fieldInfos update files
	fieldInfosFiles map[string]struct{}

	// default -1
	sizeInBytes int64

	// NOTE: only used in-RAM by IW to track buffered deletes;
	// this is never written to/read from the Directory
	// default -1
	bufferedDeletesGen int64
}

// GetDocValuesUpdatesFiles Returns the per-field DocValues updates files.
func (s *SegmentCommitInfo) GetDocValuesUpdatesFiles() map[int]map[string]struct{} {
	return s.dvUpdatesFiles
}

// SetDocValuesUpdatesFiles Sets the DocValues updates file names, per field number. Does not deep clone the map.
func (s *SegmentCommitInfo) SetDocValuesUpdatesFiles(dvUpdatesFiles map[int]map[string]struct{}) {
	s.dvUpdatesFiles = map[int]map[string]struct{}{}

	for k, set := range dvUpdatesFiles {
		data := make(map[string]struct{})
		for file, _ := range set {
			data[file] = struct{}{}
		}
		s.dvUpdatesFiles[k] = data
	}
}

// GetFieldInfosFiles Returns the FieldInfos file names.
func (s *SegmentCommitInfo) GetFieldInfosFiles() map[string]struct{} {
	return s.fieldInfosFiles
}

// SetFieldInfosFiles Sets the FieldInfos file names.
func (s *SegmentCommitInfo) SetFieldInfosFiles(fieldInfosFiles map[string]struct{}) {
	s.fieldInfosFiles = make(map[string]struct{}, len(fieldInfosFiles))
	for file, _ := range fieldInfosFiles {
		s.fieldInfosFiles[file] = struct{}{}
	}
}

// AdvanceDelGen Called when we succeed in writing deletes
func (s *SegmentCommitInfo) AdvanceDelGen() {
	s.delGen = s.nextWriteDelGen
	s.nextWriteDelGen++
	s.generationAdvanced()
}

// AdvanceNextWriteDelGen Called if there was an exception while writing deletes,
// so that we don't try to write to the same file more than once.
func (s *SegmentCommitInfo) AdvanceNextWriteDelGen() {
	s.nextWriteDelGen++
}

// GetNextWriteDelGen Gets the nextWriteDelGen.
func (s *SegmentCommitInfo) GetNextWriteDelGen() int64 {
	return s.nextWriteDelGen
}

// SetNextWriteDelGen Sets the nextWriteDelGen.
func (s *SegmentCommitInfo) SetNextWriteDelGen(v int64) {
	s.nextWriteDelGen = v
}

// AdvanceFieldInfosGen Called when we succeed in writing a new FieldInfos generation.
func (s *SegmentCommitInfo) AdvanceFieldInfosGen() {
	s.fieldInfosGen = s.nextWriteFieldInfosGen
	s.nextWriteDocValuesGen++
	s.generationAdvanced()
}

// AdvanceNextWriteFieldInfosGen Called if there was an exception while writing a new generation of FieldInfos,
// so that we don't try to write to the same file more than once.
func (s *SegmentCommitInfo) AdvanceNextWriteFieldInfosGen() {
	s.nextWriteFieldInfosGen++
}

// GetNextWriteFieldInfosGen Gets the nextWriteFieldInfosGen.
func (s *SegmentCommitInfo) GetNextWriteFieldInfosGen() int64 {
	return s.nextWriteFieldInfosGen
}

// SetNextWriteFieldInfosGen Sets the nextWriteFieldInfosGen.
func (s *SegmentCommitInfo) SetNextWriteFieldInfosGen(v int64) {
	s.nextWriteFieldInfosGen = v
}

// AdvanceDocValuesGen Called when we succeed in writing a new DocValues generation.
func (s *SegmentCommitInfo) AdvanceDocValuesGen() {
	s.docValuesGen = s.nextWriteDocValuesGen
	s.nextWriteDocValuesGen++
	s.generationAdvanced()
}

// AdvanceNextWriteDocValuesGen Called if there was an exception while writing a new generation of DocValues,
// so that we don't try to write to the same file more than once.
func (s *SegmentCommitInfo) AdvanceNextWriteDocValuesGen() {
	s.nextWriteDocValuesGen++
}

// GetNextWriteDocValuesGen Gets the nextWriteDocValuesGen.
func (s *SegmentCommitInfo) GetNextWriteDocValuesGen() int64 {
	return s.nextWriteDocValuesGen
}

// SetNextWriteDocValuesGen Sets the nextWriteDocValuesGen.
func (s *SegmentCommitInfo) SetNextWriteDocValuesGen(v int64) {
	s.nextWriteDocValuesGen = v
}

// SizeInBytes Returns total size in bytes of all files for this segment.
func (s *SegmentCommitInfo) SizeInBytes() (int64, error) {
	if s.sizeInBytes == -1 {
		s.sizeInBytes = 0
		for _, file := range s.Files() {
			size, err := s.info.Dir().FileLength(file)
			if err != nil {
				return 0, err
			}
			s.sizeInBytes += size
		}
	}
	return s.sizeInBytes, nil
}

func (s *SegmentCommitInfo) Files() []string {
	set := s.info.Files()
	files := make([]string, 0, len(set))
	for file, _ := range set {
		files = append(files, file)
	}
	return files
}

func (s *SegmentCommitInfo) getBufferedDeletesGen() int64 {
	return s.bufferedDeletesGen
}

func (s *SegmentCommitInfo) SetBufferedDeletesGen(v int64) {
	if s.bufferedDeletesGen == -1 {
		s.bufferedDeletesGen = v
		s.generationAdvanced()
	}
}

// HasDeletions Returns true if there are any deletions for the segment at this commit.
func (s *SegmentCommitInfo) HasDeletions() bool {
	return s.delGen != -1
}

// HasFieldUpdates Returns true if there are any field updates for the segment in this commit.
func (s *SegmentCommitInfo) HasFieldUpdates() bool {
	return s.fieldInfosGen != -1
}

// GetNextFieldInfosGen Returns the next available generation number of the FieldInfos files.
func (s *SegmentCommitInfo) GetNextFieldInfosGen() int64 {
	return s.nextWriteFieldInfosGen
}

// GetFieldInfosGen Returns the generation number of the field infos file or -1 if there are no field updates yet.
func (s *SegmentCommitInfo) GetFieldInfosGen() int64 {
	return s.fieldInfosGen
}

// GetNextDocValuesGen Returns the next available generation number of the DocValues files.
func (s *SegmentCommitInfo) GetNextDocValuesGen() int64 {
	return s.nextWriteDocValuesGen
}

// GetDocValuesGen Returns the generation number of the DocValues file or -1 if there are no doc-values updates yet.
func (s *SegmentCommitInfo) GetDocValuesGen() int64 {
	return s.docValuesGen
}

// GetNextDelGen Returns the next available generation number of the live docs file.
func (s *SegmentCommitInfo) GetNextDelGen() int64 {
	return s.nextWriteDelGen
}

// GetDelGen Returns generation number of the live docs file or -1 if there are no deletes yet.
func (s *SegmentCommitInfo) GetDelGen() int64 {
	return s.delGen
}

// GetDelCount Returns the number of deleted docs in the segment.
func (s *SegmentCommitInfo) GetDelCount() int {
	return s.delCount
}

// GetSoftDelCount Returns the number of only soft-deleted docs.
func (s *SegmentCommitInfo) GetSoftDelCount() int {
	return s.delCount
}

func (s *SegmentCommitInfo) setDelCount(delCount int) error {
	if delCount < 0 || delCount > s.info.MaxDoc() {
		return fmt.Errorf("invalid delCount=%d  (maxDoc=%d)", delCount, s.info.MaxDoc())
	}

	if s.softDelCount+delCount <= s.info.MaxDoc() {
		s.delCount = delCount
		return nil
	}

	return fmt.Errorf("maxDoc=%d,delCount=%d,softDelCount=%d",
		s.info.MaxDoc(), delCount, s.softDelCount)
}

func (s *SegmentCommitInfo) setSoftDelCount(softDelCount int) error {
	if softDelCount < 0 || softDelCount > s.info.MaxDoc() {
		return fmt.Errorf("invalid softDelCount=%d (maxDoc=%d)", softDelCount, s.info.MaxDoc())
	}

	if softDelCount+s.delCount <= s.info.MaxDoc() {
		s.softDelCount = softDelCount
		return nil
	}
	return fmt.Errorf("maxDoc=%d,delCount=%d,softDelCount=%d",
		s.info.MaxDoc(), s.delCount, softDelCount)
}

func (s *SegmentCommitInfo) GetDelCountV1(includeSoftDeletes bool) int {
	if includeSoftDeletes {
		return s.delCount + s.softDelCount
	}
	return s.delCount
}

func (s *SegmentCommitInfo) generationAdvanced() {
	s.sizeInBytes = -1
	s.id, _ = uuid.New().MarshalText()
}

// GetId Returns and Id that uniquely identifies this segment commit or null if there is no
// ID assigned. This ID changes each time the the segment changes due to a delete, doc-value or
// field update.
func (s *SegmentCommitInfo) GetId() []byte {
	if s.id == nil {
		return nil
	}
	bs := make([]byte, len(s.id))
	copy(bs, s.id)
	return bs
}

func NewSegmentCommitInfo(info *SegmentInfo, delCount, softDelCount int,
	delGen, fieldInfosGen, docValuesGen int64, id []byte) *SegmentCommitInfo {

	nextWriteDelGen := delGen + 1
	if delGen == -1 {
		nextWriteDelGen = -1
	}

	nextWriteFieldInfosGen := fieldInfosGen + 1
	if fieldInfosGen == -1 {
		nextWriteFieldInfosGen = -1
	}

	nextWriteDocValuesGen := docValuesGen + 1
	if docValuesGen == -1 {
		nextWriteDocValuesGen = -1
	}

	return &SegmentCommitInfo{
		info:                   info,
		id:                     id,
		delCount:               delCount,
		softDelCount:           softDelCount,
		delGen:                 delGen,
		nextWriteDelGen:        nextWriteDelGen,
		fieldInfosGen:          fieldInfosGen,
		nextWriteFieldInfosGen: nextWriteFieldInfosGen,
		docValuesGen:           docValuesGen,
		nextWriteDocValuesGen:  nextWriteDocValuesGen,
		dvUpdatesFiles:         make(map[int]map[string]struct{}),
		fieldInfosFiles:        make(map[string]struct{}),
		sizeInBytes:            -1,
		bufferedDeletesGen:     -1,
	}
}

func (s *SegmentCommitInfo) Clone() *SegmentCommitInfo {
	other := NewSegmentCommitInfo(s.info, s.delCount, s.softDelCount,
		s.delGen, s.fieldInfosGen, s.docValuesGen, s.GetId())

	// Not clear that we need to carry over nextWriteDelGen
	// (i.e. do we ever clone after a failed write and
	// before the next successful write?), but just do it to
	// be safe:
	other.nextWriteDelGen = s.nextWriteDelGen
	other.nextWriteFieldInfosGen = s.nextWriteFieldInfosGen
	other.nextWriteDocValuesGen = s.nextWriteDocValuesGen

	// deep clone
	for key, set := range s.dvUpdatesFiles {
		data := make(map[string]struct{}, len(set))
		for v, _ := range set {
			data[v] = struct{}{}
		}
		other.dvUpdatesFiles[key] = data
	}

	for file, _ := range s.fieldInfosFiles {
		other.fieldInfosFiles[file] = struct{}{}
	}
	return other
}
