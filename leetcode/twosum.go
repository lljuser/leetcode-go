package leetcode

func TwoSum(nums []int,target int) []int {
	var arr []int
	for i:=0;i< len(nums);i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				arr = []int{i, j}
				break
			}
		}
	}

	return arr
}


