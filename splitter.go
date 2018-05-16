package main

import (
	"errors"
	"path/filepath"
	"strings"
)

var (
	Separator string = string(filepath.Separator)

	// ErrInvalidPath is returned when path is invalid
	ErrInvalidPath error = errors.New("invalid path string")

	DefaultSplitter *Splitter = &Splitter{
		Sep:   Separator,
		Clean: false,
	}
)

type Splitter struct {
	Sep   string
	Clean bool
}

// clean the invalid path, and split by separator
func (s *Splitter) Split(path string) ([]string, error) {
	if path == "" {
		return nil, ErrInvalidPath
	}

	if s.Clean {
		path = filepath.Clean(path)
	}
	dir, file := filepath.Split(path)

	dirs := strings.Split(dir, Separator)

	if filepath.IsAbs(path) {
		dirs = dirs[1 : len(dirs)-1]
	} else {
		dirs = dirs[:len(dirs)-1]
	}

	if file != "" {
		dirs = append(dirs, file)
	}

	return dirs, nil
}

func splitPath(path string) ([]string, error) {
	return DefaultSplitter.Split(path)
}