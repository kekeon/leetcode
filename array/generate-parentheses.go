package array

func _generateParenthesis(left, right, n int, s string, r *[]string) *[]string {

	// 终止条件
	if left == n && right == n {
		*r = append(*r, s)
		return r
	}

	if left < n {
		_generateParenthesis(left+1, right, n, s+"(", r )
	}

	if right < left {
		_generateParenthesis(left, right+1, n, s+")", r)
	}

	return r
}

func generateParenthesis(n int) []string {

	v := []string{}

	v1 := _generateParenthesis(0, 0, 3, "", &v)
	return *v1
}
