package handler

import (
	"net/http"

	"github.com/galahade/bus_incomes/domain"
	"github.com/galahade/bus_incomes/service"
	s "github.com/galahade/bus_incomes/service"
	"github.com/galahade/bus_incomes/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// DeleteDepartment used to delete a deprtment
func DeleteDepartment(c *gin.Context) {
	methodName := "DeleteDepartment"
	var ok bool
	var err error
	status := http.StatusBadRequest
	dep := new(domain.Department)

	err = c.ShouldBindJSON(dep)
	ok, err = util.WrapErrIfNotOK(err, "bind json to domain Department err")

	if err != nil {
		glog.Errorf("%s method has an error, which is : %s\n", methodName, err)
		c.JSON(status, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	} else {
		if ok {
			glog.Infof("Request is: %#v", dep)
			ok = service.RemoveDepartment(*dep)
			if ok {
				status = http.StatusOK
			}
		}
		c.JSON(status, gin.H{
			"success": ok,
		})
	}

}

// AddDepartment used to add a line monthly income
func AddDepartment(c *gin.Context) {
	methodName := "AddDepartment"
	var status int
	var err error
	var ok bool

	dep := new(domain.Department)

	err = c.ShouldBindJSON(dep)
	ok, err = util.WrapErrIfNotOK(err, "bind json to domain Department err")
	if ok {
		glog.Infof("Request is: %#v", dep)
		ok, err = service.AddDepartment(dep)
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
		glog.Errorf("%s method has an error, which is : %s\n", methodName, err)
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

// GetAllDepartments get all lines monthly incomes
func GetAllDepartments(c *gin.Context) {
	var results []domain.Department
	ok := false
	status := http.StatusBadRequest
	if ok, results = s.GetAllDepartments(); ok {
		status = http.StatusOK
		c.JSON(status, gin.H{
			"department": results,
		})
	} else {
		c.JSON(status, gin.H{})
	}
}
