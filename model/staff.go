package model

import (
	"time"

	"github.com/galahade/bus_incomes/domain"
	"github.com/golang/glog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Model is base struct for all model
type Model struct {
	ID string `json:"id,omitempty"`
}

// StaffModel is used to represent staff for display
type StaffModel struct {
	Model            `json:",inline" bson:",inline"`
	Name             string            `json:"name"`
	EmployeeID       string            `json:"employee_id"`
	OnboardTime      string            `json:"onboard_time"`
	PersonalID       string            `json:"personal_id"`
	IsMultitimeHired bool              `json:"is_multitime_hired"`
	IsResign         bool              `json:"is_resign"`
	FirstOnboardTime string            `json:"first_onboard_time"`
	Phone            string            `json:"phone"`
	Department       domain.Department `json:"department"`
	Job              string            `json:"job"`
}

//InitFromDomain is used to fill inital data from domain
func (staffModel *StaffModel) InitFromDomain(staff domain.Staff) {
	staffModel.ID = staff.ID.Hex()
	staffModel.Name = staff.Name
	staffModel.EmployeeID = staff.EmployeeID
	staffModel.OnboardTime = staff.OnboardTime.Format("2006-01-02")
	staffModel.PersonalID = staff.PersonalID
	staffModel.IsMultitimeHired = staff.IsMultitimeHired
	staffModel.IsResign = staff.IsResign
	staffModel.FirstOnboardTime = staff.FirstOnboardTime.Format("2006-01-02")
	staffModel.Phone = staff.Phone
	staffModel.Department.ID = staff.DepartmentID
	staffModel.Job = staff.Job
}

//PassToDomain is used to fill inital data from domain
func (staffModel *StaffModel) PassToDomain() (staff *domain.Staff, err error) {
	staff = new(domain.Staff)
	if staffModel.ID != "" {
		if staff.ID, err = primitive.ObjectIDFromHex(staffModel.ID); err != nil {
			glog.Errorf("staff pass to domain err : %s", err)
			return staff, err
		}
	}

	staff.Name = staffModel.Name
	staff.EmployeeID = staffModel.EmployeeID
	timeLayout := "2006-01-02"
	if staff.OnboardTime, err = time.Parse(timeLayout, staffModel.OnboardTime); err == nil {
		staff.PersonalID = staffModel.PersonalID
		staff.IsMultitimeHired = staffModel.IsMultitimeHired
		staff.IsResign = staffModel.IsResign
		staff.Phone = staffModel.Phone
		staff.DepartmentID = staffModel.Department.ID
		staff.Job = staffModel.Job
		if staffModel.FirstOnboardTime != "" {
			staff.FirstOnboardTime, err = time.Parse(timeLayout, staffModel.FirstOnboardTime)
		}
	}
	if err != nil {
		glog.Errorf("staff pass to domain err : %s", err)

	}

	return staff, err
}
