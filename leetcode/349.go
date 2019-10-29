package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	m1 := make(map[int]bool)
	for _, v := range nums1 {
		m1[v] = true
	}

	for _, v := range nums2 {
		if m1[v] {
			result = append(result, v)
			delete(m1, v)
		}
	}
	return result
}
