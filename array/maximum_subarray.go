package array

//Given an integer array nums, find the
//subarray
//with the largest sum, and return its sum.
//
//
//
//Example 1:
//
//Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
//Output: 6
//Explanation: The subarray [4,-1,2,1] has the largest sum 6.
//Example 2:
//
//Input: nums = [1]
//Output: 1
//Explanation: The subarray [1] has the largest sum 1.
//Example 3:
//
//Input: nums = [5,4,-1,7,8]
//Output: 23
//Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.

func MaxSubArray(nums []int) int {
	// kadane's algorithm
	maxSum := nums[0]
	currentSum := 0

	for i := range nums {
		currentSum += nums[i]

		if currentSum < 0 {
			currentSum = 0
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	// one issue, kadane won't work for all negative numbers in an array
	// to resolve this, return the highest negative number
	// this is not optimized intentionally for better representation of the issue

	allNegatives := true
	maxValue := nums[0]

	for i := range nums {
		if nums[i] >= 0 {
			allNegatives = false
			break
		}
		if nums[i] > maxValue {
			maxValue = nums[i]
		}
	}

	if allNegatives {
		return maxValue
	}

	return maxSum
}
