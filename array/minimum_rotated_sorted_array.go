package array

//Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:
//
//[4,5,6,7,0,1,2] if it was rotated 4 times.
//[0,1,2,4,5,6,7] if it was rotated 7 times.
//Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
//
//Given the sorted rotated array nums of unique elements, return the minimum element of this array.
//
//You must write an algorithm that runs in O(log n) time.
//
//
//
//Example 1:
//
//Input: nums = [3,4,5,1,2]
//Output: 1
//Explanation: The original array was [1,2,3,4,5] rotated 3 times.
//Example 2:
//
//Input: nums = [4,5,6,7,0,1,2]
//Output: 0
//Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.
//Example 3:
//
//Input: nums = [11,13,15,17]
//Output: 11
//Explanation: The original array was [11,13,15,17] and it was rotated 4 times.

/*
Analysis:
[1, 2, 3, 4 , 5] l = 1, r = 5, m = 3 [l < m < r]
[5, 1, 2, 3, 4] l = 5, r = 4, m = 2 [l > r > m]
[3, 4, 5, 1, 2] l = 3, r = 2, m = 5 [ m > l > r]
[2, 3, 4, 5, 1]
*/

func FindMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	lowest := -1

	for {
		mid := (right + left) / 2

		if nums[mid] <= nums[left] && nums[mid] <= nums[right] {
			lowest = nums[mid]
			break
		}

		/*
			[3,4,1,2]
			If left elem is greater than the right element,
			lower element is on the right sight, else on the left side

			0 + 4 / 2 = 2, arr[2] = 1
			arr[2] < left && arr[2] < right [Solved]

			[3,4,5,6,7,0,1,2]
			0 + 7 /2 = 3, arr[3] = 6
			arr[3] > left [Go right], l = 3, r = 7
			3 + 7 = 10 / 2 = 5, arr[5] = 0
			arr[5] < left && arr[5] < right [Solved]

			[6,0,1,2,3,4,5]
			0+6 / 2 = 3, arr[3] = 2
			arr[3] < left [Go left], l = 0, r = 2
			2 + 6 / 2 = 4, arr[4] = 3
		*/

		if nums[right] > nums[left] {
			right = mid
			left = 0
		} else {
			left = mid
			right = len(nums) - 1
		}

	}

	return lowest
}
