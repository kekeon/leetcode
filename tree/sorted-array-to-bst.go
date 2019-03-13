package tree

// 108. 将有序数组转换为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode1 {

	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	return &TreeNode1{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}

}
