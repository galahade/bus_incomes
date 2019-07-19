package domain

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/galahade/bus_incomes/util"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindMonthLinesIncomesOrderByLineNo query monthly all lines incomes. opt pass query otpions like sort order.
func FindMonthLinesIncomesOrderByLineNo(year int, month time.Month, opt ...*options.FindOptions) (bool, []LineMonthIncome) {
	findOpt := options.Find().SetSort(bson.M{"line.line_no": 1})
	return FindMonthlyLinesIncomes(year, month, findOpt)
}

// FindMonthlyLinesIncomes query monthly all lines incomes. opt pass query otpions like sort order.
func FindMonthlyLinesIncomes(year int, month time.Month, opt ...*options.FindOptions) (bool, []LineMonthIncome) {
	ok := true
	methodName := "QueryMonthlyLinesIncome"
	glog.Warning("Enter methtod :" + methodName)
	filter := filterConditionByMonthOrAndLine(util.GetMonthlyStandardTime(year, month))
	var results []LineMonthIncome
	collection := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionIncomes)

	if cursor, err := collection.Find(ctx, filter, opt...); err != nil {
		ok = false
		glog.Errorf("Error happened on mongo find command. Error is %s", err)
	} else {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var elem LineMonthIncome
			err := cursor.Decode(&elem)
			util.CheckErr(err, fmt.Sprintf("%s decode bson to MonthLineIncome error", methodName))
			results = append(results, elem)
		}
	}

	return ok, results
}

//InsertLineMonthIncome add a LineMonthIncome to DB
func InsertLineMonthIncome(income *LineMonthIncome) error {
	result, err := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionIncomes).InsertOne(context.TODO(), income)
	if err == nil {
		id, ok := result.InsertedID.(primitive.ObjectID)
		if !ok {
			log.Printf("Can't get ID from insert result: %s./n", result.InsertedID)
		} else {
			income.ID = id
		}
	}
	return err
}

//DeleteLineMonthIncome delete a LineMonthIncome to DB
func DeleteLineMonthIncome(year int, month time.Month, line int) (err error) {
	filter := filterConditionByMonthOrAndLine(util.GetMonthlyStandardTime(year, month), line)
	result, err := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionIncomes).DeleteOne(context.TODO(), filter)
	glog.Infof("Delete count is %d", result.DeletedCount)
	if result.DeletedCount == int64(0) && err == nil {
		err = fmt.Errorf("Can't delete Line %d Month %d-%d's income from DB", line, year, month)
	}
	return err
}

// SelectLineMonthIncome query monthly all lines incomes. opt pass query otpions like sort order.
func SelectLineMonthIncome(year int, month time.Month, lineNO int) (bool, LineMonthIncome) {
	methodName := "SelectLineMonthlyIncome"
	ok := true
	filter := filterConditionByMonthOrAndLine(util.GetMonthlyStandardTime(year, month), lineNO)
	collection := Client.Database(util.MongoDBName).Collection(util.BusDBCollectionIncomes)
	var income LineMonthIncome
	err := collection.FindOne(ctx, filter).Decode(&income)
	if err != nil {
		ok = false
		if err != mongo.ErrNoDocuments {
			log.Printf("%s select %d line's %d-%d's income error : %#v\n", methodName, lineNO, year, month, err)
		}
	}
	return ok, income
}
