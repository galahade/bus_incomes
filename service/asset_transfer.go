package service

import (
	"github.com/galahade/bus_incomes/domain"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddAssetTransfer is used to insert an asset record into db
func AddAssetTransfer(assetTransfer *domain.AssetTransfer) (ok bool, err error) {
	err = assetTransfer.Insert()
	if err == nil {
		ok = true
	} else {
		glog.Errorf("Add asset transfer error: %s", err)
	}

	return ok, err
}

//RemoveAssetTransfer is used to remove a staff
func RemoveAssetTransfer(assetTransfer domain.AssetTransfer) bool {
	ok, _ := assetTransfer.Delete()
	return ok
}

// GetAssetTransferByID is used to get a asset by id
func GetAssetTransferByID(id primitive.ObjectID) (bool, domain.AssetTransfer) {
	ok := false
	asset := domain.AssetTransfer{
		Domain: domain.Domain{
			ID: id,
		},
	}
	if err := asset.SelectByID(); err == nil {
		ok = true
	}
	return ok, asset
}

// GetAssetTransferByAssetID is used to get asset array by staff id
func GetAssetTransferByAssetID(assetID string) (ok bool, results []domain.AssetTransfer) {
	if objectID, err := primitive.ObjectIDFromHex(assetID); err != nil {
		glog.Errorf("asset transfer parse asset id to ObjectID error: %s", err)
	} else {
		results = domain.SelectAssetTransferByAssetID(objectID)
		ok = true
	}
	return ok, results

}

// GetAllAssetTransfers is used to get All Asset to display
func GetAllAssetTransfers() (ok bool, assetTransfers []domain.AssetTransfer) {
	assetTransfers = domain.SelectAllAssetTransfers()
	ok = true
	return ok, assetTransfers
}
