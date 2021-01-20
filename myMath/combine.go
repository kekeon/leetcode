package myMath

func combine(n, k int) [][]int {
	result := [][]int{}

	combineRecursion(1, n,k, []int{}, &result)

	return result
}

func combineRecursion(start, n, k int, path []int, result *[][]int) {

	if len(path) == k {
		*result = append(*result, append([]int{}, path...))
		return
	}

	// 剪枝/
	// 如果 n = 7, k = 4，从 5 开始搜索就已经没有意义了，这是因为：即使把 5 选上，后面的数只有 6 和 77，一共就 3 个候选数，凑不出 4 个数的组合。因此，搜索起点有上界，这个上界是多少，可以举几个例子分析
	for a := start; a <= n - k + len(path) + 1; a++ {

		// 基础版本
	// for a := start; a <= n; a++ {
		path = append(path, a)
		combineRecursion(a+1, n, k, path, result)
		path = path[:len(path) - 1]
	}
}
