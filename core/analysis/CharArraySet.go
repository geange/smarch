package analysis

// CharArraySet A simple class that stores Strings as char[]'s in a hash table. Note that this is not a general
// purpose class. For example, it cannot remove items from the set, nor does it resize its hash
// table to be smaller, etc. It is designed to be quick to test if a char[] is in the set without
// the necessity of converting it to a String first.
//
// Please note: This class implements {@link java.util.Set Set} but does not behave like
// it should in all cases. The generic type is {@code Set<Object>}, because you can add any object
// to it, that has a string representation. The add methods will use {@link Object#toString} and
// store the result using a {@code char[]} buffer. The same behavior have the {@code contains()}
// methods. The {@link #iterator()} returns an {@code Iterator<char[]>}.
type CharArraySet struct {
}
