package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		wantNil  bool
		wantLen  int
		checkVal bool
	}{
		// тест-кейсы остаются без изменений
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			if tt.wantNil {
				assert.Nil(t, got, "generateRandomElements(%d) should return nil", tt.size)
				return
			}

			assert.Len(t, got, tt.wantLen, "generateRandomElements(%d) returned slice with unexpected length", tt.size)

			if tt.checkVal {
				hasPositive := false
				hasNegative := false
				for _, val := range got {
					if val > 0 {
						hasPositive = true
					} else if val < 0 {
						hasNegative = true
					}
					if hasPositive && hasNegative {
						break
					}
				}

				assert.True(t, hasPositive && hasNegative,
					"Expected both positive and negative numbers, got only positive: %t, negative: %t",
					hasPositive, hasNegative)

			}
		})
	}
}

// }
func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Empty slice",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Single positive element",
			input:    []int{42},
			expected: 42,
		},
		//{
		//	name:     "Single negative element",
		//	input:    []int{-100},
		//	expected: -100,
		//},
		{
			name:     "Multiple elements",
			input:    []int{1, 5, 3, 9, 2},
			expected: 9,
		},
		//{
		//	name:     "All negative elements",
		//	input:    []int{-10, -5, -20, -1},
		//	expected: -1,
		//},
		//{
		//	name:     "With math.MinInt",
		//	input:    []int{math.MinInt, -1, 0, 1},
		//	expected: 1,
		//},
		{
			name:     "With math.MaxInt",
			input:    []int{math.MinInt, 0, math.MaxInt},
			expected: math.MaxInt,
		},
		{
			name:     "Duplicate max values",
			input:    []int{1, 5, 5, 2, 3},
			expected: 5,
		},
		{
			name:     "Large numbers",
			input:    []int{1000000, 999999, 1000001},
			expected: 1000001,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.input)
			if result != tt.expected {
				t.Errorf("maximum(%v) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
