package service

import (
	"github.com/galahade/bus_incomes/domain"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddAsset is used to insert an asset record into db
func AddAsset(asset *domain.Asset) (ok bool, err error) {
	if !(asset.SN == "") {
		if !IsAssetExist(asset) {
			err := asset.Insert()
			if err == nil {
				ok = true
			}
		}
	}
	return ok, err
}

//IsAssetExist is used to check an asset if exist by it's sn
func IsAssetExist(asset *domain.Asset) bool {
	ok := true
	err := asset.Select()
	if err != nil {
		ok = false
	}
	return ok
}

//RemoveAsset is used to remove a staff
func RemoveAsset(asset domain.Asset) bool {
	ok, _ := asset.Delete()

	return ok
}

// GetAssetByID is used to get a asset by id
func GetAssetByID(id primitive.ObjectID) (bool, domain.Asset) {
	ok := false
	asset := domain.Asset{
		Domain: domain.Domain{
			ID: id,
		},
	}
	if err := asset.SelectByID(); err == nil {
		ok = true
	} else {
		glog.Errorf("get asset by id: %s error: %s", id.Hex(), err)
	}
	return ok, asset
}

// GetAssetBySID is used to get a asset by id
func GetAssetBySID(id string) (ok bool, asset domain.Asset) {

	if objectID, err := primitive.ObjectIDFromHex(id); err != nil {
		glog.Errorf("asset parse staff id to ObjectID error: %s", err)
	} else {
		asset := domain.Asset{
			Domain: domain.Domain{
				ID: objectID,
			},
		}
		if err := asset.SelectByID(); err == nil {
			ok = true
		}
	}

	return ok, asset
}

// GetAssetByStaffID is used to get asset array by staff id
func GetAssetByStaffID(staffID string) (ok bool, results []domain.Asset) {
	if objectID, err := primitive.ObjectIDFromHex(staffID); err != nil {
		glog.Errorf("asset parse staff id to ObjectID error: %s", err)
	} else {
		results = domain.SelectAssetsByStaffID(objectID)
		ok = true
	}
	return ok, results

}

// GetAllAsset is used to get All Asset to display
func GetAllAsset() (ok bool, assets []domain.Asset) {
	assets = domain.SelectAllAssets()
	ok = true
	return ok, assets
}

// ModifyAsset is used to get All Asset to display
func ModifyAsset(assetTransfer domain.AssetTransfer) (ok bool) {
	var asset domain.Asset
	if ok, asset = GetAssetByID(assetTransfer.AssetID); ok {
		asset.StartTime = assetTransfer.TransferTime
		asset.StaffID = assetTransfer.StaffID
		asset.Note = assetTransfer.Note + asset.Note
		if err := asset.UpdateOne(); err == nil {
			ok = true
		}
	}
	return ok
}
