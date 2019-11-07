package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {

	m := make(map[int]int)

	for i, v := range nums {
		if m[v] != 0 && (i+1-m[v]) <= k {
			return true
		}
		m[v] = i + 1
	}

	return false
}

func containsNearbyDuplicate(nums []int, k int) bool {

	for warp := 1; warp <= k; warp++ {

		for i := 0; i+warp < len(nums); i++ {
			if nums[i] == nums[i+warp] {
				return true
			}
		}
	}
	return false
}
