package index

import (
	. "github.com/geange/smarch/core/analysis"
	. "github.com/geange/smarch/core/document"
	. "github.com/geange/smarch/core/search"
	. "github.com/geange/smarch/core/store"
	. "github.com/geange/smarch/core/util"
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
	directoryOrig Directory
	// wrapped with additional checks
	directory Directory

	// increments every time a change is completed
	changeCount int64
	// last changeCount that was committed
	lastCommitChangeCount int64

	// list of segmentInfo we will fallback to if the commit fails
	rollbackSegments []*SegmentCommitInfo

	// set when a commit is pending (after prepareCommit() & before commit())
	pendingCommit *SegmentInfos

	pendingSeqNo             int64
	pendingCommitChangeCount int64
	filesToCommit            map[string]struct{}

	segmentInfos         SegmentInfos
	globalFieldNumberMap *FieldNumbers

	docWriter *DocumentsWriter

	mergeSource MergeSource

	//writeDocValuesLock ReentrantLock
	//deleter IndexFileDeleter

	//segmentsToMerge map[SegmentCommitInfo]bool
	mergeMaxNumSegments int

	writeLock Lock

	closed  bool
	closing bool

	maybeMerge bool

	commitUserData map[string]string

	// 	Holds all SegmentInfo instances currently involved in
	// 	merges
	//  private final HashSet<SegmentCommitInfo> mergingSegments = new HashSet<>();
	//  private final MergeScheduler mergeScheduler;
	//  private final Set<SegmentMerger> runningAddIndexesMerges = new HashSet<>();
	//  private final Deque<MergePolicy.OneMerge> pendingMerges = new ArrayDeque<>();
	//  private final Set<MergePolicy.OneMerge> runningMerges = new HashSet<>();
	//  private final List<MergePolicy.OneMerge> mergeExceptions = new ArrayList<>();
	//  private long mergeGen;
	//  private Merges merges = new Merges();
	//  private boolean didMessageState;
	//  private final AtomicInteger flushCount = new AtomicInteger();
	//  private final AtomicInteger flushDeletesCount = new AtomicInteger();
	//  private final ReaderPool readerPool;
	//  private final BufferedUpdatesStream bufferedUpdatesStream;
	//
	//  private final IndexWriterEventListener eventListener;

	// Counts how many merges have completed; this is used by
	// #forceApply(FrozenBufferedUpdates) to handle concurrently apply deletes/updates with merges
	// completing.
	mergeFinishedGen int64

	// The instance that was passed to the constructor. It is saved only in order
	// to allow users to query an IndexWriter settings.
	config *LiveIndexWriterConfig

	// System.nanoTime() when commit started; used to write an infoStream message about how long
	// commit took.
	startCommitTime int64

	// How many documents are in the index, or are in the process of being added (reserved). E.g.,
	// operations like addIndexes will first reserve the right to add N docs, before they actually
	// change the index, much like how hotels place an "authorization hold" on your credit card to
	// make sure they can later charge you when you check out.
	pendingNumDocs int64

	softDeletesEnabled bool

	infoStream InfoStream
}

func (i *IndexWriter) GetActualMaxDocs() int {
	return i.actualMaxDocs
}

// SetMaxDocs Used only for testing.
func (i *IndexWriter) SetMaxDocs(maxDocs int) {
	if maxDocs > MaxDocs {
		maxDocs = MaxDocs
	}
	i.actualMaxDocs = maxDocs
}

// GetReader Expert: returns a readonly reader, covering all committed as well as un-committed changes to
// the index. This provides "near real-time" searching, in that changes made during an IndexWriter
// session can be quickly made available for searching without closing the writer nor calling
// #commit.
//
// Note that this is functionally equivalent to calling #flush and then opening a new reader.
// But the turnaround time of this method should be faster since it avoids the potentially costly
// #commit.
//
// You must close the IndexReader returned by this method once you are done using it.
//
// It's near real-time because there is no hard guarantee on how quickly you can get a
// new reader after making changes with IndexWriter. You'll have to experiment in your situation
// to determine if it's fast enough. As this is a new and experimental feature, please report back
// on your findings so we can learn, improve and iterate.
//
// The resulting reader supports DirectoryReader#openIfChanged, but that call will
// simply forward back to this method (though this may change in the future).
//
// The very first time this method is called, this writer instance will make every effort to
// pool the readers that it opens for doing merges, applying deletes, etc. This means additional
// resources (RAM, file descriptors, CPU time) will be consumed.
//
// For lower latency on reopening a reader, you should call
// IndexWriterConfig#setMergedSegmentWarmer to pre-warm a newly merged segment before it's
// committed to the index. This is important for minimizing index-to-search delay after a large
// merge.
//
// If an addIndexes* call is running in another thread, then this reader will only search those
// segments from the foreign index that have been successfully copied over, so far.
//
// NOTE: Once the writer is closed, any outstanding readers may continue to be used.
// However, if you attempt to reopen any of those readers, you'll hit an
// AlreadyClosedException.
func (i *IndexWriter) GetReader(applyAllDeletes, writeAllDeletes bool) (DirectoryReader, error) {
	panic("")
}

