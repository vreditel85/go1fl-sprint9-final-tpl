package main

import (
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
		{
			name:    "Negative size",
			size:    -1,
			wantNil: true,
		},
		{
			name:    "Zero size",
			size:    0,
			wantNil: true,
		},
		{
			name:     "Small size",
			size:     5,
			wantLen:  5,
			checkVal: true,
		},
		{
			name:     "Large size",
			size:     1000,
			wantLen:  1000,
			checkVal: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			if tt.wantNil {
				if got != nil {
					t.Errorf("generateRandomElements(%d) = %v, want nil", tt.size, got)
				}
				return
			}

			if len(got) != tt.wantLen {
				t.Errorf("generateRandomElements(%d) returned slice with length %d, want %d",
					tt.size, len(got), tt.wantLen)
			}

			if tt.checkVal {
				// Проверяем, что все элементы в слайсе
				for i, val := range got {
					if i > 0 && val == got[i-1] {
						t.Logf("Duplicate value found at index %d", i)
					}
				}
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
				if !hasPositive || !hasNegative {
					t.Errorf("Expected both positive and negative numbers, got only %s",
						map[bool]string{true: "positive", false: "negative"}[hasPositive])
				}
			}
		})
	}
}
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
		{
			name:     "Single negative element",
			input:    []int{-100},
			expected: -100,
		},
		{
			name:     "Multiple elements",
			input:    []int{1, 5, 3, 9, 2},
			expected: 9,
		},
		{
			name:     "All negative elements",
			input:    []int{-10, -5, -20, -1},
			expected: -1,
		},
		{
			name:     "With math.MinInt",
			input:    []int{math.MinInt, -1, 0, 1},
			expected: 1,
		},
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
