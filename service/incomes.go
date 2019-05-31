package service

import (
	"fmt"
	"time"

	"github.com/galahade/bus_incomes/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddLineMonthlyIncome used to add a Line's monthly income to DB
func AddLineMonthlyIncome(income *domain.LineMonthIncome) (ok bool, err error) {
	ok = false
	isExist, _ := IsLineMonthIncomeExist(*income)
	if isExist {
		err = fmt.Errorf("Line %d's %s income already exist", income.Line.No, income.IncomeMonth.Format("2006-01"))
	} else {
		if err = domain.InsertLineMonthIncome(income); err == nil {
			ok = true
		}
	}
	return ok, err
}

//DeleteLineMonthlyIncome used to add a Line's monthly income to DB
func DeleteLineMonthlyIncome(year int, month time.Month, line int) (err error) {
	return domain.DeleteLineMonthIncome(year, month, line)
}

//QueryLineMonthIncome get LineMonthIncome from DB
func QueryLineMonthIncome(year int, month time.Month, lineNO int) (bool, domain.LineMonthIncome) {
	return domain.SelectLineMonthIncome(year, month, lineNO)
}

//IsLineMonthIncomeExist used to check if a line's monthly income has in DB
func IsLineMonthIncomeExist(income domain.LineMonthIncome) (bool, primitive.ObjectID) {
	ok, result := QueryLineMonthIncome(income.IncomeMonth.Year(), income.IncomeMonth.Month(), income.Line.No)
	return ok, result.ID
}

// QueryMonthlyLinesIncomes get one month all lines incomes by line no. asc
func QueryMonthlyLinesIncomes(year int, month time.Month) ([]domain.LineMonthIncome, error) {
	if ok, results := domain.FindMonthLinesIncomesOrderByLineNo(year, month); ok {
		return results, nil
	}
	return nil, fmt.Errorf("fail to Get lines incomes by month - Server error")

}