// GetFlushingBytes Returns the number of bytes currently being flushed
func (i *IndexWriter) GetFlushingBytes() int64 {
	panic("")
}

func (i *IndexWriter) writeSomeDocValuesUpdates() error {
	panic("")
}

// NumDeletedDocs Obtain the number of deleted docs for a pooled reader.
// If the reader isn't being pooled, the segmentInfo's delCount is returned.
func (i *IndexWriter) NumDeletedDocs(info *SegmentCommitInfo) int {
	panic("")
}

// Used internally to throw an AlreadyClosedException} if this IndexWriter has been closed
// or is in the process of closing.
//
// failIfClosing if true, also fail when  IndexWriter} is in the process of closing
// (closing=true) but not yet done closing (closed=false)
func (i IndexWriter) ensureOpenV1(failIfClosing bool) error {
	panic("")
}

// Used internally to throw an AlreadyClosedException} if this IndexWriter has been closed
// (closed=true}) or is in the process of closing (closing=true}).
func (i *IndexWriter) ensureOpen() error {
	panic("")
}

//Confirms that the incoming index sort (if any) matches the existing index sort (if any).
func (i *IndexWriter) validateIndexSort() error {
	panic("")
}

// IsCongruentSort Returns true if indexSort is a prefix of otherSort.
func (i *IndexWriter) IsCongruentSort() bool {
	panic("")
}

// ReadFieldInfos reads latest field infos for the commit
// this is used on IW init and addIndexes(Dir) to create/update the global field map.
// TODO: fix tests abusing this method!
func (i *IndexWriter) ReadFieldInfos(si *SegmentCommitInfo) (*FieldInfos, error) {
	panic("")
}

// GetFieldNumberMap Loads or returns the already loaded the global field number map for this SegmentInfos.
// If this SegmentInfos has no global field number map the returned instance is empty
func (i *IndexWriter) GetFieldNumberMap() (*FieldNumbers, error) {
	panic("")
}

// Returns a LiveIndexWriterConfig, which can be used to query the IndexWriter current
// settings, as well as modify "live" ones.
func (i *IndexWriter) getConfig() *LiveIndexWriterConfig {
	return i.config
}

func (i *IndexWriter) messageState() error {
	panic("")
}

// Gracefully closes (commits, waits for merges), but calls rollback if there's an exc so the
// IndexWriter is always closed. This is called from #close when
// IndexWriterConfig#commitOnClose is true.
func (i *IndexWriter) shutdown() error {
	panic("")
}

// Close Closes all open resources and releases the write lock.
//
// If IndexWriterConfig#commitOnClose is true, this will attempt to
// gracefully shut down by writing any changes, waiting for any running merges, committing, and
// closing. In this case, note that:
//
// <ul>
//   <li>If you called prepareCommit but failed to call commit, this method will throw
//       IllegalStateException and the IndexWriter will not be closed.
//   <li>If this method throws any other exception, the IndexWriter will be closed, but
//       changes may have been lost.
// </ul>
//
// Note that this may be a costly operation, so, try to re-use a single writer instead of
// closing and opening a new one. See #commit() for caveats about write caching done by
// some IO devices.
//
// NOTE: You must ensure no other threads are still making changes at the same time that
// this method is invoked.
func (i *IndexWriter) Close() error {
	panic("")
}

// Returns true if this thread should attempt to close, or
// false if IndexWriter is now closed; else,
// waits until another thread finishes closing
func (i *IndexWriter) shouldClose(waitForClose bool) bool {
	panic("")
}

