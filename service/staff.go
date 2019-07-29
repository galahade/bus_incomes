package service

import (
	"github.com/galahade/bus_incomes/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddStaff is used to insert a staff record into db
// name and employee id can't be empty
func AddStaff(staff *domain.Staff) (ok bool, err error) {
	if !(staff.Name == "" || staff.EmployeeID == "") {
		if !IsStaffExist(staff) {
			err := staff.Insert()
			if err == nil {
				ok = true
			}
		}
	}
	return ok, err
}

//IsStaffExist is used to check a staff if exist
func IsStaffExist(staff *domain.Staff) bool {
	ok := true
	err := staff.Select()
	if err != nil {
		ok = false
	}
	return ok
}

//RemoveStaff is used to remove a staff
func RemoveStaff(staff domain.Staff) bool {
	ok, _ := staff.Delete()

	return ok
}

// GetStaffByID is used to get a staff by id
func GetStaffByID(id primitive.ObjectID) (bool, domain.Staff) {
	ok := false
	staff := domain.Staff{
		Domain: domain.Domain{
			ID: id,
		},
	}
	if err := staff.SelectByID(); err == nil {
		ok = true
	}

	return ok, staff

}

// GetStaffByEmployeeID is used to get a staff by id
func GetStaffByEmployeeID(employeeID string) (bool, domain.Staff) {
	ok := false
	staff := domain.Staff{
		EmployeeID: employeeID,
	}
	if err := staff.Select(); err == nil {
		ok = true
	}

	return ok, staff

}
