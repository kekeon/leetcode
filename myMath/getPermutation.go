package myMath

import (
	"strconv"
	"strings"
	"fmt"
)

func getPermutationRecursion(n int, k int, depth int, countMap map[int]int, arr []string, res *string, num int) {
	num++
	fmt.Println("getPermutationRecursion", num)

	// 终止条件
	if k <= 1 {
		*res += strings.Join(arr, "")
		return
	}

	// 大于单位内阶乘
	if k >= countMap[depth-1] {

		v := k / countMap[depth-1]

		// 整除需要减一，非整除会向下取整
		if k%countMap[depth-1] == 0 {
			v = v - 1
		}
		*res += arr[v]
		arr = append(arr[0:v], arr[v+1:]...)
		getPermutationRecursion(n, k-(v*countMap[depth-1]), depth-1, countMap, arr, res, num)

	} else {
		// 查询 N 内第一个未使用的数
		*res += arr[0]
		arr = arr[1:]
		getPermutationRecursion(n, k, depth-1, countMap, arr, res, num)
	}
}

func getPermutation(n int, k int) string {

	count := 1
	arr := []string{}
	countMap := map[int]int{}
	for a := 1; a < n; a++ {
		count = count * a
		countMap[a] = count
	}

	for a := 1; a <= n; a++ {
		arr = append(arr, strconv.Itoa(a))
	}

	res := ""
	getPermutationRecursion(n, k, n, countMap, arr, &res, 0)
	return res
}
