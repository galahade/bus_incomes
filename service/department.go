package service

import (
	"github.com/galahade/bus_incomes/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddDepartment is used to insert a department record into db
func AddDepartment(dep *domain.Department) (ok bool, err error) {
	if !(dep.Name == "" || dep.SN == "") {
		if !IsDepartmentExist(dep) {
			err := dep.Insert()
			if err == nil {
				ok = true
			}
		}
	}
	return ok, err
}

//IsDepartmentExist is used to check a department if exist
func IsDepartmentExist(dep *domain.Department) bool {
	ok := true
	err := dep.Select()
	if err != nil {
		ok = false
	}
	return ok
}

//RemoveDepartment is used to remove a department
func RemoveDepartment(dep domain.Department) bool {
	ok, _ := dep.Delete()

	return ok
}

// GetDepartmentByID is used to get a department by id
func GetDepartmentByID(id primitive.ObjectID) (bool, domain.Department) {
	ok := false
	dep := domain.Department{
		Domain: domain.Domain{
			ID: id,
		},
	}
	if err := dep.SelectByID(); err == nil {
		ok = true
	}

	return ok, dep

}

// GetDepartmentByName is used to get a department by id
func GetDepartmentByName(name string) (bool, domain.Department) {
	ok := false
	dep := domain.Department{
		Name: name,
	}
	if err := dep.Select(); err == nil {
		ok = true
	}

	return ok, dep

}
