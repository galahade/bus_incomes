package domain

import (
	"context"
	"strings"
	"time"

	"github.com/golang/glog"

	"github.com/galahade/bus_incomes/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx context.Context
	//Client is mongodb client
	Client   *mongo.Client
	database *mongo.Database
)

// Domain is basic struct for all Domain
type Domain struct {
	ID  primitive.ObjectID `bson:"_id,omitempty"`
	JID string             `json:"id,omitempty"`
}

//Line is Bus line info
type Line struct {
	No int `json:"line_no" bson:"line_no"`
}

// MonthIncome is for line monthly cash and card income
type MonthIncome struct {
	Cash        int64     `json:"cash" bson:"cash"`
	Card        int64     `json:"card" bson:"card"`
	IncomeMonth time.Time `json:"month" bson:"income_month"`
	Attendance  int64     `json:"attendance" bson:"attendance"`
}

//MonthIncomes is all lines one month incomes
type MonthIncomes struct {
	LineMonthIncomes []LineMonthIncome `json:"lineMonthIncomes"`
	Month            string            `json:"month"`
}

//TowMonthsIncomes include last two month incomes
type TowMonthsIncomes struct {
	FinalIncomes    MonthIncomes `json:"finalIncomes"`
	PreviousIncomes MonthIncomes `json:"previousIncomes"`
	Error           string       `json:"error,omitempty"`
}

//SortedTwoMonthsIncomes is sorted Incomes by total incomes and average income
type SortedTwoMonthsIncomes struct {
	SortedIncomesByScore []TwoMonthsLineIncome `json:"two_month_incomes"`
	FinalMonth           string                `json:"final_month"`
	PreviousMonth        string                `json:"previous_month"`
}

//IncomeQuery is query string format like "yyyy-MM"
type IncomeQuery struct {
	Month string `json:"month"` //month string format like "yyyy-MM"
}

func init() {
	var err error
	ctx = context.Background()

	Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	util.CheckErr(err, "mongo connect error")
	// Check the connection
	util.CheckErr(Client.Ping(ctx, nil), "mongo client ping error")
}

// filterConditionByMonthOrAndLine can construct mongo query filter by time or time and line no.
func filterConditionByMonthOrAndLine(time time.Time, lineNo ...int) bson.M {
	var filter bson.M
	if lineNo != nil {
		filter = bson.M{
			"income.income_month": bson.M{"$eq": time},
			"line.line_no":        bson.M{"$eq": lineNo[0]},
		}
	} else {
		filter = bson.M{
			"income.income_month": bson.M{"$eq": time},
		}
	}
	glog.Infof("filter content is : %#v", filter)
	return filter
}

// GetYearAndMonth get year and month from query string
func (query IncomeQuery) GetYearAndMonth() (bool, int, time.Month) {
	s := strings.Split(query.Month, "-")
	return util.ConvertYearAndMonth(s[0], s[1])
}
