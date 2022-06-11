package interfaces

import (
	"github.com/anuragmahadik004/hr_api/datalayer"
	"github.com/anuragmahadik004/hr_api/models"
)

type IEmployee interface {
	GetAllEmployees() []models.Employee
}

type EmployeeRepository struct {
}

var dlEmp datalayer.DLEmployee

func (EmployeeRepository) GetAllEmployees() []models.Employee {
	return dlEmp.GetAllEmployees()
}

func (EmployeeRepository) GetEmployee(EmployeeId string) models.Employee {
	return dlEmp.GetEmployee(EmployeeId)
}

func (EmployeeRepository) SaveEmployee(Employee models.Employee) bool {
	return dlEmp.SaveEmployee(Employee)
}

func (EmployeeRepository) DeleteEmployee(EmployeeId string) bool {
	return dlEmp.DeleteEmployee(EmployeeId)
}
