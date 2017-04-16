package selectionSort

import "testing"

func TestSelectionSort(t *testing.T) {
	nums := []int{5, 3, 6, 8, 1, 3, 10, 3, 15}
	expected := []int{1, 3, 3, 3, 5, 6, 8, 10, 15}
	answer := SelectionSort(nums)

	for i, v := range expected {
		if len(answer) != len(expected) {
			t.Errorf("Expected %v, got %v\n", expected, answer)
		}

		if v != answer[i] {
			t.Errorf("Expected %v, got %v\n", expected, answer)
		}
	}
}
