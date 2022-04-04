package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCalculatePathEnergy(t *testing.T) {
	testCases := []struct {
		inParameters *Input
		expected     []int
	}{
		{
			inParameters: &Input{
				Shortcuts: []int{2, 2, 3},
			},
			expected: []int{0, 1, 2},
		},
		{
			inParameters: &Input{
				Shortcuts: []int{1, 2, 3, 4, 5},
			},
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			inParameters: &Input{
				Shortcuts: []int{4, 4, 4, 4, 7, 7, 7},
			},
			expected: []int{0, 1, 2, 1, 2, 3, 3},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", i), func(t *testing.T) {
			actual := CalculatePathEnergy(tc.inParameters)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("expected result: %v, got: %v", tc.expected, actual)
				return
			}
		})
	}
}
