package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	p1 = "├──" //Todo change to appropriate name

	p2 = "│  "

	p3 = "└── "
)

func walkDir(dir string) []string {
	elements, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	paths := make([]string, 0)
	for _, element := range elements {
		if element.IsDir() {
			paths = append(paths, walkDir(filepath.Join(dir, element.Name()))...)
		} else {
			paths = append(paths, filepath.Join(dir, element.Name()))
		}
	}
	return paths
}

func deleteSeparator(paths string) []string {
	p := strings.Split(paths, "/")
	return p[1:]
}

func makeBranch(paths []string) string {
	var tree string
	lastIndex := len(paths) - 1
	for i, path := range paths {
		if path != "//duplicate" {
			if i == lastIndex { 
				tree = tree + strings.Repeat(p2, i) + p3 + path + "\n"
			} else {
				tree = tree + strings.Repeat(p2, i) + p1 + path + "\n"
			}
		} 
	}
	return tree
}

func init() {
	// --help option
	flag.Usage = func() {
		fmt.Printf("Usage: %v [OPTIONS] COMMAND \n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func tree(paths []string, rootDir string) {
	fmt.Println(rootDir)
	var branch string
	for _, path := range paths {
		d := deleteSeparator(path)
		for i, e := range d[:len(d)-1] {
			if strings.Contains(branch, e) {
				d[i] = "//duplicate" // add prefix
			}
		}
		branch = branch + makeBranch(d)
	}
	fmt.Println(branch)
}

func main() {
	flag.Parse()
	args := flag.Args() 
	rootDir := args[0]
	paths := walkDir(rootDir)
	tree(paths, rootDir)
}