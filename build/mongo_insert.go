package build

import (
	"fmt"
	"time"

	"context"

	"github.com/galahade/bus_incomes/domain"
	"github.com/galahade/bus_incomes/util"
)

//InsertTwoMonthTestData use to prepare test data in mongo
func InsertTwoMonthTestData() {

	collection := domain.Client.Database("bus").Collection("incomes")
	lineOneAprilIncome := domain.LineMonthIncome{
		Line: domain.Line{
			No: 1,
		},
		MonthIncome: domain.MonthIncome{
			Cash:        41000,
			Card:        6000,
			IncomeMonth: time.Date(2019, time.April, 1, 0, 0, 0, 0, time.UTC),
			Attendance:  200,
		},
	}

	lineTwoAprilIncome := domain.LineMonthIncome{
		Line: domain.Line{
			No: 2,
		},
		MonthIncome: domain.MonthIncome{
			Cash:        62000,
			Card:        8000,
			IncomeMonth: time.Date(2019, time.April, 1, 0, 0, 0, 0, time.UTC),
			Attendance:  255,
		},
	}

	lineThreeAprilIncome := domain.LineMonthIncome{
		Line: domain.Line{
			No: 3,
		},
		MonthIncome: domain.MonthIncome{
			Cash:        41000,
			Card:        5000,
			IncomeMonth: time.Date(2019, time.April, 1, 0, 0, 0, 0, time.UTC),
			Attendance:  190,
		},
	}

	lineSixAprilIncome := domain.LineMonthIncome{
		Line: domain.Line{
			No: 6,
		},
		MonthIncome: domain.MonthIncome{
			Cash:        71000,
			Card:        8000,
			IncomeMonth: time.Date(2019, time.April, 1, 0, 0, 0, 0, time.UTC),
			Attendance:  210,
		},
	}

	lineEightAprilIncome := domain.LineMonthIncome{
		Line: domain.Line{
			No: 8,
		},
		MonthIncome: domain.MonthIncome{
			Cash:        32000,
			Card:        4000,
			IncomeMonth: time.Date(2019, time.April, 1, 0, 0, 0, 0, time.UTC),
			Attendance:  180,
		},
	}

	lineTenAprilIncome := domain.LineMonthIncome{
		Line: domain.Line{
			No: 10,
		},
		MonthIncome: domain.MonthIncome{
			Cash:        82000,
			Card:        9000,
			IncomeMonth: time.Date(2019, time.April, 1, 0, 0, 0, 0, time.UTC),
			Attendance:  240,
		},
	}

	lineElevenAprilIncome := domain.LineMonthIncome{
		Line: domain.Line{
			No: 11,
		},
		MonthIncome: domain.MonthIncome{
			Cash:        42000,
			Card:        5000,
			IncomeMonth: time.Date(2019, time.April, 1, 0, 0, 0, 0, time.UTC),
			Attendance:  180,
		},
	}

	lineOneMarchIncome := domain.LineMonthIncome{
		Line: domain.Line{
			No: 1,
		},
		MonthIncome: domain.MonthIncome{
			Cash:        31000,
			Card:        3500,
			IncomeMonth: time.Date(2019, time.March, 1, 0, 0, 0, 0, time.UTC),
			Attendance:  180,
		},
	}

	lineTwoMarchIncome := domain.LineMonthIncome{
		Line: domain.Line{
			No: 2,
		},
		MonthIncome: domain.MonthIncome{
			Cash:        58000,
			Card:        4000,
			IncomeMonth: time.Date(2019, time.March, 1, 0, 0, 0, 0, time.UTC),
			Attendance:  240,
		},
	}

	lineThreeMarchIncome := domain.LineMonthIncome{
		Line: domain.Line{
			No: 3,
		},
		MonthIncome: domain.MonthIncome{
			Cash:        40000,
			Card:        7000,
			IncomeMonth: time.Date(2019, time.March, 1, 0, 0, 0, 0, time.UTC),
			Attendance:  190,
		},
	}

	lineSixMarchIncome := domain.LineMonthIncome{
		Line: domain.Line{
			No: 6,
		},
		MonthIncome: domain.MonthIncome{
			Cash:        65000,
			Card:        9000,
			IncomeMonth: time.Date(2019, time.March, 1, 0, 0, 0, 0, time.UTC),
			Attendance:  230,
		},
	}

	lineEightMarchIncome := domain.LineMonthIncome{
		Line: domain.Line{
			No: 8,
		},
		MonthIncome: domain.MonthIncome{
			Cash:        34000,
			Card:        5000,
			IncomeMonth: time.Date(2019, time.March, 1, 0, 0, 0, 0, time.UTC),
			Attendance:  175,
		},
	}

	lineTenMarchIncome := domain.LineMonthIncome{
		Line: domain.Line{
			No: 10,
		},
		MonthIncome: domain.MonthIncome{
			Cash:        79000,
			Card:        8000,
			IncomeMonth: time.Date(2019, time.March, 1, 0, 0, 0, 0, time.UTC),
			Attendance:  220,
		},
	}

	lineElevenMarchIncome := domain.LineMonthIncome{
		Line: domain.Line{
			No: 11,
		},
		MonthIncome: domain.MonthIncome{
			Cash:        39000,
			Card:        3000,
			IncomeMonth: time.Date(2019, time.March, 1, 0, 0, 0, 0, time.UTC),
			Attendance:  187,
		},
	}
	incomes := []interface{}{
		lineOneAprilIncome, lineTwoAprilIncome, lineThreeAprilIncome,
		lineSixAprilIncome, lineEightAprilIncome, lineTenAprilIncome, lineElevenAprilIncome,
		lineOneMarchIncome, lineTwoMarchIncome, lineThreeMarchIncome,
		lineSixMarchIncome, lineEightMarchIncome, lineTenMarchIncome, lineElevenMarchIncome,
	}

	insertResult, err := collection.InsertMany(context.TODO(), incomes)
	util.CheckErr(err)
	fmt.Println("Inserted multiple documents: ", insertResult.InsertedIDs)
}
