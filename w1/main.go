package main

func main() {
	// [1,1,1]
	// 2
	testData := []int{1, 1, 1, 2}
	target := 2
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