// GetDirectory Returns the Directory used by this index.
func (i *IndexWriter) GetDirectory() Directory {
	return i.directoryOrig
}

// GetInfoStream Returns InfoStream used for debugging.
func (i *IndexWriter) GetInfoStream() InfoStream {
	return i.infoStream
}

// GetAnalyzer Returns the analyzer used by this index.
func (i *IndexWriter) GetAnalyzer() Analyzer {
	return i.config.analyzer
}

// AdvanceSegmentInfosVersion If SegmentInfos#getVersion is below newVersion then update it to this value.
func (i *IndexWriter) AdvanceSegmentInfosVersion(newVersion int64) {

}

// HasDeletions Returns true if this index has deletions (including buffered deletions). Note that this will
// return true if there are buffered Term/Query deletions, even if it turns out those buffered
// deletions don't match any documents.
func (i *IndexWriter) HasDeletions() bool {
	panic("")
}

// AddDocument Adds a document to this index.
//
// Note that if an Exception is hit (for example disk full) then the index will be consistent,
// but this document may not have been added. Furthermore, it's possible the index will have one
// segment in non-compound format even when using compound files (when a merge has partially
// succeeded).
//
// This method periodically flushes pending documents to the Directory (see above,
// and also periodically triggers segment merges in the index according
// to the MergePolicy in use.
//
// Merges temporarily consume space in the directory. The amount of space required is up to 1X
// the size of all segments being merged, when no readers/searchers are open against the index,
// and up to 2X the size of all segments being merged when readers/searchers are open against the
// index (see #forceMerge(int) for details). The sequence of primitive merge operations
// performed is governed by the merge policy.
//
// Note that each term in the document can be no longer than #MAX_TERM_LENGTH in bytes,
// otherwise an IllegalArgumentException will be thrown.
//
// Note that it's possible to create an invalid Unicode string in java if a UTF16 surrogate
// pair is malformed. In this case, the invalid characters are silently replaced with the Unicode
// replacement character U+FFFD.
func (i *IndexWriter) AddDocument(doc *Document) (int64, error) {
	panic("")
}

// AddDocuments Atomically adds a block of documents with sequentially assigned document IDs,
// such that an external reader will see all or none of the documents.
//
// <b>WARNING</b>: the index does not currently record which documents were added as a block.
// Today this is fine, because merging will preserve a block. The order of documents within a
// segment will be preserved, even when child documents within a block are deleted. Most search
// features (like result grouping and block joining) require you to mark documents; when these
// documents are deleted these search features will not work as expected. Obviously adding
// documents to an existing block will require you the reindex the entire block.
//
// However it's possible that in the future Lucene may merge more aggressively re-order
// documents (for example, perhaps to obtain better index compression), in which case you may need
// to fully re-index your documents at that time.
//
// See #addDocument(Iterable)} for details on index and IndexWriter state after an
// Exception, and flushing/merging temporary free space requirements.
//
// <b>NOTE</b>: tools that do offline splitting of an index (for example, IndexSplitter in
// contrib) or re-sorting of documents (for example, IndexSorter in contrib) are not aware of
// these atomically added documents and will likely break them up. Use such tools at your own
// risk!
func (i *IndexWriter) AddDocuments(doc []*Document) (int64, error) {
	panic("")
}

// UpdateDocuments Atomically deletes documents matching the provided delTerm and adds a block of documents with
// sequentially assigned document IDs, such that an external reader will see all or none of the
// documents.
func (i *IndexWriter) UpdateDocuments(delTerm *Term, docs []*Document) int64 {
	panic("")
}

func (i *IndexWriter) updateDocuments(delNode interface{}, docs []*Document) int64 {
	panic("")
}

// SoftUpdateDocuments Expert: Atomically updates documents matching the provided term with the given doc-values
// fields and adds a block of documents with sequentially assigned document IDs, such that an
// external reader will see all or none of the documents.
//
// One use of this API is to retain older versions of documents instead of replacing them. The
// existing documents can be updated to reflect they are no longer current while atomically adding
// new documents at the same time.
//
// In contrast to #updateDocuments(Term, Iterable)} this method will not delete
// documents in the index matching the given term but instead update them with the given
// doc-values fields which can be used as a soft-delete mechanism.
//
// See #addDocuments(Iterable)} and #updateDocuments(Term, Iterable)}.
func (i *IndexWriter) SoftUpdateDocuments(term *Term, docs []*Document, softDeletes ...Field) (int64, error) {
	panic("")
}

