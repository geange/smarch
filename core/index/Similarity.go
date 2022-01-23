package index

// Similarity defines the components of Lucene scoring.
//
// Expert: Scoring API.
//
// This is a low-level API, you should only extend this API if you want to implement an
// information retrieval model. If you are instead looking for a convenient way to alter
// Lucene's scoring, consider just tweaking the default implementation: BM25Similarity or
// extend SimilarityBase, which makes it easy to compute a score from index statistics.
//
// Similarity determines how Lucene weights terms, and Lucene interacts with this class at both
// index-time and query-time.
//
// Indexing Time At indexing time, the indexer calls
// #computeNorm(FieldInvertState), allowing the Similarity implementation to set a per-document
// value for the field that will be later accessible via
// org.apache.lucene.index.LeafReader#getNormValues(String). Lucene makes no assumption about what
// is in this norm, but it is most useful for encoding length normalization information.
//
// Implementations should carefully consider how the normalization is encoded: while Lucene's
// BM25Similarity encodes length normalization information with SmallFloat into a
// single byte, this might not be suitable for all purposes.
//
// Many formulas require the use of average document length, which can be computed via a
// combination of CollectionStatistics#sumTotalTermFreq() and
// CollectionStatistics#docCount().
//
// Additional scoring factors can be stored in named NumericDocValuesFields and accessed
// at query-time with org.apache.lucene.index.LeafReader#getNumericDocValues(String).
// However this should not be done in the Similarity but externally, for instance by using
// FunctionScoreQuery.
//
// Finally, using index-time boosts (either via folding into the normalization byte or via
// DocValues), is an inefficient way to boost the scores of different fields if the boost will be
// the same for every document, instead the Similarity can simply take a constant boost parameter
// C, and PerFieldSimilarityWrapper can return different instances with different
// boosts depending upon field name.
//
// Query time At query-time, Queries interact with the Similarity via these
// steps:
//
//
//   The #scorer(float, CollectionStatistics, TermStatistics...) method is called a
//       single time, allowing the implementation to compute any statistics (such as IDF, average
//       document length, etc) across the entire collection. The TermStatistics and
//       CollectionStatistics passed in already contain all of the raw statistics involved,
//       so a Similarity can freely use any combination of statistics without causing any additional
//       I/O. Lucene makes no assumption about what is stored in the returned
//       Similarity.SimScorer object.
//   Then SimScorer#score(float, long) is called for every matching document to compute
//       its score.
//
//
// <a id="explaintime">Explanations</a> When
// IndexSearcher#explain(org.apache.lucene.search.Query, int) is called, queries consult the
// Similarity's DocScorer for an explanation of how it computed its score. The query passes in a the
// document id and an explanation of how the frequency was computed.
type Similarity interface {

	// Computes the normalization value for a field, given the accumulated state of term processing
	// for this field (see FieldInvertState).
	//
	// Matches in longer fields are less precise, so implementations of this method usually set
	// smaller values when state.getLength() is large, and larger values when
	// state.getLength() is small.
	//
	// Note that for a given term-document frequency, greater unsigned norms must produce scores
	// that are lower or equal, ie. for two encoded norms n1 and n2 so that
	// Long.compareUnsigned(n1, n2) > 0 then SimScorer.score(freq, n1) <=
	// SimScorer.score(freq, n2) for any legal freq.
	//
	// 0 is not a legal norm, so 1 is the norm that produces the highest scores.
}
