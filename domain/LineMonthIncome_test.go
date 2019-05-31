package domain

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestTwoMonthsLineIncomePrint(t *testing.T) {
	result := TwoMonthsLineIncome{}
	t.Logf("decimal value compare: %t", result.FinalTotalIncome.Equal(decimal.Zero))
	fmt.Printf("Empty TwoMonthsLineIncome struct is : %s", result)

}
