package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTreeEq(t *testing.T) {

	myTree := &Node{
		Value: 1,
		LeftNode: &Node{
			Value: 2,
			LeftNode: &Node{
				Value:     4,
				LeftNode:  nil,
				RightNode: nil,
			},
			RightNode: &Node{
				Value:     5,
				LeftNode:  nil,
				RightNode: nil,
			},
		},
		RightNode: &Node{
			Value: 3,
			LeftNode: &Node{
				Value:     6,
				LeftNode:  nil,
				RightNode: nil,
			},
			RightNode: nil,
		},
	}

	reversedTree := &Node{
		Value: 1,
		LeftNode: &Node{
			Value:    3,
			LeftNode: nil,
			RightNode: &Node{
				Value:     6,
				LeftNode:  nil,
				RightNode: nil,
			},
		},
		RightNode: &Node{
			Value: 2,
			LeftNode: &Node{
				Value:     5,
				LeftNode:  nil,
				RightNode: nil,
			},
			RightNode: &Node{
				Value:     4,
				LeftNode:  nil,
				RightNode: nil,
			},
		},
	}

	mirror := mirrorTree(myTree)

	assert.True(t, isEqual(reversedTree, mirror))
}
