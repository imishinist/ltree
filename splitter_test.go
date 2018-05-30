package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleSplit(t *testing.T) {
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

func TestMultiSplit(t *testing.T) {
	assert := assert.New(t)
	seps := []string{Separator}

	cases := []struct {
		Object *MultiSplitter
		Input  string
		Output []string
		Error  error
	}{
		{&MultiSplitter{seps, true}, "/usr/bin/cd", []string{"usr", "bin", "cd"}, nil},
		{&MultiSplitter{seps, true}, "/usr/bin/", []string{"usr", "bin"}, nil},
		{&MultiSplitter{seps, true}, "/usr/bin//tmp", []string{"usr", "bin", "tmp"}, nil},
		{&MultiSplitter{seps, true}, "usr/bin", []string{"usr", "bin"}, nil},
		{&MultiSplitter{seps, true}, "usr/bin/", []string{"usr", "bin"}, nil},
		{&MultiSplitter{seps, true}, "./usr/bin", []string{"usr", "bin"}, nil},
		{&MultiSplitter{[]string{"/", ","}, true}, "./usr,bin", []string{"usr", "bin"}, nil},
		{&MultiSplitter{[]string{"/", ","}, true}, "./usr,,bin", []string{"usr", "bin"}, nil},
		{&MultiSplitter{seps, false}, "/usr/bin//tmp", []string{"usr", "bin", "", "tmp"}, nil},
		{&MultiSplitter{seps, false}, "usr/bin", []string{"usr", "bin"}, nil},
		{&MultiSplitter{seps, false}, "usr/bin/", []string{"usr", "bin"}, nil},
		{&MultiSplitter{seps, false}, "./usr/bin", []string{".", "usr", "bin"}, nil},
	}

	for _, c := range cases {
		output, err := c.Object.Split(c.Input)
		assert.Equal(c.Output, output)
		assert.Equal(c.Error, err)
	}
}

func TestReverseSplit(t *testing.T) {
	assert := assert.New(t)
	seps := []string{Separator}

	cases := []struct {
		Object Splitter
		Input  string
		Output []string
		Error  error
	}{
		{NewReverseSplitter(seps), "/usr/bin/cd", []string{"cd", "bin", "usr"}, nil},
		{NewReverseSplitter(seps), "/usr/bin/", []string{"bin", "usr"}, nil},
		{NewReverseSplitter(seps), "/usr/bin//tmp", []string{"tmp", "", "bin", "usr"}, nil},
		{NewReverseSplitter(seps), "usr/bin", []string{"bin", "usr"}, nil},
		{NewReverseSplitter(seps), "usr/bin/", []string{"bin", "usr"}, nil},
		{NewReverseSplitter(seps), "./usr/bin", []string{"bin", "usr", "."}, nil},
		{NewReverseSplitter([]string{"/", ","}), "./usr,bin", []string{"bin", "usr", "."}, nil},
		{NewReverseSplitter([]string{"/", ","}), "./usr,,bin", []string{"bin", "", "usr", "."}, nil},
		{NewReverseSplitter(seps), "/usr/bin//tmp", []string{"tmp", "", "bin", "usr"}, nil},
	}

	for _, c := range cases {
		output, err := c.Object.Split(c.Input)
		assert.Equal(c.Output, output)
		assert.Equal(c.Error, err)
	}
}
