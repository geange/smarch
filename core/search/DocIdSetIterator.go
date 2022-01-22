package search

// DocIdSetIterator This abstract class defines methods to iterate over a set of non-decreasing doc ids. Note that
// this class assumes it iterates on doc Ids, and therefore #NO_MORE_DOCS} is set to {@value
// #NO_MORE_DOCS} in order to be used as a sentinel object. Implementations of this class are
// expected to consider Integer#MAX_VALUE} as an invalid value.
type DocIdSetIterator interface {
}
