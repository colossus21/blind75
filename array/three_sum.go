package array

import "fmt"

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
*/

func ThreeSum(nums []int) [][]int {
	usedIndices := make(map[string]bool)
	var res [][]int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			for k := 0; k < len(nums); k++ {
				if i != j && i != k && j != k {
					if nums[i]+nums[j]+nums[k] == 0 {
						hash := h(nums[i], nums[j], nums[k])
						if !usedIndices[hash] {
							res = append(res, []int{nums[i], nums[j], nums[k]})
							usedIndices[hash] = true
						}
					}
				}
			}
		}
	}
	return res
}

func h(i, j, k int) string {
	if i < j {
		i, j = j, i
	}
	if j < k {
		j, k = k, j
	}
	if i < j {
		i, j = j, i
	}
	return fmt.Sprintf(`%v-%v-%v`, i, j, k)
}
