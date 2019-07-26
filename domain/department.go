package domain

import (
	"context"

	"github.com/galahade/bus_incomes/util"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Department is for staff's
type Department struct {
	Domain `json:",inline" bson:",inline"`
	SN     string `json:"sn" bson:"sn"`
	Name   string `json:"name" bson:"name"`
}

//Insert add a Department to DB
func (department *Department) Insert() error {
	result, err := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionDepartment).InsertOne(context.TODO(), department)
	if err == nil {
		id, ok := result.InsertedID.(primitive.ObjectID)
		if !ok {
			glog.Errorf("Can't get mongo ID from insert result: %s./n", result.InsertedID)
		} else {
			department.ID = id
		}
	}
	return err
}

// Get query monthly all lines incomes. opt pass query otpions like sort order.
func (department *Department) Get() error {
	methodName := "department.Get"
	name := department.Name
	filter := bson.M{
		"name": bson.M{"$eq": name},
	}
	collection := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionDepartment)
	err := collection.FindOne(ctx, filter).Decode(&department)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			glog.Errorf("%s %s error : %#v\n", methodName, name, err)
		}
	}
	return err
}

//Delete a department from DB
func (department Department) Delete() (ok bool, err error) {
	filter := bson.M{
		"name": bson.M{"$eq": department.Name},
		"sn":   bson.M{"$eq": department.SN},
	}
	result, err := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionDepartment).DeleteOne(context.TODO(), filter)
	glog.Infof("Delete department name is %s ,count is %d", department.Name, result.DeletedCount)
	if result.DeletedCount != int64(0) && err == nil {
		ok = true
	}
	return ok, err
}
