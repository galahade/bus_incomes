package domain

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/galahade/bus_incomes/util"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Use command `go test -run TestMongoDBInsert` to test this method
func TestDeleteLineMonthIncome(t *testing.T) {
	income := LineMonthIncome{
		Line: Line{
			No: 20,
		},
		MonthIncome: MonthIncome{
			IncomeMonth: util.GetMonthlyStandardTime(2000, 1),
		},
	}
	//ok, err := DeleteLineMonthIncome(income)
	err := DeleteLineMonthIncome(income.IncomeMonth.Year(), income.IncomeMonth.Month(), income.Line.No)

	log.Print(err)
	assert.Nil(t, err)
}
func TestInsertLineMonthIncome(t *testing.T) {
	income := &LineMonthIncome{
		Line: Line{
			No: 20,
		},
		MonthIncome: MonthIncome{
			Cash:        100000,
			Card:        50000,
			IncomeMonth: util.GetMonthlyStandardTime(2000, 1),
			Attendance:  500,
		},
	}
	err := InsertLineMonthIncome(income)
	assert.Nil(t, err)
	assert.False(t, income.ID.IsZero())
	err = DeleteLineMonthIncome(income.IncomeMonth.Year(), income.IncomeMonth.Month(), income.Line.No)
	assert.Nil(t, err)
}

func TestSelectLineMonthlyIncome(t *testing.T) {
	ok, income := SelectLineMonthIncome(2000, 1, 20)
	assert.True(t, ok)
	assert.Equal(t, income.Cash, int64(41000))
	assert.Equal(t, income.Card, int64(6000))
	assert.Equal(t, income.Attendance, int64(200))
	assert.NotNil(t, income.ID)
}

func TestMongoDBQuery(t *testing.T) {
	fromDate := time.Date(2019, time.April, 1, 0, 0, 0, 0, time.UTC)
	toDate :=
		time.Date(2019, time.April, 2, 0, 0, 0, 0, time.UTC)
	findOpt := options.Find().SetSort(bson.M{"line.line_no": -1})
	filter := bson.M{"income.income_month": bson.M{
		"$gte": fromDate,
		"$lte": toDate,
	}}

	var results []LineMonthIncome
	collection := Client.Database("bus").Collection("incomes")
	cursor, err := collection.Find(ctx, filter, findOpt)

	defer cursor.Close(ctx)
	util.TestCheckErr(err, t, "collection find err")
	for cursor.Next(ctx) {
		var elem LineMonthIncome
		err := cursor.Decode(&elem)
		util.TestCheckErr(err, t, "Decode err")
		results = append(results, elem)
	}
	util.TestCheckErr(cursor.Err(), t, "cursor err")
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

}
