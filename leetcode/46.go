package leetcode

var res [][]int

func permute(nums []int) [][]int {
	res = make([][]int, 0)
	back(nums, nil)
	return res
}

func back(nums []int, track []int) {
	if len(nums) == len(track) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		hadDone := false
		for _, v := range track {
			if nums[i] == v {
				hadDone = true
				break
			}
		}
		if hadDone {
			continue
		}
		track = append(track, nums[i])
		back(nums, track)
		track = track[:len(track)-1]
	}
}
