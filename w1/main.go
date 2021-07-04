package main

func main() {
	// [1,1,1]
	// 2
	testData := []int{1, 2, 1, 1, 3}
	target := 3
	sum := subarraySum(testData, target)
	if sum == 2 {
		// test pass
	}
}

// 和为 K 的子数组（Medium）
func subarraySum(nums []int, k int) int {
	hashmap, pre, res := map[int]int{0: 1}, 0, 0
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		// 如果没有的话这里是0，不影响结果
		res += hashmap[pre-k]
		hashmap[pre]++
	}
	return res
}

// O(n^2)
func subarraySum(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end >= 0; end-- {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

//func subarraySum2(nums []int, k int) int {
//	hashmap, pre, res := map[int]int{0: 1}, 0, 0
//
//	for i := 0; i < len(nums); i++ {
//		pre += nums[i]
//		if hashmap[pre-k] > 0 {
//			res += hashmap[pre-k]
//		}
//		hashmap[pre]++
//	}
//	return res
//}
