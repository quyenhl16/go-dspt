package main

import (
	"fmt"
)

// Refer to: https://refactoring.guru/design-patterns/composite
/*
Composite is a structural design pattern that lets you compose objects into tree structures
and then work with these structures as if they were individual objects.
*/

// Component interface (common interface for files and folders)
type Component interface {
	Show(indent string)
}

// File is a Leaf
type File struct {
	Name string
}

func (f *File) Show(indent string) {
	fmt.Println(indent+"- File:", f.Name)
}

// Folder is a Composite (can contain files or other folders)
type Folder struct {
	Name     string
	Children []Component
}

func (f *Folder) Show(indent string) {
	fmt.Println(indent+"+ Folder:", f.Name)
	for _, child := range f.Children {
		child.Show(indent + "  ")
	}
}

// Add a component to the folder
func (f *Folder) Add(component Component) {
	f.Children = append(f.Children, component)
}

func main() {
	// Create files
	file1 := &File{Name: "file1.txt"}
	file2 := &File{Name: "file2.txt"}
	file3 := &File{Name: "file3.txt"}

	// Create folders
	folder1 := &Folder{Name: "Documents"}
	folder2 := &Folder{Name: "Pictures"}

	// Add files to folders
	folder1.Add(file1)
	folder1.Add(file2)

	// Add another folder inside
	folder2.Add(file3)
	folder1.Add(folder2)

	// Show folder structure
	folder1.Show("")
}
