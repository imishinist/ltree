package main

// Tree contains file or directory status
type Tree struct {
	Parent   *Tree
	Children []*Tree
	Name     string
	splitter Splitter
}

// NewTree returns Node tree by path
func NewTree(path string) (*Tree, error) {
	splitter := DefaultSplitter
	return NewSTree(path, splitter)
}

func NewSTree(path string, splitter Splitter) (*Tree, error) {
	paths, err := splitter.Split(path)
	if err != nil {
		return nil, err
	}
	tree, err := pathsToTree(paths, splitter)
	if err != nil {
		return nil, err
	}
	tree.splitter = splitter
	return tree, nil
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
	paths, err := n.splitter.Split(path)
	if err != nil {
		return err
	}
	now := n
	for i, path := range paths {
		if now.Name == path {
			continue
		}
		if child := now.Child(path); child == nil {
			tree, err := pathsToTree(paths[i:], n.splitter)
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
func pathsToTree(paths []string, splitter Splitter) (*Tree, error) {
	var nodes []*Tree
	for _, path := range paths {
		node := &Tree{
			Parent:   nil,
			Children: []*Tree{},
			Name:     path,
			splitter: splitter,
		}
		nodes = append(nodes, node)
	}
	for i := 1; i < len(nodes); i++ {
		nodes[i].Parent = nodes[i-1]
		nodes[i-1].Children = append(nodes[i-1].Children, nodes[i])
	}
	return nodes[0], nil
}
