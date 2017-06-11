package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	root, err := NewTree(dir)
	if err != nil {
		log.Fatal(err)
	}
	printTree(root, 0)
}

func printTree(root *Tree, depth int) {
	indent := ""
	for i := 0; i < depth; i++ {
		if i == depth-1 {
			indent += "└── "
			//indent += "─── "
			// indent += "└── "
			//indent += "├── "
		} else {
			indent += "    "
		}
	}
	fmt.Printf("%s%s\n", indent, root.Name)
	for _, child := range root.Children {
		printTree(child, depth+1)
	}
}
