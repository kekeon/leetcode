package array

import "sort"

// 暴力解法
func majorityElement(nums []int) int {
	return majorityElementRecursion(nums, 0, len(nums) - 1)
}

// map 解法
func majorityElementMap(nums []int) int {

	m := map[int]int{}

	l := float32(len(nums)) / 2
	for _, v := range nums {
		m[v]++
		if float32(m[v]) > l {
			return v
		}
	}

	return 0

}

// 排序解法。 如果定存在众数，则1/2 处定为众数

func majorityElementSort(nums []int) int {

	sort.Ints(nums)

	return nums[len(nums)/2]
}

// 分治解法
func majorityElementDivide(nums []int) int {

	sort.Ints(nums)

	return nums[len(nums)/2]
}

func majorityElementRecursionCount(nums []int, num, start, end int) int {
	count := 0
	for  := start;  <= end; ++ {
		if nums[a] == num {
			count++
		}
	}

	return count
}

func majorityElementRecursion(nums []int, start, end int) int {

	// 终止条件
	if start == end {
		return nums[start]
	}

	// 做加法防止越界
	mid := (end-start)/2 + start

	// 分治
	left := majorityElementRecursion(nums, start, mid)
	right := majorityElementRecursion(nums, mid + 1, end)

	if left == right {
		return left
	}

	// 统计分治结果
	leftCount := majorityElementRecursionCount(nums, left, start, end)
	rightCount := majorityElementRecursionCount(nums, right, start, end)

	// 比较结果
	if leftCount > rightCount {
		return left
	}

	return right
}
