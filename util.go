package go_fileutil

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func callerPath() (string, error) {
	_, p, _, ok := runtime.Caller(2)
	if !ok {
		return "", fmt.Errorf("can't get caller path")
	}
	return p, nil
}

// ReadFile reads the named file and returns the contents.
// A successful call returns err == nil, not err == EOF.
// Because ReadFile reads the whole file, it does not treat an EOF from Read
// as an error to be reported.
func ReadFile(name string) ([]byte, error) {
	p, err := callerPath()
	if err != nil {
		return nil, err
	}
	basePath := filepath.Dir(p)
	return os.ReadFile(filepath.Join(basePath, name))
}

// Open opens the named file for reading. If successful, methods on
// the returned file can be used for reading; the associated file
// descriptor has mode O_RDONLY.
// If there is an error, it will be of type *PathError.
func Open(name string) (*os.File, error) {
	p, err := callerPath()
	if err != nil {
		return nil, err
	}
	return os.Open(p)
}
