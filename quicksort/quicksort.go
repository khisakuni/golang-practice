package quicksort

func Quicksort(nums []int, left int, right int) []int {
	index := partition(nums, left, right)
	if left < index-1 {
		Quicksort(nums, left, index-1)
	}
	if index < right {
		Quicksort(nums, index, right)
	}
	return nums
}

func partition(nums []int, left int, right int) int {
	pivot := nums[(left+right)/2]
	for left <= right {
		for nums[left] < pivot {
			left++
		}

		for nums[right] > pivot {
			right--
		}

		if left <= right {
			swap(nums, left, right)
			left++
			right--
		}
	}
	return left
}

func swap(nums []int, i1 int, i2 int) {
	nums[i1], nums[i2] = nums[i2], nums[i1]
}
