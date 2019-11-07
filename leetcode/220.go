package leetcode
// 1.
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

	for warp := 1; warp <= k; warp++ {

		for i := 0; i+warp < len(nums); i++ {
			if abs(nums[i] , nums[i+warp])<=t {
				return true
			}
		}
	}
	return false
}


func abs(a,b int)int{
	if a>b{
		return a-b
	}
	return b-a
}




//2 .