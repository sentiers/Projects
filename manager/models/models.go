package models

import (
	"manager/config"
)

// Company Model=====================================================

type Company struct { // preload
	Id          uint   `gorm:"primarykey" json:"id"`
	CompanyName string `gorm:"type:varchar(255); not null" json:"companyname"`
}

func (c *Company) TableName() string {
	return "company"
}

func GetAllCompanies(company *[]Company) (err error) {
	if err = config.DB.Find(company).Error; err != nil {
		return err
	}
	return nil
}

func CreateCompany(company *Company) (err error) {
	if err = config.DB.Create(company).Error; err != nil {
		return err
	}
	return nil
}

func GetCompanyById(company *Company, key string) (err error) {
	if err = config.DB.First(&company, key).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCompany(company *Company) (err error) {
	config.DB.Save(company)
	return nil
}

func DeleteCompany(company *Company, key string) (err error) {
	config.DB.Delete(company, key)
	return nil
}

// Department Model=====================================================

type Department struct {
	Id             uint    `gorm:"primarykey" json:"id"`
	DepartmentName string  `gorm:"type:varchar(255); not null" json:"departmentname"`
	Company        Company `gorm:"foreignkey:CompanyId; references:Id"`
	CompanyId      uint    `json:"companyid"`
}

func (c *Department) TableName() string {
	return "department"
}

func GetAllDepartments(department *[]Department) (err error) {
	if err = config.DB.Find(department).Error; err != nil {
		return err
	}
	return nil
}

func CreateDepartment(department *Department) (err error) {
	if err = config.DB.Create(department).Error; err != nil {
		return err
	}
	return nil
}

func GetDepartmentById(department *Department, key string) (err error) {
	if err = config.DB.First(&department, key).Error; err != nil {
		return err
	}
	return nil
}

func UpdateDepartment(department *Department) (err error) {
	config.DB.Save(department)
	return nil
}

func DeleteDepartment(department *Department, key string) (err error) {
	config.DB.Delete(department, key)
	return nil
}

// Team Model=====================================================

type Team struct {
	Id           int        `gorm:"primarykey" json:"id"`
	TeamName     string     `gorm:"type:varchar(255); not null" json:"teamname"`
	Department   Department `gorm:"foreignkey:DepartmentId; references:Id"`
	DepartmentId uint       `json:"departmentid"`
}

func (c *Team) TableName() string {
	return "team"
}

func GetAllTeams(team *[]Team) (err error) {
	if err = config.DB.Find(team).Error; err != nil {
		return err
	}
	return nil
}

func CreateTeam(team *Team) (err error) {
	if err = config.DB.Create(team).Error; err != nil {
		return err
	}
	return nil
}

func GetTeamById(team *Team, key string) (err error) {
	if err = config.DB.First(&team, key).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTeam(team *Team) (err error) {
	config.DB.Save(team)
	return nil
}

func DeleteTeam(team *Team, key string) (err error) {
	config.DB.Delete(team, key)
	return nil
}

// Employee Model=====================================================

type Employee struct {
	Id           int    `gorm:"primarykey" json:"id"`
	EmployeeName string `gorm:"type:varchar(255); not null" json:"employeename"`
	Email        string `gorm:"type:varchar(255); not null" json:"email"`
	PhoneNumber  string `gorm:"type:varchar(255); not null" json:"phonenumber"`
	Team         Team   `gorm:"foreignkey:TeamId; references:Id"`
	TeamId       uint   `json:"teamid"`
}

func (c *Employee) TableName() string {
	return "employee"
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
