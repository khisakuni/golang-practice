package radixsort

func RadixSort(nums []int) []int {
	maxNum := max(nums)
	m := 10
	n := 1

	for m < maxNum*10 {
		list := make([][]int, 10, 10)
		for _, v := range nums {
			index := (v % m) / n
			sublist := list[index]
			sublist = append(sublist, v)
			list[index] = sublist
		}
		nums = flatten(list)

		m = m * 10
		n = n * 10
	}

	return nums
}

func max(nums []int) int {
	max := -1
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func flatten(nums [][]int) []int {
	list := make([]int, 0)
	for _, v := range nums {
		list = append(list, v...)
	}
	return list
}
