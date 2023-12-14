package array

/*
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right)/2

		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return nums[left]
*/

func Search(nums []int, target int) int {
	// Find the lowest number
	left := 0
	right := len(nums)

	/*
		1 2 3 4 5
		1..3..5 M < R
		1 2 3 4 5
		1..3..5 M > R
	*/

	index := -1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] == target {
			index = mid
		}

		// Lower side on left
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return index
}
