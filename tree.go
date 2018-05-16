package main

import (
	"errors"
	"path/filepath"
	"strings"
)

// ErrInvalidPath is returned when path is invalid
var ErrInvalidPath = errors.New("invalid path string")

// Tree contains file or directory status
type Tree struct {
	Parent   *Tree
	Children []*Tree
	Name     string
}

// NewTree returns Node tree by path
func NewTree(path string) (*Tree, error) {
	paths, err := splitPath(path)
	if err != nil {
		return nil, err
	}
	return pathsToTree(paths)
}

// IsRoot returns whether or not it's root
func (n *Tree) IsRoot() bool {
	return n.Parent == nil
}

// IsLeaf returns wether or not it's leaf
func (n *Tree) IsLeaf() bool {
	return len(n.Children) == 0
}

func (n *Tree) Merge(path string) error {
	paths, err := splitPath(path)
	if err != nil {
		return err
	}
	now := n
	for i, path := range paths {
		if child := now.Child(path); child == nil {
			tree, err := pathsToTree(paths[i:])
			if err != nil {
				return err
			}
			tree.Parent = now
			now.Children = append(now.Children, tree)
			break
		} else {
			now = child
		}
	}
	return nil
}

// Child searches child from Children by name
func (n *Tree) Child(name string) *Tree {
	for _, child := range n.Children {
		if child.Name == name {
			return child
		}
	}
	return nil
}

// path array convert to Node tree
func pathsToTree(paths []string) (*Tree, error) {
	var nodes []*Tree
	for _, path := range paths {
		node := &Tree{
			Parent:   nil,
			Children: []*Tree{},
			Name:     path,
		}
		nodes = append(nodes, node)
	}
	for i := 1; i < len(nodes); i++ {
		nodes[i].Parent = nodes[i-1]
		nodes[i-1].Children = append(nodes[i-1].Children, nodes[i])
	}
	return nodes[0], nil
}

// clean the invalid path, and split by separator
func splitPath(path string) ([]string, error) {
	if path == "" {
		return nil, ErrInvalidPath
	}
	path = filepath.Clean(path)
	dir, file := filepath.Split(path)

	dirs := strings.Split(dir, string(filepath.Separator))

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