// TryDeleteDocument Expert: attempts to delete by document ID, as long as the provided reader is a near-real-time
// reader (from DirectoryReader#open(IndexWriter)}). If the provided reader is an NRT
// reader obtained from this writer, and its segment has not been merged away, then the delete
// succeeds and this method returns a valid (&gt; 0) sequence number; else, it returns -1 and the
// caller must then separately delete by Term or Query.
//
// <b>NOTE</b>: this method can only delete documents visible to the currently open NRT reader.
// If you need to delete documents indexed after opening the NRT reader you must use
// #deleteDocuments(Term...)}).
func (i *IndexWriter) TryDeleteDocument(readerIn IndexReader, docID int) (int64, error) {
	panic("")
}

// TryUpdateDocValue Expert: attempts to update doc values by document ID, as long as the provided reader is a
// near-real-time reader (from {@link DirectoryReader#open(IndexWriter)}). If the provided reader
// is an NRT reader obtained from this writer, and its segment has not been merged away, then the
// update succeeds and this method returns a valid (&gt; 0) sequence number; else, it returns -1
// and the caller must then either retry the update and resolve the document again. If a doc
// values fields data is <code>null</code> the existing value is removed from all documents
// matching the term. This can be used to un-delete a soft-deleted document since this method will
// apply the field update even if the document is marked as deleted.
//
// NOTE: this method can only updates documents visible to the currently open NRT
// reader. If you need to update documents indexed after opening the NRT reader you must use
// {@link #updateDocValues(Term, Field...)}.
func (i *IndexWriter) TryUpdateDocValue(readerIn IndexReader, docID int, fields ...Field) (int64, error) {
	panic("")
}

func (i *IndexWriter) tryModifyDocument(readerIn IndexReader, docID int, toApply DocModifier) error {
	panic("")
}

//Drops a segment that has 100% deleted documents.
func (i *IndexWriter) dropDeletedSegment(info *SegmentCommitInfo) error {
	panic("")
}

// DeleteDocuments Deletes the document(s) containing any of the terms. All given deletes are applied and flushed
// atomically at the same time.
func (i *IndexWriter) DeleteDocuments(terms ...Terms) (int64, error) {
	panic("")
}

// DeleteDocumentsByQuery Deletes the document(s) matching any of the provided queries.
// All given deletes are applied and flushed atomically at the same time.
func (i *IndexWriter) DeleteDocumentsByQuery(queries ...Query) (int64, error) {
	panic("")
}

// UpdateDocument Updates a document by first deleting the document(s) containing term and then
// adding the new document. The delete and then add are atomic as seen by a reader on the same
// index (flush may happen only after the add).
func (i *IndexWriter) UpdateDocument(term *Term, doc *Document) (int64, error) {
	panic("")
}

// SoftUpdateDocument Expert: Updates a document by first updating the document(s) containing <code>term</code> with
// the given doc-values fields and then adding the new document. The doc-values update and then
// add are atomic as seen by a reader on the same index (flush may happen only after the add).
//
// One use of this API is to retain older versions of documents instead of replacing them. The
// existing documents can be updated to reflect they are no longer current while atomically adding
// new documents at the same time.
//
// In contrast to {@link #updateDocument(Term, Iterable)} this method will not delete documents
// in the index matching the given term but instead update them with the given doc-values fields
// which can be used as a soft-delete mechanism.
//
// See {@link #addDocuments(Iterable)} and {@link #updateDocuments(Term, Iterable)}.
func (i *IndexWriter) SoftUpdateDocument(term *Term, doc *Document, softDeletes ...Field) (int64, error) {
	panic("")
}

// UpdateNumericDocValue Updates a document's NumericDocValues for field to the given value.
// You can only update fields that already exist in the index, not add new fields through
// this method. You can only update fields that were indexed with doc values only.
func (i *IndexWriter) UpdateNumericDocValue(term *Term, field string, value int64) (int64, error) {
	panic("")
}

// UpdateBinaryDocValue Updates a document's BinaryDocValues for field to the given value.
// You can only update fields that already exist in the index, not add new fields through
// this method. You can only update fields that were indexed only with doc values.
//
// NOTE: this method currently replaces the existing value of all affected documents
// with the new value.
func (i *IndexWriter) UpdateBinaryDocValue(term *Term, field string, value []byte) (int64, error) {
	panic("")
}

