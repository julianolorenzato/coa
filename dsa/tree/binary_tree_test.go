package tree

import (
	"reflect"
	"testing"
)

var exampleTree = &BinaryTree{
	Root: &node{
		Val: 5,
		Left: &node{
			Val:  9,
			Left: nil,
			Right: &node{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &node{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	},
}

func TestBinaryTree_PreOrder(t *testing.T) {
	got := make([]any, 0)
	want := []any{5, 9, 2, 1}

	exampleTree.PreOrder(func(n *node) {
		got = append(got, n.Val)
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("BinaryTree_PreOrder = %v, want %v", got, want)
	}
}

func TestBinaryTree_InOrder(t *testing.T) {
	got := make([]any, 0)
	want := []any{9, 2, 5, 1}

	exampleTree.InOrder(func(n *node) {
		got = append(got, n.Val)
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("BinaryTree_InOrder = %v, want %v", got, want)
	}
}

func TestBinaryTree_PostOrder(t *testing.T) {
	got := make([]any, 0)
	want := []any{2, 9, 1, 5}

	exampleTree.PostOrder(func(n *node) {
		got = append(got, n.Val)
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("BinaryTree_PostOrder = %v, want %v", got, want)
	}
}

func TestNewCompleteBinaryTree(t *testing.T) {
	want := BinaryTree{
		Root: &node{
			Val: 2,
			Left: &node{
				Val:   3,
				Left:  &node{Val: 2},
				Right: &node{Val: 5},
			},
			Right: &node{Val: 4},
		},
	}

	got := NewCompleteBinaryTree([]any{2, 3, 4, 2, 5})

	if !CompareSubTrees(want.Root, got.Root) {
		t.Errorf("NewCompleteBinaryTree want = %v, got = %v", want, got)
	}
}
