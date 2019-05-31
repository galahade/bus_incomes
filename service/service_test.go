package service

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTwoNextMonthIncomes(t *testing.T) {

	results := GetTwoNextMonthIncomes(time.Date(2019, time.April, 1, 0, 0, 0, 0, time.UTC))
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
}

func TestSortMonthLinesIncomes(t *testing.T) {

	results := GetTwoNextMonthIncomes(time.Date(2019, time.April, 1, 0, 0, 0, 0, time.UTC))
	sortedResults := SortMonthLinesIncomes(results)
	fmt.Printf(" Sorted result: %+v\n", sortedResults)
}
