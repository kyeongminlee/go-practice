package money

import (
	"reflect"
	"testing"
)

func TestMultipleMoney(t *testing.T) {
	tests := []struct {
		description string
		firstMoney  *Money
		secondMoney *Money
		expected    *Money
		expectedErr error
	}{
		{
			description: "Money 3 * 5",
			firstMoney:  Dollar(5),
			secondMoney: Dollar(3),
			expected:    &Money{Price: 15, Currency: DollarString},
			expectedErr: nil,
		},
		{
			description: "Money 7 * 2",
			firstMoney:  Dollar(7),
			secondMoney: Dollar(2),
			expected:    &Money{Price: 14, Currency: DollarString},
			expectedErr: nil,
		},
		{
			description: "Money 1 * 3",
			firstMoney:  Franc(1),
			secondMoney: Franc(3),
			expected:    &Money{Price: 3, Currency: FrancString},
			expectedErr: nil,
		},
		{
			description: "Money 0 * 100",
			firstMoney:  Dollar(3),
			secondMoney: Franc(4),
			expected:    nil,
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := test.firstMoney.Multiple(*test.secondMoney)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}

func TestIsEqualPrice(t *testing.T) {
	tests := []struct {
		description string
		firstMoney  *Money
		secondMoney *Money
		expected    bool
		expectedErr error
	}{
		{
			description: "Money 3 == 3",
			firstMoney:  Dollar(3),
			secondMoney: Dollar(3),
			expected:    true,
			expectedErr: nil,
		},
		{
			description: "Money 3 != 4",
			firstMoney:  Dollar(3),
			secondMoney: Dollar(4),
			expected:    false,
			expectedErr: nil,
		},
		{
			description: "Money 3 == 3",
			firstMoney:  Franc(3),
			secondMoney: Franc(3),
			expected:    true,
			expectedErr: nil,
		},
		{
			description: "Money 3 != 4",
			firstMoney:  Franc(3),
			secondMoney: Franc(4),
			expected:    false,
			expectedErr: nil,
		},
		{
			description: "Money 3 != Franc 3",
			firstMoney:  Dollar(3),
			secondMoney: Franc(3),
			expected:    false,
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := test.firstMoney.IsEqualPrice(*test.secondMoney)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
