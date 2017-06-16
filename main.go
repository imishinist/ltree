package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	root := &Tree{Name: ".", Children: []*Tree{}}
	for _, file := range files {
		err := root.Merge(file.Name())
		if err != nil {
			fmt.Println(err)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
	printTree(root, 0, true)
}

func printTree(root *Tree, depth int, last bool) {
	indent := ""
	for i := 0; i < depth; i++ {
		if i == depth-1 {
			if last {
				indent += "└── "
			} else {
				indent += "├── "
			}
		} else {
			indent += "    "
		}
	}
	fmt.Printf("%s%s\n", indent, root.Name)
	for i, child := range root.Children {
		printTree(child, depth+1, i == len(root.Children)-1)
	}
}
