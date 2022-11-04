package main

import "fmt"

// file.go: Component interface
type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
	if f.getName() == keyword {
		fmt.Printf("found file with keyword %s\n", keyword)
	}
}

func (f *File) getName() string {
	return f.name
}
