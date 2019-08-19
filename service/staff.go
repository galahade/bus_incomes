package service

import (
	"github.com/galahade/bus_incomes/domain"
	"github.com/galahade/bus_incomes/model"
	"github.com/galahade/bus_incomes/util"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddStaff is used to insert a staff record into db
// name and employee id can't be empty
func AddStaff(staffModel *model.StaffModel) (ok bool, err error) {
	methodName := "AddStaff"
	glog.Info("Enter service methtod :" + methodName)

	if !(staffModel.Name == "" || staffModel.EmployeeID == "") {
		if staff, err1 := staffModel.PassToDomain(); err1 == nil {
			if !IsStaffExist(staff) {
				err := staff.Insert()
				if err == nil {
					ok = true
				} else {
					glog.Errorf("service %s instert err: %s", methodName, err)
				}
			}
		} else {
			err = err1
		}
	} else {
		glog.Errorf("staff's name or staff's employeeID is empty.")
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

// GetStaffByEmployeeID is used to get a staff by Employee id
func GetStaffByEmployeeID(employeeID string) (bool, domain.Staff) {
	ok := false
	staff := domain.Staff{
		EmployeeID: employeeID,
	}
	// Todo
	if err := staff.Select(); err == nil {
		ok = true
	}

	return ok, staff

}

// GetStaffByDepartmentID is used to get a staff by id
func GetStaffByDepartmentID(departmentID string) (bool, []model.StaffModel) {
	ok := false
	var results []model.StaffModel
	if objectID, err := primitive.ObjectIDFromHex(departmentID); err != nil {
		glog.Errorf("staff pass to domain err : %s", err)

	} else {

		staff := domain.SelectStaffByDepartmentID(objectID)

		if len(staff) > 0 {
			ok = true
			for _, s := range staff {
				staffModel := new(model.StaffModel)
				staffModel.InitFromDomain(s)
				(&staffModel.Department).SelectByID()

				results = append(results, *staffModel)
			}
		}
	}
	return ok, results

}

// GetAllStaff is used to get All staff to display
func GetAllStaff() (bool, []model.StaffModel) {
	staff := domain.SelectAllStaff()
	ok := false
	var results []model.StaffModel

	if len(staff) > 0 {
		ok = true
		for _, s := range staff {
			staffModel := new(model.StaffModel)
			staffModel.InitFromDomain(s)
			(&staffModel.Department).SelectByID()

			results = append(results, *staffModel)
		}
	}

	return ok, results
}

// GetAllJobType is used to get All staff to display
func GetAllJobType() (bool, []string) {
	ok := true
	var results []string

	results = append(results, util.JobDriver)
	results = append(results, util.JobMananger)
	results = append(results, util.JobHeadman)
	results = append(results, util.JobDispatcher)
	results = append(results, util.JobCommonStaff)
	results = append(results, util.JobCharger)
	results = append(results, util.JobMaintenanceMan)
	results = append(results, util.JobIT)

	return ok, results
}
