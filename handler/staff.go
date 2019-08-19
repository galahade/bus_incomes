package handler

import (
	"net/http"

	"github.com/galahade/bus_incomes/domain"
	"github.com/galahade/bus_incomes/model"
	"github.com/galahade/bus_incomes/service"
	"github.com/galahade/bus_incomes/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// DeleteStaff used to delete a deprtment
func DeleteStaff(c *gin.Context) {
	methodName := "DeleteStaff"
	var ok bool
	var err error
	status := http.StatusBadRequest
	staff := new(domain.Staff)

	err = c.ShouldBindJSON(staff)
	ok, err = util.WrapErrIfNotOK(err, "bind json to domain Staff err")

	if err != nil {
		glog.Errorf("%s method has an error, which is : %s\n", methodName, err)
		c.JSON(status, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	} else {
		glog.Infof("Request is: %#v", staff)
		ok = service.RemoveStaff(*staff)
		if ok {
			status = http.StatusOK
		}
		c.JSON(status, gin.H{
			"success": ok,
		})
	}

}

// AddStaff used to add a line monthly income
func AddStaff(c *gin.Context) {
	methodName := "AddStaff"
	status := http.StatusBadRequest
	var err error
	var ok bool

	staffModel := new(model.StaffModel)
	err = c.ShouldBindJSON(staffModel)
	ok, err = util.WrapErrIfNotOK(err, "bind json to domain Staff err")

	if err == nil {
		glog.Infof("Request is: %#v", staffModel)

		ok, err = service.AddStaff(staffModel)
		if ok || err == nil {
			if ok {
				status = http.StatusOK
			} else {
				status = http.StatusConflict
			}

			c.JSON(status, gin.H{
				"success": ok,
			})
		} else {
			c.JSON(status, gin.H{
				"success": ok,
				"error":   err,
			})
		}

		return
	}

	glog.Errorf("%s method has an error, which is : %s\n", methodName, err)
	c.JSON(status, gin.H{
		"success": false,
		"error":   err.Error(),
	})

}

// GetStaffByDepartment get all staff
func GetStaffByDepartment(c *gin.Context) {
	departmentID := c.Param("departmentID")

	var results []model.StaffModel
	ok := false
	status := http.StatusBadRequest
	if ok, results = service.GetStaffByDepartmentID(departmentID); ok {
		status = http.StatusOK
		c.JSON(status, gin.H{
			"staff": results,
		})
	} else {
		c.JSON(status, gin.H{})
	}
}

// GetAllStaff get all staff
func GetAllStaff(c *gin.Context) {
	var results []model.StaffModel
	ok := false
	status := http.StatusBadRequest
	if ok, results = service.GetAllStaff(); ok {
		status = http.StatusOK
		c.JSON(status, gin.H{
			"staff": results,
		})
	} else {
		c.JSON(status, gin.H{})
	}
}

// GetAllJobType get all staff
func GetAllJobType(c *gin.Context) {
	var results []string
	ok := false
	status := http.StatusBadRequest
	if ok, results = service.GetAllJobType(); ok {
		status = http.StatusOK
		c.JSON(status, gin.H{
			"jobTypes": results,
		})
	} else {
		c.JSON(status, gin.H{})
	}
}
