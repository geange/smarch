package search

// Sort Encapsulates sort criteria for returned hits.
// A Sort can be created with an empty constructor, yielding an object that will instruct
// searches to return their hits sorted by relevance; or it can be created with one or more
// SortFields.
type Sort interface {
}
