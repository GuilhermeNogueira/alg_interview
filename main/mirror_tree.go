package main

type Node struct {
	Value     int
	LeftNode  *Node
	RightNode *Node
}

func (n *Node) AddLeftNode(node *Node) {
	n.LeftNode = node
}

func (n *Node) AddRightNode(node *Node) {
	n.RightNode = node
}

func isEqual(n *Node, node *Node) bool {

	if n == nil && node == nil {
		return true
	}

	if (node == nil && n != nil) || (n == nil && node != nil) {
		return false
	}

	if n.Value != node.Value {
		return false
	}

	if n.LeftNode == nil && node.LeftNode != nil {
		return false
	}

	if n.RightNode == nil && node.RightNode != nil {
		return false
	}

	if !isEqual(n.RightNode, node.RightNode) {
		return false
	}

	if !isEqual(n.LeftNode, node.LeftNode) {
		return false
	}

	return true
}

func mirrorTree(node *Node) *Node {

	if node == nil {
		return nil
	}

	left := node.LeftNode
	right := node.RightNode

	node.AddRightNode(mirrorTree(left))

	node.AddLeftNode(mirrorTree(right))

	return node
}
