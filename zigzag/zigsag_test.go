package zigzag

import "testing"

type TestCase struct {
	nums   []int
	answer int
}

var testCases = []TestCase{
	TestCase{nums: []int{1, 7, 4, 9, 2, 5}, answer: 6},
	TestCase{nums: []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, answer: 7},
	TestCase{nums: []int{44}, answer: 1},
	TestCase{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, answer: 2},
	TestCase{nums: []int{70, 55, 13, 2, 99, 2, 80, 80, 80, 80, 100, 19, 7, 5, 5, 5, 1000, 32, 32}, answer: 8},
	TestCase{
		nums: []int{
			374, 40, 854, 203, 203, 156, 362, 279, 812, 955,
			600, 947, 978, 46, 100, 953, 670, 862, 568, 188,
			67, 669, 810, 704, 52, 861, 49, 640, 370, 908,
			477, 245, 413, 109, 659, 401, 483, 308, 609, 120,
			249, 22, 176, 279, 23, 22, 617, 462, 459, 244,
		},
		answer: 36,
	},
}

func TestLongestZigZag(t *testing.T) {
	for _, testCase := range testCases {
		answer := LongestZigZag(testCase.nums)

		if answer != testCase.answer {
			t.Errorf("Expected %d, got %d\n", testCase.answer, answer)
		}
	}
}