// UpdateDocValues Updates documents' DocValues fields to the given values. Each field update is applied to the
// set of documents that are associated with the {@link Term} to the same value. All updates are
// atomically applied and flushed together. If a doc values fields data is <code>null</code> the
// existing value is removed from all documents matching the term.
func (i *IndexWriter) UpdateDocValues(term *Term, updates ...Field) (int64, error) {
	panic("")
}

func (i *IndexWriter) buildDocValuesUpdate(term *Term, updates []Field) []DocValuesUpdate {
	panic("")
}

// GetFieldNames Return an unmodifiable set of all field names as visible from this IndexWriter,
// across all segments of the index.
func (i *IndexWriter) GetFieldNames() map[string]struct{} {
	panic("")
}

func (i *IndexWriter) newSegmentName() string {
	panic("")
}

// ForceMerge Forces merge policy to merge segments until there are {@code <= maxNumSegments}. The actual
// merges to be executed are determined by the {@link MergePolicy}.
//
// This is a horribly costly operation, especially when you pass a small {@code
// maxNumSegments}; usually you should only call this if the index is static (will no longer be
// changed).
//
// Note that this requires free space that is proportional to the size of the index in your
// Directory: 2X if you are not using compound file format, and 3X if you are. For example, if
// your index size is 10 MB then you need an additional 20 MB free for this to complete (30 MB if
// you're using compound file format). This is also affected by the {@link Codec} that is used to
// execute the merge, and may result in even a bigger index. Also, it's best to call {@link
// #commit()} afterwards, to allow IndexWriter to free up disk space.
//
// If some but not all readers re-open while merging is underway, this will cause {@code > 2X}
// temporary space to be consumed as those new readers will then hold open the temporary segments
// at that time. It is best not to re-open readers while merging is running.
//
// The actual temporary usage could be much less than these figures (it depends on many
// factors).
//
// In general, once this completes, the total size of the index will be less than the size of
// the starting index. It could be quite a bit smaller (if there were many pending deletes) or
// just slightly smaller.
//
// If an Exception is hit, for example due to disk full, the index will not be corrupted and no
// documents will be lost. However, it may have been partially merged (some segments were merged
// but not all), and it's possible that one of the segments in the index will be in non-compound
// format even when using compound file format. This will occur when the Exception is hit during
// conversion of the segment into compound format.
//
// This call will merge those segments present in the index when the call started. If other
// threads are still adding documents and flushing segments, those newly created segments will not
// be merged unless you call forceMerge again.
func (i *IndexWriter) ForceMerge(maxNumSegments int) error {
	panic("")
}

// ForceMergev1 Just like {@link #forceMerge(int)}, except you can specify whether the call should block until
// all merging completes. This is only meaningful with a {@link MergeScheduler} that is able to
// run merges in background threads.
func (i *IndexWriter) ForceMergev1(maxNumSegments int, doWait bool) error {
	panic("")
}

// ForceMergeDeletes Just like #forceMergeDeletes(), except you can specify whether the call should block
// until the operation completes. This is only meaningful with a MergeScheduler that is
// able to run merges in background threads.
func (i *IndexWriter) ForceMergeDeletes(doWait bool) error {
	panic("")
}

// ForceMergeDeletesV1 Forces merging of all segments that have deleted documents. The actual merges to be executed
// are determined by the {@link MergePolicy}. For example, the default {@link TieredMergePolicy}
// will only pick a segment if the percentage of deleted docs is over 10%.
//   *
// <p>This is often a horribly costly operation; rarely is it warranted.
//   *
// <p>To see how many deletions you have pending in your index, call {@link
// IndexReader#numDeletedDocs}.
//   *
// <p><b>NOTE</b>: this method first flushes a new segment (if there are indexed documents), and
// applies all buffered deletes.
func (i *IndexWriter) ForceMergeDeletesV1() error {
	panic("")
}

// MaybeMerge Expert: asks the mergePolicy whether any merges are necessary now and if so, runs the requested
// merges and then iterate (test again if merges are needed) until no more merges are returned by
// the mergePolicy.
//   *
// <p>Explicit calls to maybeMerge() are usually not necessary. The most common case is when merge
// policy parameters have changed.
func (i *IndexWriter) MaybeMerge() error {
	panic("")
}

