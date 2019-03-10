package tree

// 1. tree to array
func isSameTree(p *TreeNode, q *TreeNode) bool {
	pList := travPre(p).([]interface{})
	qList := travPre(q).([]interface{})

	pl := len(pList)
	ql := len(qList)

	    if pl != ql {
	        return false
	    }else {
	        for i := 0; i < pl; i++ {
	            if pList[i] != qList[i] {
	                return false
	            }
	        }
	    }

	    return true
}

// 2. recursion tree
func isSameTree2(p *TreeNode, q *TreeNode) bool{

	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree2(p.Left, q.Left) && isSameTree2(p.Right, q.Right)
	}

	return false
}
