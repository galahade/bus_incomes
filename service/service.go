package service

import (
	"sort"
	"time"

	"github.com/galahade/bus_incomes/domain"
)

// SortMonthLinesIncomes used to sort
func SortMonthLinesIncomes(twoMonthsIncomes domain.TowMonthsIncomes) domain.SortedTwoMonthsIncomes {
	finalIncomeArray := twoMonthsIncomes.FinalIncomes.LineMonthIncomes
	previousIncomeArray := twoMonthsIncomes.PreviousIncomes.LineMonthIncomes
	var tmliArray []domain.TwoMonthsLineIncome
	for i := 0; i < len(finalIncomeArray); i++ {
		cmli := domain.TwoMonthsLineIncome{
			Line:           finalIncomeArray[i].Line,
			FinalIncome:    finalIncomeArray[i],
			PreviousIncome: previousIncomeArray[i],
		}
		tmliArray = append(tmliArray, cmli)
	}
	sort.Sort(domain.CMLISortByTotalIncome(tmliArray))
	for i := 0; i < len(tmliArray); i++ {
		tmliArray[i].TotalIncomeRank = i
		tmliArray[i].Score += tmliArray[i].TotalIncomeRank * 10
	}

	sort.Sort(domain.CMLISortByAverageIncome(tmliArray))
	for i := 0; i < len(tmliArray); i++ {
		tmliArray[i].CapitaIncomeRank = i
		tmliArray[i].Score += tmliArray[i].CapitaIncomeRank * 10
	}
	sort.Sort(domain.CMLISortByScore(tmliArray))
	result := domain.SortedTwoMonthsIncomes{
		SortedIncomesByScore: tmliArray,
		FinalMonth:           twoMonthsIncomes.FinalIncomes.Month,
		PreviousMonth:        twoMonthsIncomes.PreviousIncomes.Month,
	}
	return result
}

//QueryLastTowMonthLinesIncomes used to query current and last month incomes by lines.
func QueryLastTowMonthLinesIncomes() domain.TowMonthsIncomes {
	return GetTwoNextMonthIncomes(time.Now())

}

// GetTwoNextMonthIncomes query two next month all lines incomes
func GetTwoNextMonthIncomes(final time.Time) domain.TowMonthsIncomes {
	previous := final.AddDate(0, -1, 0)
	var result domain.TowMonthsIncomes
	result = domain.TowMonthsIncomes{}
	if finalMonthLinesIncomes, err1 := QueryMonthlyLinesIncomes(final.Year(), final.Month()); err1 == nil {
		result.FinalIncomes = domain.MonthIncomes{
			LineMonthIncomes: finalMonthLinesIncomes,
			Month:            final.Format("2006-01"),
		}
		if previousMonthLinesIncomes, err2 := QueryMonthlyLinesIncomes(previous.Year(), previous.Month()); err2 == nil {
			result.PreviousIncomes = domain.MonthIncomes{
				LineMonthIncomes: previousMonthLinesIncomes,
				Month:            previous.Format("2006-01"),
			}
		} else {
			result.Error = err2.Error()
		}
	} else {
		result.Error = err1.Error()
	}

	return result
}
