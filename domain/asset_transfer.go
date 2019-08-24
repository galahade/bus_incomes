package domain

import (
	"context"
	"fmt"
	"time"

	"github.com/galahade/bus_incomes/util"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AssetTransfer info
type AssetTransfer struct {
	Domain           `json:",inline" bson:",inline"`
	AssetID          primitive.ObjectID `json:"asset_id" bson:"asset_id"`
	OrginalStaffID   primitive.ObjectID `json:"orginal_staff_id" bson:"orginal_staff_id"`
	DepartmentID     primitive.ObjectID `json:"department_id" bson:"department_id"`
	StaffID          primitive.ObjectID `json:"staff_id" bson:"staff_id"`
	TransferTime     time.Time          `json:"transfer_time" bson:"transfer_time"`
	OrginalStartTime time.Time          `json:"orginal_start_time" bson:"orginal_start_time"`
	Note             string             `json:"note" bson:"note,omitempty"`
}

// SelectByID query monthly all lines incomes. opt pass query otpions like sort order.
func (assetTransfer *AssetTransfer) SelectByID() error {
	methodName := "domain assetTransfer.SelectByID"
	id := assetTransfer.Domain.ID
	filter := bson.M{
		"_id": bson.M{"$eq": id},
	}
	collection := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionAssetTransfer)
	err := collection.FindOne(ctx, filter).Decode(&assetTransfer)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			glog.Errorf("%s %s error : %#v\n", methodName, id, err)
		}
	}
	return err
}

//Insert add a asset to DB
func (assetTransfer *AssetTransfer) Insert() error {
	methodName := "domain asset_transfer.Insert"
	glog.Info("Enter methtod :" + methodName)
	result, err := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionAssetTransfer).InsertOne(context.TODO(), assetTransfer)
	if err == nil {
		id, ok := result.InsertedID.(primitive.ObjectID)
		if !ok {
			glog.Errorf("%s can't get mongo ID from insert result: %s./n", methodName, result.InsertedID)
		} else {
			assetTransfer.ID = id
		}
	} else {
		glog.Errorf("%s error: %s", methodName, err)
	}
	return err
}

//Delete a AssetTransfer from DB by id
func (assetTransfer AssetTransfer) Delete() (ok bool, err error) {
	filter := bson.M{
		"_id": bson.M{"$eq": assetTransfer.Domain.ID},
	}
	result, err := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionAssetTransfer).DeleteOne(context.TODO(), filter)
	if result.DeletedCount != int64(0) && err == nil {
		ok = true
	}
	return ok, err
}

//SelectAssetTransferByAssetID is
func SelectAssetTransferByAssetID(assetID primitive.ObjectID) []AssetTransfer {
	methodName := "SelectAssetTransferByAssetID"
	glog.Info("Enter domain methtod :" + methodName)
	findOpt := options.Find().SetSort(bson.M{"transfer_time": 1})
	filter := bson.M{
		"asset_id": bson.M{"$eq": assetID},
	}
	var results []AssetTransfer
	collection := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionAssetTransfer)
	if cursor, err := collection.Find(ctx, filter, findOpt); err != nil {
		glog.Errorf("Error happened on mongo find command. Error is %s", err)
	} else {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var assetTranfer AssetTransfer
			err := cursor.Decode(&assetTranfer)
			util.CheckErr(err, fmt.Sprintf("%s decode bson error", methodName))
			results = append(results, assetTranfer)
		}
	}

	return results
}

//SelectAllAssetTransfers is
func SelectAllAssetTransfers() []AssetTransfer {
	methodName := "SelectAllAssetTransfers"
	glog.Info("Enter domain methtod :" + methodName)
	findOpt := options.Find().SetSort(bson.M{"transfer_time": 1})

	var results []AssetTransfer
	collection := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionAssetTransfer)
	// Passing bson.D{{}} as the filter matches all documents in the collection
	if cursor, err := collection.Find(ctx, bson.D{{}}, findOpt); err != nil {
		glog.Errorf("Error happened on mongo find command. Error is %s", err)
	} else {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var staffTransfer AssetTransfer
			err := cursor.Decode(&staffTransfer)
			util.CheckErr(err, fmt.Sprintf("%s decode bson error", methodName))
			results = append(results, staffTransfer)
		}
	}
	return results
}
