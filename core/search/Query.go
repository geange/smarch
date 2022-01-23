package search

// Query The abstract base class for queries.
//
// <p>Instantiable subclasses are:
//
// <ul>
//   <li>{@link TermQuery}
//   <li>{@link BooleanQuery}
//   <li>{@link WildcardQuery}
//   <li>{@link PhraseQuery}
//   <li>{@link PrefixQuery}
//   <li>{@link MultiPhraseQuery}
//   <li>{@link FuzzyQuery}
//   <li>{@link RegexpQuery}
//   <li>{@link TermRangeQuery}
//   <li>{@link PointRangeQuery}
//   <li>{@link ConstantScoreQuery}
//   <li>{@link DisjunctionMaxQuery}
//   <li>{@link MatchAllDocsQuery}
// </ul>
//
// See also additional queries available in the <a
// href="{@docRoot}/../queries/overview-summary.html">Queries module</a>
type Query interface {
}
