package models

import (
	"manager/config"
)

type Employee struct {
	Id             int    `json:"id"`
	EmployeeName   string `json:"employeename"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phonenumber"`
	TeamName       string `json:"teamname"`
	DepartmentName string `json:"departmentname"`
	CompanyName    string `json:"companyname"`
}

func (c *Employee) TableName() string {
	return "Employee"
}

func GetAllEmployees(employee *[]Employee) (err error) {
	if err = config.DB.Find(employee).Error; err != nil {
		return err
	}
	return nil
}

func CreateEmployee(employee *Employee) (err error) {
	if err = config.DB.Create(employee).Error; err != nil {
		return err
	}
	return nil
}

func GetEmployeeById(employee *Employee, key string) (err error) {
	if err = config.DB.First(&employee, key).Error; err != nil {
		return err
	}
	return nil
}

func UpdateEmployee(employee *Employee) (err error) {
	config.DB.Save(employee)
	return nil
}

func DeleteEmployee(employee *Employee, key string) (err error) {
	config.DB.Delete(employee, key)
	return nil
}
