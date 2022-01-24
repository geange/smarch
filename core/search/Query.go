package search

// Query The abstract base class for queries.
//
// <p>Instantiable subclasses are:
//
// <ul>
//   <li>TermQuery}
//   <li>BooleanQuery}
//   <li>WildcardQuery}
//   <li>PhraseQuery}
//   <li>PrefixQuery}
//   <li>MultiPhraseQuery}
//   <li>FuzzyQuery}
//   <li>RegexpQuery}
//   <li>TermRangeQuery}
//   <li>PointRangeQuery}
//   <li>ConstantScoreQuery}
//   <li>DisjunctionMaxQuery}
//   <li>MatchAllDocsQuery}
// </ul>
//
// See also additional queries available in the <a
// href="{@docRoot}/../queries/overview-summary.html">Queries module</a>
type Query interface {
}
