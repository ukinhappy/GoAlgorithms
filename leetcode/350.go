package leetcode

func intersect(nums1 []int, nums2 []int) []int {

	var result []int
	tmp := make(map[int]int)

	for _, v := range nums1 {
		tmp[v] = tmp[v] + 1
	}
	for _, v := range nums2 {
		if tmp[v] > 0 {
			tmp[v] = tmp[v] - 1
			result = append(result, v)
		}
	}
	return result
}
