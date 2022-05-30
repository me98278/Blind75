package two_sum

func TwoSum(nums []int, target int) []int {
	// take the current number and sustract from the target
	// if the difference exists in the nums array then we are done
	// if diff does not exist in the nums array,
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		x := nums[i]
		y := target - x
		
	}
}
