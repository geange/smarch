package index

// TermsHashPerField This class stores streams of information per term without knowing the size of the
// stream ahead of time. Each stream typically encodes one level of information like term frequency
// per document or term proximity. Internally this class allocates a linked list of slices that can
// be read by a ByteSliceReader for each term. Terms are first deduplicated in a BytesRefHash
// once this is done internal data-structures point to the current offset of each stream that can be
// written to.
//
// 此类存储每个术语的信息流，而无需提前知道流的大小。 每个流通常编码一个级别的信息，例如每个文档的词频或词的接近度。
// 在内部，此类分配切片的链接列表，ByteSliceReader 可以为每个术语读取这些切片。
// 完成后，首先在 BytesRefHash 中对术语进行重复数据删除，内部数据结构指向可以写入的每个流的当前偏移量。
type TermsHashPerField struct {
}
