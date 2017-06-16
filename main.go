package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
	root := &Tree{Name: "/", Children: []*Tree{}}
	for _, file := range files {
		path, err := filepath.Abs(file.Name())
		if err != nil {
			fmt.Println(err)
		}
		err = root.Merge(path)
		if err != nil {
			fmt.Println(err)
		}
	}

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
