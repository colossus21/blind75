package array

//Given an integer array nums, find a
//subarray
//that has the largest product, and return the product.
//
//The test cases are generated so that the answer will fit in a 32-bit integer.
//
//
//
//Example 1:
//
//Input: nums = [2,3,-2,4]
//Output: 6
//Explanation: [2,3] has the largest product 6.
//Example 2:
//
//Input: nums = [-2,0,-1]
//Output: 0
//Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

func maxProduct(nums []int) int {
	maximumProduct := nums[0]
	currentProductPositive := 1
	currentProductNegative := 1

	for i := range nums {
		if nums[i] > 0 {
			currentProductPositive, currentProductNegative = currentProductPositive*nums[i], currentProductNegative*nums[i]
		} else if nums[i] < 0 {
			currentProductPositive, currentProductNegative = currentProductNegative*nums[i], currentProductPositive*nums[i]
		} else {
			currentProductPositive = 0
			currentProductNegative = 1
		}
		if currentProductPositive > maximumProduct {
			maximumProduct = currentProductPositive
		}
		if currentProductPositive <= 0 {
			currentProductPositive = 1
		}
	}

	return maximumProduct
}
