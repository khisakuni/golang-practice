package selectionSort

// SelectionSort implements selection sort O(n^2)
func SelectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		unsorted := nums[i:]
		smallestIndex := min(unsorted)
		swap(smallestIndex+i, i, nums)
	}
	return nums
}

func swap(i1 int, i2 int, nums []int) {
	nums[i1], nums[i2] = nums[i2], nums[i1]
}

func min(nums []int) int {
	smallestIndex := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[smallestIndex] {
			smallestIndex = i
		}
	}

	return smallestIndex
}
