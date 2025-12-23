package twoSum

func twoSum(nums []int, target int) []int {
	hashm := make(map[int]int)
	res := []int{}
	for i := 0; i < len(nums); i++ {

		if j, exist := hashm[target-nums[i]]; exist {
			res = append(res, j, i)
			break
		} else {
			hashm[nums[i]] = i
		}

	}
	return res
}
