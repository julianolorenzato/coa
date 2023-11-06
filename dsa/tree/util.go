package tree

func CompareSubTrees(root1, root2 *node) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 != nil && root2 != nil {
		if root1.Val == root2.Val {
			return true
		}
	}

	return false
}
