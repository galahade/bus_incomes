package service

import "github.com/galahade/bus_incomes/domain"

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
	err := dep.Get()
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
