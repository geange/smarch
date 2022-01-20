package index

import "io"

type IndexReader interface {
	io.Closer
}
