package domain

import (
	"context"
	"time"

	"github.com/galahade/bus_incomes/util"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Staff info
type Staff struct {
	Domain           `json:",inline" bson:",inline"`
	Name             string             `json:"name" bson:"name"`
	EmployeeID       string             `json:"employee_id" bson:"employee_id"`
	OnboardTime      time.Time          `json:"" bson:"onboard_time"`
	PersonalID       string             `json:"personal_id" bson:"personal_id"`
	IsMultitimeHired bool               `json:"is_multitime_hired" bson:"is_multitime_hired"`
	IsResign         bool               `json:"is_resign" bson:"is_resign"`
	FirstOnboardTime time.Time          `json:"first_onboard_time" bson:"first_onboard_time"`
	Phone            string             `json:"phone" bson:"phone"`
	DepartmentID     primitive.ObjectID `json:"department_id" bson:"department_id"`
	Job              string             `json:"job" bson:"job"`
}

// SelectByID query monthly all lines incomes. opt pass query otpions like sort order.
func (staff *Staff) SelectByID() error {
	methodName := "staff.SelectByID"
	id := staff.Domain.ID
	filter := bson.M{
		"_id": bson.M{"$eq": id},
	}
	collection := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionStaff)
	err := collection.FindOne(ctx, filter).Decode(&staff)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			glog.Errorf("%s %s error : %#v\n", methodName, id, err)
		}
	}
	return err
}

//Insert add a staff to DB
func (staff *Staff) Insert() error {
	result, err := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionStaff).InsertOne(context.TODO(), staff)
	if err == nil {
		id, ok := result.InsertedID.(primitive.ObjectID)
		if !ok {
			glog.Errorf("Can't get mongo ID from insert result: %s./n", result.InsertedID)
		} else {
			staff.ID = id
		}
	}
	return err
}

// Select query staff by employee id.
func (staff *Staff) Select() error {
	methodName := "department.Select"
	name := staff.Name
	filter := bson.M{
		"employee_id": bson.M{"$eq": staff.EmployeeID},
	}
	collection := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionStaff)
	err := collection.FindOne(ctx, filter).Decode(&staff)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			glog.Errorf("%s %s error : %#v\n", methodName, name, err)
		}
	}
	return err
}

//Delete a department from DB
func (staff Staff) Delete() (ok bool, err error) {
	filter := bson.M{
		"name":        bson.M{"$eq": staff.Name},
		"employee_id": bson.M{"$eq": staff.EmployeeID},
	}
	result, err := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionStaff).DeleteOne(context.TODO(), filter)
	glog.Infof("Delete department name is %s ,count is %d", staff.Name, result.DeletedCount)
	if result.DeletedCount != int64(0) && err == nil {
		ok = true
	}
	return ok, err
}
