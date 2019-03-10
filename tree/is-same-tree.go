package tree

import "fmt"

func isSameTree(p *TreeNode, q *TreeNode) bool {
	pList := travPre(p).([]interface{})
	qList := travPre(q).([]interface{})

	fmt.Println(pList)
	fmt.Println(qList)

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
