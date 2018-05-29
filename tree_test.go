package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func create(name string) *Tree {
	var splitter Splitter = DefaultSplitter
	return &Tree{nil, []*Tree{}, name, splitter}
}

func concat(t *Tree, children ...*Tree) *Tree {
	// Copy
	c := &Tree{
		Parent:   t.Parent,
		Children: t.Children,
		Name:     t.Name,
		splitter: t.splitter,
	}
	for _, child := range children {
		child.Parent = c
		c.Children = append(c.Children, child)
	}
	return c
}

func TestPathsToTree(t *testing.T) {
	assert := assert.New(t)
	var splitter Splitter = DefaultSplitter

	cases := []struct {
		Input  []string
		Splitter Splitter
		Output *Tree
	}{
		{[]string{"usr", "bin"}, splitter, concat(create("usr"), create("bin"))},
	}
	for _, c := range cases {
		output, _ := pathsToTree(c.Input, c.Splitter)
		assert.Equal(c.Output, output)
	}
}

func TestMerge(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Object *Tree
		Input  string
		Output *Tree
		Error  error
	}{
		{create("/"), "/usr/bin", concat(create("/"), concat(create("usr"), create("bin"))), nil},
		{create("."), "./usr/bin", concat(create("."), concat(create("usr"), create("bin"))), nil},
		{create("/"), "", create("/"), ErrInvalidPath},
		{concat(create("/"), concat(create("usr"), create("bin"))), "/usr/lib", concat(create("/"), concat(create("usr"), create("bin"), create("lib"))), nil},
	}

	for _, c := range cases {
		in := c.Object
		err := in.Merge(c.Input)
		assert.Equal(c.Output, in)
		assert.Equal(c.Error, err)
	}
}

func TestIsRoot(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Input  *Tree
		Output bool
	}{
		{create("usr"), true},
		{concat(create("usr"), create("bin")).Children[0], false},
	}

	for _, c := range cases {
		assert.Equal(c.Output, c.Input.IsRoot())
	}
}

func TestIsLeaf(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		Input  *Tree
		Output bool
	}{
		{create("usr"), true},
		{concat(create("usr"), create("bin")), false},
		{concat(create("usr"), create("bin")).Children[0], true},
	}

	for _, c := range cases {
		assert.Equal(c.Output, c.Input.IsLeaf())
	}
}
