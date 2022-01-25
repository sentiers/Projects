package models

import (
	"time"

	"github.com/sentiers/api-server/v2/config"
)

// Company Model=====================================================

type Company struct {
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
	Id             uint   `gorm:"primarykey" json:"id"`
	DepartmentName string `gorm:"type:varchar(255); not null" json:"departmentname"`
	// set foreignkey
	Company   Company `gorm:"foreignkey:CompanyId; references:Id" json:"company"`
	CompanyId uint    `json:"companyid"`
}

func (c *Department) TableName() string {
	return "department"
}

func GetAllDepartments(department *[]Department) (err error) {
	if err = config.DB.Preload("Company").Find(department).Error; err != nil {
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
	if err = config.DB.Preload("Company").First(&department, key).Error; err != nil {
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
	Id       uint   `gorm:"primarykey" json:"id"`
	TeamName string `gorm:"type:varchar(255); not null" json:"teamname"`
	// set foreignkey
	Department   Department `gorm:"foreignkey:DepartmentId; references:Id" json:"department"`
	DepartmentId uint       `json:"departmentid"`
	// set relation
	Employee []*Employee `gorm:"many2many:team_emp;" json:"employee"`
}

func (c *Team) TableName() string {
	return "team"
}

func GetAllTeams(team *[]Team) (err error) {
	if err = config.DB.Preload("Department.Company").Find(team).Error; err != nil {
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
	// with team member data
	if err = config.DB.Preload("Department.Company").Preload("Employee.Team.Department.Company").First(&team, key).Error; err != nil {
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
	Id           uint      `gorm:"primarykey" json:"id"`
	EmployeeName string    `gorm:"type:varchar(255); not null" json:"employeename"`
	Email        string    `gorm:"type:varchar(255); not null" json:"email"`
	PhoneNumber  string    `gorm:"type:varchar(255); not null" json:"phonenumber"`
	CreatedAt    time.Time `json:"createdat"`
	// set relation
	Team []*Team `gorm:"many2many:team_emp;" json:"team"`
}

func (c *Employee) TableName() string {
	return "employee"
}

func GetAllEmployees(employee *[]Employee, pagination *Pagination) (err error) {
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(Employee{}).Where(employee).Preload("Team.Department.Company").Find(employee)
	if err = result.Error; err != nil {
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

// ===========================================================

func GetEmployeeById(employee *Employee, key string) (err error) {
	if err = config.DB.Preload("Team.Department.Company").First(&employee, key).Error; err != nil {
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

// =========================================

func GetEmployeeByName(employee *[]Employee, key string) (err error) {
	if err = config.DB.Preload("Team.Department.Company").Find(&employee, "employee_name = ?", key).Error; err != nil {
		return err
	}
	return nil
}

func GetEmployeeByDate(employee *[]Employee, key string) (err error) {
	if err = config.DB.Preload("Team.Department.Company").Find(&employee, "created_at >= ?", key).Error; err != nil {
		return err
	}
	return nil
}

// Pagination Model=====================================================

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}
