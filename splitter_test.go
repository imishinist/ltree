package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Object *Splitter
		Input  string
		Output []string
		Error  error
	}{
		{&Splitter{Separator, true}, "/usr/bin/cd", []string{"usr", "bin", "cd"}, nil},
		{&Splitter{Separator, true}, "/usr/bin/", []string{"usr", "bin"}, nil},
		{&Splitter{Separator, true}, "/usr/bin//tmp", []string{"usr", "bin", "tmp"}, nil},
		{&Splitter{Separator, true}, "usr/bin", []string{"usr", "bin"}, nil},
		{&Splitter{Separator, true}, "usr/bin/", []string{"usr", "bin"}, nil},
		{&Splitter{Separator, true}, "./usr/bin", []string{"usr", "bin"}, nil},
		{&Splitter{Separator, false}, "/usr/bin//tmp", []string{"usr", "bin", "", "tmp"}, nil},
		{&Splitter{Separator, false}, "usr/bin", []string{"usr", "bin"}, nil},
		{&Splitter{Separator, false}, "usr/bin/", []string{"usr", "bin"}, nil},
		{&Splitter{Separator, false}, "./usr/bin", []string{".", "usr", "bin"}, nil},
		{DefaultSplitter, "", nil, ErrInvalidPath},
	}

	for _, c := range cases {
		output, err := c.Object.Split(c.Input)
		assert.Equal(c.Output, output)
		assert.Equal(c.Error, err)
	}
}
