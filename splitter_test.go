package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Object *SingleSplitter
		Input  string
		Output []string
		Error  error
	}{
		{&SingleSplitter{Separator, true}, "/usr/bin/cd", []string{"usr", "bin", "cd"}, nil},
		{&SingleSplitter{Separator, true}, "/usr/bin/", []string{"usr", "bin"}, nil},
		{&SingleSplitter{Separator, true}, "/usr/bin//tmp", []string{"usr", "bin", "tmp"}, nil},
		{&SingleSplitter{Separator, true}, "usr/bin", []string{"usr", "bin"}, nil},
		{&SingleSplitter{Separator, true}, "usr/bin/", []string{"usr", "bin"}, nil},
		{&SingleSplitter{Separator, true}, "./usr/bin", []string{"usr", "bin"}, nil},
		{&SingleSplitter{Separator, false}, "/usr/bin//tmp", []string{"usr", "bin", "", "tmp"}, nil},
		{&SingleSplitter{Separator, false}, "usr/bin", []string{"usr", "bin"}, nil},
		{&SingleSplitter{Separator, false}, "usr/bin/", []string{"usr", "bin"}, nil},
		{&SingleSplitter{Separator, false}, "./usr/bin", []string{".", "usr", "bin"}, nil},
	}

	for _, c := range cases {
		output, err := c.Object.Split(c.Input)
		assert.Equal(c.Output, output)
		assert.Equal(c.Error, err)
	}
}
