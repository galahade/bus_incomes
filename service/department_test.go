package service

import (
	"fmt"
	"testing"

	"github.com/galahade/bus_incomes/domain"
	"github.com/stretchr/testify/assert"
)

func TestIsDepartmentExistFalse(t *testing.T) {
	department := domain.Department{
		Name: "maintian1",
	}
	results := IsDepartmentExist(&department)
	assert.False(t, results)
	fmt.Printf("Can we find department named %s ?: %+v\n", department.Name, results)
}

func TestIsDepartmentExistTrue(t *testing.T) {
	department := domain.Department{
		Name: "operation",
	}
	results := IsDepartmentExist(&department)
	assert.True(t, results)
	fmt.Printf("Can we find department named %s ?: %+v\n", department.Name, results)
}

func TestAddDepartmentFalse(t *testing.T) {
	department := domain.Department{
		Name: "maintian1",
	}
	ok, err := AddDepartment(&department)

	assert.False(t, ok)
	assert.Nil(t, err)
}

// Can't run it multi times
func TestAddDepartmentTrue(t *testing.T) {
	department := domain.Department{
		Name: "operation",
		SN:   "001",
	}
	ok, err := AddDepartment(&department)

	assert.True(t, ok)
	assert.Nil(t, err)
}

// Can't run it multi times
func TestRemoveDepartmentTrue(t *testing.T) {
	department := domain.Department{
		Name: "operation",
		SN:   "001",
	}
	ok := RemoveDepartment(department)

	assert.True(t, ok)
}
