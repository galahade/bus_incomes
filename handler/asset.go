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

// DeleteAsset used to delete a asset
func DeleteAsset(c *gin.Context) {
	methodName := "DeleteAsset"
	var ok bool
	var err error
	status := http.StatusBadRequest
	asset := new(domain.Asset)

	err = c.ShouldBindJSON(asset)
	ok, err = util.WrapErrIfNotOK(err, "bind json to domain asset err")

	if err != nil {
		glog.Errorf("%s method has an error, which is : %s\n", methodName, err)
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
	} else {
		glog.Infof("Request is: %#v", asset)
		ok = service.RemoveAsset(*asset)
		if ok {
			status = http.StatusOK
		}
		c.JSON(status, gin.H{
			"success": ok,
		})
	}

}

// AddAsset used to add a new asset
func AddAsset(c *gin.Context) {
	methodName := "AddAsset"
	status := http.StatusBadRequest
	var err error
	var ok bool

	asset := new(domain.Asset)
	err = c.ShouldBindJSON(asset)
	ok, err = util.WrapErrIfNotOK(err, "bind json to domain asset err")

	if err == nil {
		glog.Infof("Request is: %#v", asset)

		ok, err = service.AddAsset(asset)
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
				"error": err,
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

// TransferAsset used to add a new asset transfer record
func TransferAsset(c *gin.Context) {
	methodName := "TransferAsset"
	status := http.StatusBadRequest
	var err error
	var ok bool

	assetTransfer := new(domain.AssetTransfer)

	if ok, err = util.WrapErrIfNotOK(c.ShouldBindJSON(assetTransfer), "bind json to domain AssetTransfer err"); ok {
		glog.Infof("Request is: %#v", assetTransfer)
		if ok, err = service.AddAssetTransfer(assetTransfer); ok {
			if ok = service.ModifyAsset(*assetTransfer); ok {
				status = http.StatusOK
			}
			c.JSON(status, gin.H{
				"success": ok,
			})
		} else if err == nil {
			status = http.StatusConflict
		} else {
			c.JSON(status, gin.H{
				"error": err,
			})
		}
	} else {
		glog.Errorf("%s method has an error, which is : %s\n", methodName, err)
		c.JSON(status, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}
}

// GetAssetByStaff get all staff
func GetAssetByStaff(c *gin.Context) {
	staffID := c.Param("staffID")
	status := http.StatusBadRequest
	_, asset := service.GetAssetByStaffID(staffID)
	status = http.StatusOK
	c.JSON(status, gin.H{
		"asset": asset,
	})
}

// GetAllAssets get all staff
func GetAllAssets(c *gin.Context) {
	var results []model.StaffModel
	status := http.StatusBadRequest
	_, results = service.GetAllStaff()
	status = http.StatusOK
	c.JSON(status, gin.H{
		"staff": results,
	})
}
