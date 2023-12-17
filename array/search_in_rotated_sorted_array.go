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
	if nums[0] == target {
		return 0
	}

	// Find the lowest number
	left := 0
	right := len(nums) - 1

	/*
		L < M < R: 1 2 3
		L < M > R: 1 3 2
		L > M < R: 3 1 2
		L > M > R: 3 2 1 (NOT POSSIBLE)
	*/

	index := -1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] == target {
			index = mid
			break
		}

		if nums[left] == target {
			return left
		}

		if nums[right] == target {
			return right
		}

		/*
			1 2 3 4 5	-L < R
			5 1 2 3 4	=L > R	-L > M
			4 5 1 2 3	=L > R	-L > M
			3 4 5 1 2	=L > R	=L < M
			2 3 4 5 1	=L > R	=L < M
		*/
		if nums[left] <= nums[mid] && nums[mid] < nums[right] {
			// target in left-mid
			if target <= nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			/*
				5 1 2 3 4	M < R
				4 5 1 2 3	M < R
			*/
			if nums[mid] < nums[right] {
				if target >= nums[mid] && target <= nums[right] {
					left = mid + 1
				} else {
					right = mid
				}
			} else {
				/*
					3 4 5 1 2	M > R
					2 3 4 5 1	M > R
				*/
				if target >= nums[left] && target <= nums[mid] {
					right = mid
				} else {
					left = mid + 1
				}
			}
		}
	}
	return index
}