// GetMergingSegments Expert: to be used by a MergePolicy to avoid selecting merges for segments already
// being merged. The returned collection is not cloned, and thus is only safe to access if you
// hold IndexWriter's lock (which you do when IndexWriter invokes the MergePolicy).
func (i *IndexWriter) GetMergingSegments() []SegmentCommitInfo {
	panic("")
}

// HasPendingMerges Expert: returns true if there are merges waiting to be scheduled.
func (i *IndexWriter) HasPendingMerges() bool {
	panic("")
}

// Rollback Close the <code>IndexWriter</code> without committing any changes that have occurred since the
// last commit (or since it was opened, if commit hasn't been called). This removes any temporary
// files that had been created, after which the state of the index will be the same as it was when
// commit() was last called or when this writer was first opened. This also clears a previous call
// to prepareCommit.
func (i *IndexWriter) Rollback() error {
	panic("")
}

// DeleteAll Delete all documents in the index.
//
// This method will drop all buffered documents and will remove all segments from the index.
// This change will not be visible until a {@link #commit()} has been called. This method can be
// rolled back using {@link #rollback()}.
//
// NOTE: this method is much faster than using deleteDocuments( new MatchAllDocsQuery() ). Yet,
// this method also has different semantics compared to {@link #deleteDocuments(Query...)} since
// internal data-structures are cleared as well as all segment information is forcefully dropped
// anti-viral semantics like omitting norms are reset or doc value types are cleared. Essentially
// a call to {@link #deleteAll()} is equivalent to creating a new {@link IndexWriter} with {@link
// OpenMode#CREATE} which a delete query only marks documents as deleted.
//
// NOTE: this method will forcefully abort all merges in progress. If other threads are running
// {@link #forceMerge}, {@link #addIndexes(CodecReader[])} or {@link #forceMergeDeletes} methods,
// they may receive {@link MergePolicy.MergeAbortedException}s.
func (i *IndexWriter) DeleteAll() (int64, error) {
	panic("")
}

// AddIndexes Adds all segments from an array of indexes into this index.
//
// <p>This may be used to parallelize batch indexing. A large document collection can be broken
// into sub-collections. Each sub-collection can be indexed in parallel, on a different thread,
// process or machine. The complete index can then be created by merging sub-collection indexes
// with this method.
//
// <p><b>NOTE:</b> this method acquires the write lock in each directory, to ensure that no {@code
// IndexWriter} is currently open or tries to open while this is running.
//
// <p>This method is transactional in how Exceptions are handled: it does not commit a new
// segments_N file until all indexes are added. This means if an Exception occurs (for example
// disk full), then either no indexes will have been added or they all will have been.
//
// <p>Note that this requires temporary free space in the {@link Directory} up to 2X the sum of
// all input indexes (including the starting index). If readers/searchers are open against the
// starting index, then temporary free space required will be higher by the size of the starting
// index (see {@link #forceMerge(int)} for details).
//
// <p>This requires this index not be among those to be added.
//
// <p>All added indexes must have been created by the same Lucene version as this index.
func (i *IndexWriter) AddIndexes(dirs ...Directory) (int64, error) {
	panic("")
}

// AddIndexesByCodec Merges the provided indexes into this index.
//   *
// <p>The provided IndexReaders are not closed.
//   *
// <p>See {@link #addIndexes} for details on transactional semantics, temporary free space
// required in the Directory, and non-CFS segments on an Exception.
//   *
// <p><b>NOTE:</b> empty segments are dropped by this method and not added to this index.
//   *
// <p><b>NOTE:</b> this merges all given {@link LeafReader}s in one merge. If you intend to merge
// a large number of readers, it may be better to call this method multiple times, each time with
// a small set of readers. In principle, if you use a merge policy with a {@code mergeFactor} or
// {@code maxMergeAtOnce} parameter, you should pass that many readers in one call.
//   *
// <p><b>NOTE:</b> this method does not call or make use of the {@link MergeScheduler}, so any
// custom bandwidth throttling is at the moment ignored.
func (i *IndexWriter) AddIndexesByCodec(readers ...CodecReader) (int64, error) {
	panic("")
}

// A hook for extending classes to execute operations after pending added and deleted documents
// have been flushed to the Directory but before the change is committed (new segments_N file
// written).
func (i *IndexWriter) doAfterFlush() error {
	panic("")
}

