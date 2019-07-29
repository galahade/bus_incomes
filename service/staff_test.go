package service

import (
	"testing"
	"time"

	"github.com/galahade/bus_incomes/domain"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Can't run it multitimes
func TestAddStaffTrue(t *testing.T) {
	_, dep := GetDepartmentByName("operation")
	depID := dep.ID
	staff := domain.Staff{
		Name:             "许宁",
		EmployeeID:       "000005",
		OnboardTime:      time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
		PersonalID:       "132930198202071234",
		IsMultitimeHired: false,
		IsResign:         false,
		FirstOnboardTime: time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
		Phone:            "12345678987",
		DepartmentID:     depID,
		Job:              "GM",
	}

	ok, err := AddStaff(&staff)

	assert.True(t, ok)
	assert.Nil(t, err)
}

func TestIsStaffExist(t *testing.T) {
	staff := domain.Staff{
		Name:       "许宁",
		EmployeeID: "000005",
	}

	ok := IsStaffExist(&staff)
	assert.True(t, ok)
}

func TestRemoveStaff(t *testing.T) {
	staff := domain.Staff{
		Name:       "许宁",
		EmployeeID: "000005",
	}
	ok := RemoveStaff(staff)
	assert.True(t, ok)
}

func TestGetStaffByID(t *testing.T) {
	id, _ := primitive.ObjectIDFromHex("5d3f0ec033899abf7076c867")
	ok, staff := GetStaffByID(id)
	assert.True(t, ok)
	assert.Equal(t, "许宁", staff.Name)
	assert.Equal(t, "132930198202071234", staff.PersonalID)
}

func TestGetStaffByEmployeeID(t *testing.T) {
	ok, staff := GetStaffByEmployeeID("000005")
	assert.True(t, ok)
	assert.Equal(t, "132930198202071234", staff.PersonalID)
}
