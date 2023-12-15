package array

import "fmt"

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
	right := len(nums) - 1

	/*
		L < M < R: 1 2 3
		L < M > R: 1 3 2
		L > M < R: 3 1 2
		L > M > R: 3 2 1 (NOT POSSIBLE)
	*/

	index := -1
	fmt.Println(nums)
	for left < right {
		mid := (left + right) / 2

		if nums[mid] == target {
			index = mid
			break
		}
		fmt.Println("A L=", left, "R=", right, "M=", mid, " [ L , R, M ] =", "[", nums[left], ",", nums[mid], ",", nums[right], "]")
		/*
			1 2 3 4 5	-L < R
			5 1 2 3 4	=L > R	-L > M
			4 5 1 2 3	=L > R	-L > M
			3 4 5 1 2	=L > R	=L < M
			2 3 4 5 1	=L > R	=L < M
		*/

		perfectlySorted := nums[left] <= nums[right]
		if perfectlySorted {
			fmt.Println("SORTED LOGIC")
			if target <= nums[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		} else {
			if nums[left] > nums[mid] {
				if target <= nums[mid] {
					fmt.Println("L > M && T <= M :: R = M")
					right = mid
				} else {
					fmt.Println("L > M && T > M :: L = M")
					left = mid + 1
				}
			} else {
				if target <= nums[mid] {
					fmt.Println("L <= M && T <= M :: R = M")
					right = mid
				} else {
					fmt.Println("L <= M && T <= M :: L = M")
					left = mid + 1
				}
			}
		}
		fmt.Println("B L=", left, "R=", right, "M=", mid)
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target {
		return right
	}
	return index
}
