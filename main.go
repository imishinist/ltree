package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	body, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	paths := strings.Split(string(body), "\n")

	root, err := NewTree(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, path := range paths {
		root.Merge(path)
	}

	printTree(root)
}

func printTree(root *Tree) {
	indents := []bool{false}
	doTree(root, 0, indents, true)
}

func doTree(root *Tree, depth int, indents []bool, last bool) {
	indent := ""
	for i := 0; i < depth; i++ {
		if indents[i] {
			indent += "│   "
		} else {
			indent += "    "
		}
	}
	if last {
		indent += "└── "
	} else {
		indent += "├── "
	}
	fmt.Printf("%s%s\n", indent, root.Name)
	for i, child := range root.Children {
		indentsTmp := make([]bool, len(indents))
		copy(indentsTmp, indents)
		if i == len(root.Children)-1 {
			indentsTmp = append(indentsTmp, false)
		} else {
			indentsTmp = append(indentsTmp, true)
		}
		doTree(child, depth+1, indentsTmp, i == len(root.Children)-1)
	}
}
