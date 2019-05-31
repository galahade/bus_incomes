package util

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/golang/glog"
)

// GetMonthlyStandardTime generate date by year and month.
func GetMonthlyStandardTime(year int, month time.Month) time.Time {
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
}

//CheckErr is used to checke err
func CheckErr(err error, message ...string) {
	if err != nil {
		if message != nil {
			log.Fatalf("%s: %s/n", message, err)
		} else {
			log.Fatal(err)
		}
	}
}

//WrapErrIfNotOK is used to checke err
func WrapErrIfNotOK(err error, message ...string) (ok bool, result error) {
	ok = true
	if err != nil {
		ok = false
		if message != nil {
			err = fmt.Errorf("%s: %s", message, err)
		}
	}
	return ok, err
}

//TestCheckErr is used to check test errors
func TestCheckErr(err error, t *testing.T, message ...string) {
	if err != nil {
		if message != nil {
			t.Fatalf("%s: %s", message, err)
		} else {
			t.Fatal(err)
		}
	}
}

//ConvertYearAndMonth util
func ConvertYearAndMonth(yearStr string, monthStr string) (ok bool, year int, month time.Month) {
	methodName := "ConvertYearAndMonth"
	ok = false
	var err error
	if year, err = strconv.Atoi(yearStr); err == nil {
		if monthI, err := strconv.Atoi(monthStr); err == nil {
			ok = true
			month = time.Month(monthI)
		} else {
			glog.Errorf("%s method convert month error: %s", methodName, err)
		}
	} else {
		glog.Errorf("%s method convert year error: %s", methodName, err)
	}
	return ok, year, month
}

//ConvertYearMonthOrLine util
func ConvertYearMonthOrLine(yearStr string, monthStr string, lineStr ...string) (ok bool, year int, month time.Month, line int) {
	methodName := "ConvertYearMonthOrLine"
	ok = false
	var err error
	if year, err = strconv.Atoi(yearStr); err == nil {
		if monthI, err := strconv.Atoi(monthStr); err == nil {
			month = time.Month(monthI)
			if lineStr == nil {
				ok = true
			} else {
				if line, err = strconv.Atoi(lineStr[0]); err == nil {
					ok = true
				}
			}
		} else {
			glog.Errorf("%s method convert month error: %s", methodName, err)
		}
	} else {
		glog.Errorf("%s method convert year error: %s", methodName, err)
	}
	return ok, year, month, line
}
