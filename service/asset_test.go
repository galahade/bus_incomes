package service

import (
	"testing"
	"time"

	"github.com/galahade/bus_incomes/domain"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Can't run it multitimes
func TestAddAssetTrue(t *testing.T) {
	_, staff := GetStaffByEmployeeID("010001")
	asset := domain.Asset{
		Name:       "电脑",
		Model:      "Mac pro",
		SN:         "1239292921013",
		Brand:      "Apple",
		Quantity:   1,
		StartTime:  time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
		StaffID:    staff.ID,
		Price:      1200000,
		BuyingTime: time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	ok, err := AddAsset(&asset)

	assert.True(t, ok)
	assert.Nil(t, err)
}

func TestIsAssetExist(t *testing.T) {
	asset := domain.Asset{
		SN: "1239292921013",
	}
	ok := IsAssetExist(&asset)
	assert.True(t, ok)
}

func TestRemoveAsset(t *testing.T) {
	asset := domain.Asset{
		SN: "1239292921013",
	}
	ok := RemoveAsset(asset)
	assert.True(t, ok)
}

func TestGetAssetByID(t *testing.T) {
	id, _ := primitive.ObjectIDFromHex("5d5c1c02c1842c2e63c03f17")
	ok, asset := GetAssetByID(id)
	assert.True(t, ok)
	assert.Equal(t, "电脑", asset.Name)
	assert.Equal(t, "Mac pro", asset.Model)
}

func TestGetAssetByStaffID(t *testing.T) {
	_, staff := GetStaffByEmployeeID("010001")
	ok, assets := GetAssetByStaffID(staff.ID.Hex())
	assert.True(t, ok)
	assert.NotEmpty(t, assets)
}
