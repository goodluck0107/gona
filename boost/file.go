package boost

import (
	"fmt"
	"os"
)

const (
	// ReadFile is the type to open file for read
	ReadFile int = iota
	// WriteFile is the type to open file for write
	WriteFile
	// NewFile is the type to open file for create new file
	NewFile
	// AppendFile is the type to open file for appending new file
	AppendFile
)

// OpenFile opens file by type
func OpenFile(typ int, f string) (*os.File, error) {
	switch typ {
	case ReadFile:
		return os.OpenFile(f, os.O_RDONLY|os.O_CREATE, 0o644)
	case WriteFile:
		return os.OpenFile(f, os.O_WRONLY|os.O_CREATE, 0o644)
	case NewFile:
		return os.OpenFile(f, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0o644)
	case AppendFile:
		return os.OpenFile(f, os.O_APPEND|os.O_CREATE, 0o644)
	}
	return nil, fmt.Errorf("file open type %d is not supported", typ)
}

// RemoveFile removes file if exists
func RemoveFile(f string) error {
	if ExistsFile(f) {
		return os.Remove(f)
	}
	return nil
}

// ExistsFile returns if exists
func ExistsFile(f string) bool {
	if _, err := os.Stat(f); err != nil {
		return os.IsExist(err)
	}
	return true
}
