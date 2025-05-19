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

// ReadFile reads the named file with caller file path, and
// call os.ReadFile to read content
func ReadFile(name string) ([]byte, error) {
	p, err := callerPath()
	if err != nil {
		return nil, err
	}
	basePath := filepath.Dir(p)
	return os.ReadFile(filepath.Join(basePath, name))
}

// Open opens the named file for reading with caller file path,
// and call os.Open to construct *os.File
func Open(name string) (*os.File, error) {
	p, err := callerPath()
	if err != nil {
		return nil, err
	}
	return os.Open(p)
}
