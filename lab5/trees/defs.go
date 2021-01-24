package trees

// Node structure repreents tree's node
type Node struct {
	right *Node
	left  *Node
	value int
}

// Tree structure contains tree representation
type Tree struct {
	sequence []int
	root     *Node
}
