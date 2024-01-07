package dollar

import (
	"reflect"
	"testing"
)

func TestMultipleDollar(t *testing.T) {
	tests := []struct {
		description  string
		firstDollar  *Dollar
		secondDollar *Dollar
		expected     *Dollar
		expectedErr  error
	}{
		{
			description:  "Dollar 3 * 5",
			firstDollar:  &Dollar{Price: 3},
			secondDollar: &Dollar{Price: 5},
			expected:     &Dollar{Price: 15},
			expectedErr:  nil,
		},
		{
			description:  "Dollar 7 * 2",
			firstDollar:  &Dollar{Price: 7},
			secondDollar: &Dollar{Price: 2},
			expected:     &Dollar{Price: 14},
			expectedErr:  nil,
		},
		{
			description:  "Dollar 1 * 3",
			firstDollar:  &Dollar{Price: 1},
			secondDollar: &Dollar{Price: 3},
			expected:     &Dollar{Price: 3},
			expectedErr:  nil,
		},
		{
			description:  "Dollar 0 * 100",
			firstDollar:  &Dollar{Price: 0},
			secondDollar: &Dollar{Price: 100},
			expected:     &Dollar{Price: 0},
			expectedErr:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := test.firstDollar.Multiple(test.secondDollar)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}

func TestIsEqualPrice(t *testing.T) {
	tests := []struct {
		description  string
		firstDollar  *Dollar
		secondDollar *Dollar
		expected     bool
		expectedErr  error
	}{
		{
			description:  "Dollar 3 == 3",
			firstDollar:  &Dollar{Price: 3},
			secondDollar: &Dollar{Price: 3},
			expected:     true,
			expectedErr:  nil,
		},
		{
			description:  "Dollar 7 != 2",
			firstDollar:  &Dollar{Price: 7},
			secondDollar: &Dollar{Price: 2},
			expected:     false,
			expectedErr:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := test.firstDollar.IsEqualPrice(test.secondDollar)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
