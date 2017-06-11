package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitPath(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Input  string
		Output []string
		Error  error
	}{
		{"/usr/bin/cd", []string{"usr", "bin", "cd"}, nil},
		{"/usr/bin/", []string{"usr", "bin"}, nil},
		{"/usr/bin//tmp", []string{"usr", "bin", "tmp"}, nil},
		{"/", []string{}, nil},
		{"", nil, ErrInvalidPath},
	}

	for _, c := range cases {
		output, err := splitPath(c.Input)
		assert.Equal(c.Output, output)
		assert.Equal(c.Error, err)
	}
}

func TestPathsToTree(t *testing.T) {
	assert := assert.New(t)

	usrTree := &Tree{nil, []*Tree{}, "usr"}
	binTree := &Tree{usrTree, []*Tree{}, "bin"}
	usrTree.Children = append(usrTree.Children, binTree)

	cases := []struct {
		Input  []string
		Output *Tree
	}{
		{[]string{"usr", "bin"}, usrTree},
	}
	for _, c := range cases {
		output, _ := pathsToTree(c.Input)
		assert.Equal(c.Output, output)
	}
}

func TestMerge(t *testing.T) {
	assert := assert.New(t)
	in := &Tree{nil, []*Tree{}, "/"}
	root := &Tree{nil, []*Tree{}, "/"}
	usrTree := &Tree{root, []*Tree{}, "usr"}
	root.Children = append(root.Children, usrTree)
	binTree := &Tree{usrTree, []*Tree{}, "bin"}
	usrTree.Children = append(usrTree.Children, binTree)

	cases := []struct {
		Input  string
		Output *Tree
	}{
		{"/usr/bin", root},
	}

	for _, c := range cases {
		in.Merge(c.Input)
		assert.Equal(c.Output, in)
	}
}

func TestIsRoot(t *testing.T) {
	assert := assert.New(t)

	usrTree := &Tree{nil, []*Tree{}, "usr"}
	binTree := &Tree{usrTree, []*Tree{}, "bin"}
	usrTree.Children = append(usrTree.Children, binTree)

	cases := []struct {
		Input  *Tree
		Output bool
	}{
		{usrTree, true},
		{binTree, false},
	}

	for _, c := range cases {
		assert.Equal(c.Output, c.Input.IsRoot())
	}
}

func TestIsLeaf(t *testing.T) {
	assert := assert.New(t)

	usrTree := &Tree{nil, []*Tree{}, "usr"}
	binTree := &Tree{usrTree, []*Tree{}, "bin"}
	usrTree.Children = append(usrTree.Children, binTree)

	cases := []struct {
		Input  *Tree
		Output bool
	}{
		{usrTree, false},
		{binTree, true},
	}

	for _, c := range cases {
		assert.Equal(c.Output, c.Input.IsLeaf())
	}
}
