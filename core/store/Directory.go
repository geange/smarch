package store

type Directory interface {

	// FileLength Returns the byte length of a file in the directory.
	//
	// This method must throw either {@link NoSuchFileException} or {@link FileNotFoundException}
	// if code name points to a non-existing file.
	//
	// * name the name of an existing file.
	// * in case of I/O error
	FileLength(name string) (int64, error)
}
