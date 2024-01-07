package money

import (
	"reflect"
	"testing"
)

func TestTimesMoney(t *testing.T) {
	tests := []struct {
		description string
		money       *Money
		multiplier  int
		expected    *Money
		expectedErr error
	}{
		{
			description: "Money 5 * 2",
			money:       Dollar(5),
			multiplier:  2,
			expected:    &Money{Price: 10, Currency: DollarString},
			expectedErr: nil,
		},
		{
			description: "Money 5 * 3",
			money:       Dollar(5),
			multiplier:  3,
			expected:    &Money{Price: 15, Currency: DollarString},
			expectedErr: nil,
		},
		{
			description: "Franc 5 * 2",
			money:       Franc(5),
			multiplier:  2,
			expected:    &Money{Price: 10, Currency: FrancString},
			expectedErr: nil,
		},
		{
			description: "Franc 5 * 3",
			money:       Franc(5),
			multiplier:  3,
			expected:    &Money{Price: 15, Currency: FrancString},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := test.money.Times(test.multiplier)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}

func TestAddMoney(t *testing.T) {
	tests := []struct {
		description string
		firstMoney  *Money
		secondMoney *Money
		expected    *Money
		expectedErr error
	}{
		{
			description: "Money 3 + 5",
			firstMoney:  Dollar(5),
			secondMoney: Dollar(3),
			expected:    &Money{Price: 8, Currency: DollarString},
			expectedErr: nil,
		},
		{
			description: "Money 7 + 2",
			firstMoney:  Dollar(7),
			secondMoney: Dollar(2),
			expected:    &Money{Price: 9, Currency: DollarString},
			expectedErr: nil,
		},
		{
			description: "Money 1 + 3",
			firstMoney:  Franc(1),
			secondMoney: Franc(3),
			expected:    &Money{Price: 4, Currency: FrancString},
			expectedErr: nil,
		},
		{
			description: "can't not add another currency",
			firstMoney:  Dollar(3),
			secondMoney: Franc(4),
			expected:    nil,
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := test.firstMoney.Add(*test.secondMoney)

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
