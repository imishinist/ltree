package main

import (
	"errors"
	"path/filepath"
	"strings"
)

type Splitter interface {
	Split(string) ([]string, error)
}

var (
	Separator string = string(filepath.Separator)

	// ErrInvalidPath is returned when path is invalid
	ErrInvalidPath error = errors.New("invalid path string")

	DefaultSplitter Splitter = &SingleSplitter{
		Sep:   Separator,
		Clean: false,
	}
)

func NewReverseSplitter(seps []string) Splitter {
	splitter := NewSplitter(seps)
	return &ReverseSplitter{
		splitter,
	}
}

type ReverseSplitter struct {
	splitter Splitter
}

func (s *ReverseSplitter) Split(path string) ([]string, error) {
	paths, err := s.splitter.Split(path)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(paths)/2; i++ {
		index := len(paths) - 1 - i
		paths[i], paths[index] = paths[index], paths[i]
	}
	return paths, nil
}

func NewSplitter(seps []string) Splitter {
	return &MultiSplitter{
		Sep:   seps,
		Clean: false,
	}
}

type MultiSplitter struct {
	Sep   []string
	Clean bool
}

func (s *MultiSplitter) Split(path string) ([]string, error) {
	replace_strings := make([]string, 0, 2*len(s.Sep))
	for _, s := range s.Sep {
		replace_strings = append(replace_strings, s)
		replace_strings = append(replace_strings, Separator)
	}
	r := strings.NewReplacer(replace_strings...)
	path = r.Replace(path)
	single_splitter := &SingleSplitter{
		Sep:   Separator,
		Clean: s.Clean,
	}
	return single_splitter.Split(path)
}

type SingleSplitter struct {
	Sep   string
	Clean bool
}

// clean the invalid path, and split by separator
func (s *SingleSplitter) Split(path string) ([]string, error) {
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
