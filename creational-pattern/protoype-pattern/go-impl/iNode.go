package main

// prototype interface
type Inode interface {
	print(string)
	clone() Inode
}
