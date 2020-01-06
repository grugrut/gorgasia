package org

// Node indicates a block of org node.
type Node struct {
	Title    string
	Contents string
	Children []Node
}
