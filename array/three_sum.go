package array

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
	var pos = make(map[int]int)
	var neg = make(map[int]int)
	var zer int

	for i := range nums {
		n := nums[i]
		if n > 0 {
			val, ok := pos[n]
			if !ok {
				pos[n] = 1
			} else {
				pos[n] = val + 1
			}
		} else if n < 0 {
			val, ok := neg[n]
			if !ok {
				neg[n] = 1
			} else {
				neg[n] = val + 1
			}
		} else {
			zer++
		}
	}

	var res [][]int
	zeroUsed := 0

	for i := 0; i < zer; i++ {
		// for every positive there should be a negative
		for posN, posNCount := range pos {
			if posNCount > 0 {
				negN := -1 * posN
				negNCount, negNExists := neg[negN]
				if negNExists && negNCount > 0 {
					res = append(res, []int{0, negN, posN})
					pos[posN]--
					neg[negN]--
					zeroUsed++
				}
			}
		}
	}

	remainingZeros := zer - zeroUsed
	if remainingZeros > 2 {
		numArrayWithAllZeros := remainingZeros / 3
		for numArrayWithAllZeros > 0 {
			res = append(res, []int{0, 0, 0})
			numArrayWithAllZeros--
		}
	}

	prevPos := -1
	for n, c := range pos {
		if c > 0 {
			if c == -1 {
				prevPos = n
				continue
			}
			sumOfTwoPos := n + prevPos
			requiredNeg := -1 * sumOfTwoPos
			negNCount, negNExists := neg[requiredNeg]
			if !negNExists || negNCount <= 0 {
				prevPos = n
				continue
			}
			res = append(res, []int{requiredNeg, n, prevPos})
			pos[n]--
			pos[prevPos]--
			neg[requiredNeg]--
			prevPos = n
		}
	}

	return res
}
