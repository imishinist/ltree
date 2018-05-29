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

	delims := []string{".", "/", ":"}
	splitter := NewSplitter(delims)

	root, err := NewSTree("<root>", splitter)
	if err != nil {
		log.Fatal(err)
	}

	for _, path := range paths {
		root.Merge(path)
	}

	printTree(root)
}

func printTree(root *Tree) {
	indents := []bool{}
	doTree(root, indents, true)
}

func doTree(root *Tree, indents []bool, last bool) {
	indent := ""
	for i := 0; i < len(indents); i++ {
		if i == len(indents)-1 {
			if last {
				indent += "└── "
			} else {
				indent += "├── "
			}
		} else {
			if indents[i] {
				indent += "│   "
			} else {
				indent += "    "
			}
		}
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
		doTree(child, indentsTmp, i == len(root.Children)-1)
	}
}
