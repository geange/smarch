package index

var (
	IndexFileNamesInstance = &IndexFileNames{}
)

// IndexFileNames This class contains useful constants representing filenames and extensions used by lucene,
// as well as convenience methods for querying whether a file name matches an extension
// (#matchesExtension(String, String) matchesExtension), as well as generating file names from a
// segment name, generation and extension ( #fileNameFromGeneration(String, String, long)
// fileNameFromGeneration, #segmentFileName(String, String, String) segmentFileName).
//
// NOTE: extensions used by codecs are not listed here. You must interact with the Codec directly.
type IndexFileNames struct {
}

// StripSegmentName Strips the segment name out of the given file name. If you used #segmentFileName or
// #fileNameFromGeneration to create your files, then this method simply removes whatever
// comes before the first '.', or the second '_' (excluding both).
//
// the filename with the segment name removed, or the given filename if it does not
// contain a '.' and '_'.
func (i *IndexFileNames) StripSegmentName(fileName string) string {
	panic("")
}
