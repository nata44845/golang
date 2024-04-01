package main

import (
	"testing"
)

func TestTxOperator_Deposit(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		amounts  []float64
		expected float64
	}{
		{
			name:     "test_0",
			amounts:  []float64{0, 0, 0},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(
			tc.name, func(t *testing.T) {
				t.Parallel()
				op := &txOperator{}
				for _, amount := range tc.amounts {
					op.Deposit(amount)
				}

				balance := op.Balance()
				if tc.expected != balance {
					t.Errorf("got: %v, want: %v", balance)
				}
			},
		)
	}
}