// A hook for extending classes to execute operations before pending added and deleted documents
// are flushed to the Directory.
func (i *IndexWriter) doBeforeFlush() error {
	panic("")
}

// PrepareCommit Expert: prepare for commit. This does the first phase of 2-phase commit. This method does all
// steps necessary to commit changes since this writer was opened: flushes pending added and
// deleted docs, syncs the index files, writes most of next segments_N file. After calling this
// you must call either {@link #commit()} to finish the commit, or {@link #rollback()} to revert
// the commit and undo all changes done since the writer was opened.
//   *
// <p>You can also just call {@link #commit()} directly without prepareCommit first in which case
// that method will internally call prepareCommit.
func (i *IndexWriter) PrepareCommit() (int64, error) {
	panic("")
}

// FlushNextBuffer Expert: Flushes the next pending writer per thread buffer if available or the largest active
// non-pending writer per thread buffer in the calling thread. This can be used to flush documents
// to disk outside of an indexing thread. In contrast to {@link #flush()} this won't mark all
// currently active indexing buffers as flush-pending.
//
// <p>Note: this method is best-effort and might not flush any segments to disk. If there is a
// full flush happening concurrently multiple segments might have been flushed. Users of this API
// can access the IndexWriters current memory consumption via {@link #ramBytesUsed()}
func (i *IndexWriter) FlushNextBuffer() (bool, error) {
	panic("")
}

// SetLiveCommitData Sets the iterator to provide the commit user data map at commit time. Calling this method is
// considered a committable change and will be {@link #commit() committed} even if there are no
// other changes this writer. Note that you must call this method before {@link #prepareCommit()}.
// Otherwise it won't be included in the follow-on {@link #commit()}.
//   *
// <p><b>NOTE:</b> the iterator is late-binding: it is only visited once all documents for the
// commit have been written to their segments, before the next segments_N file is written
func (i *IndexWriter) SetLiveCommitData() {

}

// SetLiveCommitDataV1 Sets the commit user data iterator, controlling whether to advance the
// SegmentInfos#getVersion.
func (i *IndexWriter) SetLiveCommitDataV1(commitUserData *Iterator, doIncrementVersion bool) {
}

// GetLiveCommitData Returns the commit user data iterable previously set with #setLiveCommitData(Iterable),
// or null if nothing has been set yet.
//  Iterable[map[string]string]
func (i *IndexWriter) GetLiveCommitData() *Iterator {
	panic("")
}

// Commit Commits all pending changes (added and deleted documents, segment merges, added indexes, etc.)
// to the index, and syncs all referenced index files, such that a reader will see the changes and
// the index updates will survive an OS or machine crash or power loss. Note that this does not
// wait for any running background merges to finish. This may be a costly operation, so you should
// test the cost in your application and do it only when really necessary.
//   *
// <p>Note that this operation calls Directory.sync on the index files. That call should not
// return until the file contents and metadata are on stable storage. For FSDirectory, this calls
// the OS's fsync. But, beware: some hardware devices may in fact cache writes even during fsync,
// and return before the bits are actually on stable storage, to give the appearance of faster
// performance. If you have such a device, and it does not have a battery backup (for example)
// then on power loss it may still lose data. Lucene cannot guarantee consistency on such devices.
//   *
// <p>If nothing was committed, because there were no pending changes, this returns -1. Otherwise,
// it returns the sequence number such that all indexing operations prior to this sequence will be
// included in the commit point, and all other operations will not.
func (i *IndexWriter) Commit() (int64, error) {
	panic("")
}

// HasUncommittedChanges Returns true if there may be changes that have not been committed. There are cases where this
// may return true when there are no actual "real" changes to the index, for example if you've
// deleted by Term or Query but that Term or Query does not match any documents. Also, if a merge
// kicked off as a result of flushing a new segment during {@link #commit}, or a concurrent merged
// finished, this method may return true right after you had just called {@link #commit}.
func (i *IndexWriter) HasUncommittedChanges() bool {
	panic("")
}

//Returns true if there are any changes or deletes that are not flushed or applied.
func (i *IndexWriter) hasChangesInRam() bool {
	panic("")
}

// Flush Moves all in-memory segments to the Directory, but does not commit (fsync) them (call #commit for that).
func (i *IndexWriter) Flush() error {
	panic("")
}

// NumRamDocs Expert: Return the number of documents currently buffered in RAM.
func (i *IndexWriter) NumRamDocs() int {
	panic("")
}

