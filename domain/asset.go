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

// Asset info
type Asset struct {
	Domain     `json:",inline" bson:",inline"`
	Name       string             `json:"name" bson:"name"`
	Model      string             `json:"model" bson:"model"`
	SN         string             `json:"sn" bson:"sn"`
	Brand      string             `json:"brand" bson:"brand"`
	StartTime  time.Time          `json:"start_time" bson:"start_time"`
	Quantity   int64              `json:"quantity" bson:"quantity"`
	BuyingTime time.Time          `json:"buying_time,omitempty" bson:"buying_time,omitempty"`
	StaffID    primitive.ObjectID `json:"staff_id" bson:"staff_id"`
	Price      int64              `json:"price" bson:"price"`
	Note       string             `json:"note,omitempty" bson:"note,omitempty"`
}

// UpdateOne used to update an asset
func (asset Asset) UpdateOne() (err error) {
	methodName := "domain asset.UpdateOne"
	glog.Info("Enter domain methtod :" + methodName)
	id := asset.Domain.ID
	// set filters and updates
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{
		"start_time": asset.StartTime,
		"staff_id":   asset.StaffID,
		"note":       asset.Note,
	}}
	collection := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionAsset)
	if _, err = collection.UpdateOne(ctx, filter, update); err != nil {
		glog.Errorf("%s %s error : %#v\n", methodName, id, err)

	}
	return
}

// SelectByID query monthly all lines incomes. opt pass query otpions like sort order.
func (asset *Asset) SelectByID() error {
	methodName := "domain asset.SelectByID"
	glog.Info("Enter domain methtod :" + methodName)
	id := asset.Domain.ID
	filter := bson.M{
		"_id": bson.M{"$eq": id},
	}
	collection := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionAsset)
	err := collection.FindOne(ctx, filter).Decode(&asset)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			glog.Errorf("%s %s error : %#v\n", methodName, id, err)
		}
	}
	return err
}

//Insert add a asset to DB
func (asset *Asset) Insert() error {
	methodName := "domain asset.Insert"
	result, err := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionAsset).InsertOne(context.TODO(), asset)
	if err == nil {
		id, ok := result.InsertedID.(primitive.ObjectID)
		if !ok {
			glog.Errorf("%s can't get mongo ID from insert result: %s./n", methodName, result.InsertedID)
		} else {
			asset.ID = id
		}
	} else {
		glog.Errorf("%s error: %s", methodName, err)
	}
	return err
}

// Select query asset by sn.
func (asset *Asset) Select() error {
	methodName := "asset.Select"
	name := asset.Name
	filter := bson.M{
		"sn": bson.M{"$eq": asset.SN},
	}
	collection := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionAsset)
	err := collection.FindOne(ctx, filter).Decode(&asset)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			glog.Errorf("%s %s error : %#v\n", methodName, name, err)
		}
	}
	return err
}

//Delete a Asset from DB by SN
func (asset Asset) Delete() (ok bool, err error) {
	filter := bson.M{
		"sn": bson.M{"$eq": asset.SN},
	}
	result, err := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionAsset).DeleteOne(context.TODO(), filter)
	if result.DeletedCount != int64(0) && err == nil {
		ok = true
	}
	return ok, err
}

//SelectAssetsByStaffID is
func SelectAssetsByStaffID(staffID primitive.ObjectID) []Asset {
	methodName := "SelectAssetByDepartmentID"
	glog.Info("Enter domain methtod :" + methodName)
	findOpt := options.Find().SetSort(bson.M{"start_time": 1})
	filter := bson.M{
		"staff_id": bson.M{"$eq": staffID},
	}
	var results []Asset
	collection := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionAsset)
	if cursor, err := collection.Find(ctx, filter, findOpt); err != nil {
		glog.Errorf("Error happened on mongo find command. Error is %s", err)
	} else {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var asset Asset
			err := cursor.Decode(&asset)
			util.CheckErr(err, fmt.Sprintf("%s decode bson error", methodName))
			results = append(results, asset)
		}
	}

	return results
}

//SelectAllAssets is
func SelectAllAssets() []Asset {
	methodName := "SelectAllStaff"
	glog.Info("Enter domain methtod :" + methodName)
	findOpt := options.Find().SetSort(bson.M{"start_time": 1})

	var results []Asset
	collection := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionAsset)
	// Passing bson.D{{}} as the filter matches all documents in the collection
	if cursor, err := collection.Find(ctx, bson.D{{}}, findOpt); err != nil {
		glog.Errorf("Error happened on mongo find command. Error is %s", err)
	} else {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var staff Asset
			err := cursor.Decode(&staff)
			util.CheckErr(err, fmt.Sprintf("%s decode bson error", methodName))
			results = append(results, staff)
		}
	}
	return results
}
