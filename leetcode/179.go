package leetcode

import "strconv"

func largestNumber(nums []int) string {
	if len(nums) <= 0 {
		return ""
	}
	if len(nums) <= 1 {
		return strconv.Itoa(nums[0])
	}
	var numslice []string

	for _, num := range nums {
		numslice = append(numslice, strconv.Itoa(num))
	}
	var flag int = len(nums) - 1
	for flag > 0 {
		var tmp int
		for j := 0; j < flag; j++ {
			if numslice[j]+numslice[j+1] < numslice[j+1]+numslice[j] {
				numslice[j], numslice[j+1] = numslice[j+1], numslice[j]
				tmp = j
			}
		}
		flag = tmp
	}
	var result string
	for _, v := range numslice {
		result += v
	}
	for i, v := range result {
		if v != '0' {
			return result[i:]
		}
	}
	return "0"
}