// OnTragicEvent This method should be called on a tragic event ie. if a downstream class of the writer hits an
// unrecoverable exception. This method does not rethrow the tragic event exception.
//   *
// <p>Note: This method will not close the writer but can be called from any location without
// respecting any lock order
func (i *IndexWriter) OnTragicEvent(tragedy interface{}, location string) {
	panic("")
}

// GetTragicException If this {@code IndexWriter} was closed as a side-effect of a tragic exception, e.g. disk full
// while flushing a new segment, this returns the root cause exception. Otherwise (no tragic
// exception has occurred) it returns null.
func (i *IndexWriter) GetTragicException() interface{} {
	panic("")
}

// IsOpen Returns {@code true} if this IndexWriter is still open.
func (i *IndexWriter) IsOpen() bool {
	panic("")
}

// DeleteUnusedFiles Expert: remove any index files that are no longer used.
//   *
// <p>IndexWriter normally deletes unused files itself, during indexing. However, on Windows,
// which disallows deletion of open files, if there is a reader open on the index then those files
// cannot be deleted. This is fine, because IndexWriter will periodically retry the deletion.
//   *
// <p>However, IndexWriter doesn't try that often: only on open, close, flushing a new segment,
// and finishing a merge. If you don't do any of these actions with your IndexWriter, you'll see
// the unused files linger. If that's a problem, call this method to delete them (once you've
// closed the open readers that were preventing their deletion).
//   *
// <p>In addition, you can call this method to delete unreferenced index commits. This might be
// useful if you are using an {@link IndexDeletionPolicy} which holds onto index commits until
// some criteria are met, but those commits are no longer needed. Otherwise, those commits will be
// deleted the next time commit() is called.
func (i *IndexWriter) DeleteUnusedFiles() error {
	panic("")
}

// IncRefDeleter Record that the files referenced by this {@link SegmentInfos} are still in use.
func (i *IndexWriter) IncRefDeleter(segmentInfos *SegmentInfos) error {
	panic("")
}

// DecRefDeleter Record that the files referenced by this SegmentInfos are no longer in use. Only call
// this if you are sure you previously called #incRefDeleter.
func (i *IndexWriter) DecRefDeleter(segmentInfos *SegmentInfos) error {
	panic("")
}

// GetPendingNumDocs Returns the number of documents in the index including documents are being added
// (i.e., reserved).
func (i *IndexWriter) GetPendingNumDocs() int64 {
	panic("")
}

// GetMaxCompletedSequenceNumber Returns the highest sequence number across all completed
// operations, or 0 if no operations have finished yet. Still in-flight operations (in other
// threads) are not counted until they finish.
func (i *IndexWriter) GetMaxCompletedSequenceNumber() int64 {
	panic("")
}

// NumDeletesToMerge Returns the number of deletes a merge would claim back if the given segment is merged.
func (i *IndexWriter) NumDeletesToMerge() (int, error) {
	panic("")
}

// GetDocStats Returns accurate DocStats for this writer. The numDoc for instance can change after
// maxDoc is fetched that causes numDocs to be greater than maxDoc which makes it hard to get
// accurate document stats from IndexWriter.
func (i *IndexWriter) GetDocStats() *DocStats {
	panic("")
}

//------------------------------------

// IndexReaderWarmer If DirectoryReader#open(IndexWriter) has been called (ie, this writer is in near
// real-time mode), then after a merge completes, this class can be invoked to warm the reader on
// the newly merged segment, before the merge commits. This is not required for near real-time
// search, but will reduce search latency on opening a new near real-time reader after a merge
// completes.
type IndexReaderWarmer interface {

	// Warm Invoked on the LeafReader for the newly merged segment, before that segment is made
	// visible to near-real-time readers.
	Warm(reader LeafReader) error
}

type DocModifier interface {
	Run(docId int, readersAndUpdates *ReadersAndUpdates) error
}

//DocStats for this index
type DocStats struct {
	// The total number of docs in this index, including docs not yet flushed
	// (still in the RAM buffer), not counting deletions.
	maxDoc int

	// The total number of docs in this index, including docs not yet flushed (still in the RAM
	// buffer), and including deletions. <b>NOTE:</b> buffered deletions are not counted. If you
	// really need these to be counted you should call {@link IndexWriter#commit()} first.
	numDocs int
}
