package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"github.com/galahade/bus_incomes/domain"
	"github.com/galahade/bus_incomes/service"
	s "github.com/galahade/bus_incomes/service"
	"github.com/galahade/bus_incomes/util"
)

// AddLineMonthIncome used to add a line monthly income
func AddLineMonthIncome(c *gin.Context) {
	methodName := "AddLineMonthIncome"
	var status int
	var err error
	var ok bool

	income := new(domain.LineMonthIncome)

	err = c.ShouldBindJSON(income)
	ok, err = util.WrapErrIfNotOK(err, "bind json to LineMonthIncome err")
	if ok {
		log.Printf("Income request is: %#v", income)
		ok, err = service.AddLineMonthlyIncome(income)
		if ok {
			status = http.StatusOK
			//c.JSON(status, ok)
		} else {
			//// TODO: if has exist send different error code.
			status = http.StatusConflict
		}
	} else {
		status = http.StatusBadRequest
	}

	if err != nil {
		log.Printf("%s method has an error, which is : %s\n", methodName, err)
		c.JSON(status, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	} else {
		c.JSON(status, gin.H{
			"success": ok,
		})
	}

}

// DeleteLineMonthIncome used to add a line monthly income
func DeleteLineMonthIncome(c *gin.Context) {
	status := http.StatusOK
	var message string
	if ok, year, month, line := util.ConvertYearMonthOrLine(c.Param("year"), c.Param("month"), c.Param("line")); ok {
		if err := s.DeleteLineMonthlyIncome(year, month, line); err != nil {
			status = http.StatusInternalServerError
			message = err.Error()
		}
	} else {
		status = http.StatusBadRequest
	}

	c.JSON(status, gin.H{
		"message": message,
	})
}

// GetMonthLinesIncomes get all lines monthly incomes
func GetMonthLinesIncomes(c *gin.Context) {
	methodName := "GetMonthLinesIncomes"
	var results []domain.LineMonthIncome
	status := http.StatusBadRequest
	hasData := false
	var err error
	if ok, year, month := util.ConvertYearAndMonth(c.Param("year"), c.Param("month")); ok {
		if results, err = s.QueryMonthlyLinesIncomes(year, month); err == nil {
			status = http.StatusOK
			if len(results) > 0 {
				hasData = true
			}
		} else {
			status = http.StatusInternalServerError
		}
	}

	if status == http.StatusOK {
		c.JSON(status, gin.H{
			"hasData":    hasData,
			"incomeList": results,
		})
	} else {
		glog.Warningf("%s method fail. error is %s\n", methodName, err)
		c.JSON(http.StatusBadRequest, err)
	}
}

// GetRecentTwoMonthIncomes get all lines monthly incomes
func GetRecentTwoMonthIncomes(c *gin.Context) {
	incomes := s.QueryLastTowMonthLinesIncomes()
	status := http.StatusOK
	c.JSON(status, incomes)
}

// GetSortedTwoMonthComparedIncomesByDate get all lines monthly incomes by the passed year and month.
func GetSortedTwoMonthComparedIncomesByDate(c *gin.Context) {
	status := http.StatusBadRequest
	var results interface{}
	if ok, year, month := util.ConvertYearAndMonth(c.Param("year"), c.Param("month")); ok {
		time := util.GetMonthlyStandardTime(year, month)
		results = s.SortMonthLinesIncomes(s.GetTwoNextMonthIncomes(time))
		status = http.StatusOK
	}
	c.JSON(status, results)
}
