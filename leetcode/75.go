package leetcode

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i, j, k := 0, 0, len(nums)-1

	for j <= k {
		switch (nums[j]) {
		case 0:
			// nums[i] 是 [i:j]的第一个等于1的数字
			nums[i], nums[j] = nums[j], nums[i]
			// 交换之后,j++
			i++
			j++
		case 1:
			j++
		case 2:
			nums[j], nums[k] = nums[k], nums[j]
			k--
		}
	}
}
