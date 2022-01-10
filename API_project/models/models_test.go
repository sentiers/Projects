package models

import (
	"manager/admin"
	"manager/config"
	"strconv"
	"testing"

	"github.com/go-playground/assert/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbConnection() {
	dsn := "root:root@tcp(127.0.0.1:3306)/manager?charset=utf8mb4&parseTime=True&loc=Local"
	config.DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	_ = config.DB.AutoMigrate(&Company{}, &Department{}, &Team{}, &Employee{}, &admin.User{})
}

// Company Model Test ======================================
func TestCreateCompany(t *testing.T) {
	dbConnection()
	company := Company{CompanyName: "createcompany"}
	if err := CreateCompany(&company); err != nil {
		t.Error("not created")
	}
	var company2 Company
	config.DB.Find(&company2, "company_name = ?", "createcompany")
	assert.Equal(t, company.CompanyName, company2.CompanyName)
}

func TestGetAllCompanies(t *testing.T) {
	dbConnection()
	companies := []Company{{CompanyName: "testcompany1"},
		{CompanyName: "testcompany2"},
		{CompanyName: "testcompany3"},
	}
	var count1 int64
	config.DB.Table("company").Count(&count1)

	for _, i := range companies {
		_ = CreateCompany(&i)
	}

	var count2 int64
	config.DB.Table("company").Count(&count2)
	assert.Equal(t, count1+3, count2)
}

func TestUpdateCompany(t *testing.T) {
	dbConnection()

	company := Company{CompanyName: "updatecompany"}
	if err := CreateCompany(&company); err != nil {
		t.Error("not created")
	}

	company.CompanyName = "updatedcompany"
	_ = UpdateCompany(&company)
	assert.Equal(t, "updatedcompany", company.CompanyName)
}

func TestDeleteCompany(t *testing.T) {
	dbConnection()

	company := Company{CompanyName: "deletecompany"}
	if err := CreateCompany(&company); err != nil {
		t.Error("not created")
	}
	var count1 int64
	config.DB.Table("company").Count(&count1)

	_ = DeleteCompany(&company, strconv.FormatUint(uint64(company.Id), 10))

	var count2 int64
	config.DB.Table("company").Count(&count2)
	assert.Equal(t, count1-1, count2)
}

// Department Model Test ======================================
func TestCreateDepartment(t *testing.T) {
	dbConnection()
	department := Department{DepartmentName: "createdepartment", CompanyId: 1}
	if err := CreateDepartment(&department); err != nil {
		t.Error("not created")
	}
	var department2 Department
	config.DB.Find(&department2, "department_name = ?", "createdepartment")
	assert.Equal(t, department.DepartmentName, department2.DepartmentName)
}

func TestGetAllDepartments(t *testing.T) {
	dbConnection()
	departments := []Department{{DepartmentName: "testdepartment1", CompanyId: 1},
		{DepartmentName: "testdepartment2", CompanyId: 1},
		{DepartmentName: "testdepartment3", CompanyId: 1},
	}
	var count1 int64
	config.DB.Table("department").Count(&count1)

	for _, i := range departments {
		_ = CreateDepartment(&i)
	}

	var count2 int64
	config.DB.Table("department").Count(&count2)
	assert.Equal(t, count1+3, count2)
}

func TestUpdateDepartment(t *testing.T) {
	dbConnection()

	department := Department{DepartmentName: "updatedepartment", CompanyId: 1}
	if err := CreateDepartment(&department); err != nil {
		t.Error("not created")
	}

	department.DepartmentName = "updateddepartment"
	_ = UpdateDepartment(&department)
	assert.Equal(t, "updateddepartment", department.DepartmentName)
}

func TestDeleteDepartment(t *testing.T) {
	dbConnection()

	department := Department{DepartmentName: "deletedepartment", CompanyId: 1}
	if err := CreateDepartment(&department); err != nil {
		t.Error("not created")
	}
	var count1 int64
	config.DB.Table("department").Count(&count1)

	_ = DeleteDepartment(&department, strconv.FormatUint(uint64(department.Id), 10))

	var count2 int64
	config.DB.Table("department").Count(&count2)
	assert.Equal(t, count1-1, count2)
}

// Team Model Test ======================================
func TestCreateTeam(t *testing.T) {
	dbConnection()
	team := Team{TeamName: "createteam", DepartmentId: 1}
	if err := CreateTeam(&team); err != nil {
		t.Error("not created")
	}
	var team2 Team
	config.DB.Find(&team2, "team_name = ?", "createteam")
	assert.Equal(t, team.TeamName, team2.TeamName)
}

func TestGetAllTeams(t *testing.T) {
	dbConnection()
	teams := []Team{{TeamName: "testteam1", DepartmentId: 1},
		{TeamName: "testteam2", DepartmentId: 1},
		{TeamName: "testteam3", DepartmentId: 1},
	}
	var count1 int64
	config.DB.Table("team").Count(&count1)

	for _, i := range teams {
		_ = CreateTeam(&i)
	}

	var count2 int64
	config.DB.Table("team").Count(&count2)
	assert.Equal(t, count1+3, count2)
}

func TestUpdateTeam(t *testing.T) {
	dbConnection()

	team := Team{TeamName: "updateteam", DepartmentId: 1}
	if err := CreateTeam(&team); err != nil {
		t.Error("not created")
	}

	team.TeamName = "updatedteam"
	_ = UpdateTeam(&team)
	assert.Equal(t, "updatedteam", team.TeamName)
}

func TestDeleteTeam(t *testing.T) {
	dbConnection()

	team := Team{TeamName: "createteam", DepartmentId: 1}
	if err := CreateTeam(&team); err != nil {
		t.Error("not created")
	}
	var count1 int64
	config.DB.Table("team").Count(&count1)

	_ = DeleteTeam(&team, strconv.FormatUint(uint64(team.Id), 10))

	var count2 int64
	config.DB.Table("team").Count(&count2)
	assert.Equal(t, count1-1, count2)
}

// Employee Model Test ======================================
// func TestCreateEmployee(t *testing.T) {
// 	dbConnection()
// 	team := Team{TeamName: "createteam", DepartmentId: 1}
// 	if err := CreateTeam(&team); err != nil {
// 		t.Error("not created")
// 	}
// 	var team2 Team
// 	config.DB.Find(&team2, "team_name = ?", "createteam")
// 	assert.Equal(t, team.TeamName, team2.TeamName)
// }

// func TestGetAllEmployees(t *testing.T) {
// 	dbConnection()
// 	teams := []Team{{TeamName: "testteam1", DepartmentId: 1},
// 		{TeamName: "testteam2", DepartmentId: 1},
// 		{TeamName: "testteam3", DepartmentId: 1},
// 	}
// 	var count1 int64
// 	config.DB.Table("team").Count(&count1)

// 	for _, i := range teams {
// 		_ = CreateTeam(&i)
// 	}

// 	var count2 int64
// 	config.DB.Table("team").Count(&count2)
// 	assert.Equal(t, count1+3, count2)
// }

// func TestUpdateEmployee(t *testing.T) {
// 	dbConnection()

// 	team := Team{TeamName: "updateteam", DepartmentId: 1}
// 	if err := CreateTeam(&team); err != nil {
// 		t.Error("not created")
// 	}

// 	team.TeamName = "updatedteam"
// 	_ = UpdateTeam(&team)
// 	assert.Equal(t, "updatedteam", team.TeamName)
// }

// func TestDeleteEmployee(t *testing.T) {
// 	dbConnection()

// 	team := Team{TeamName: "createteam", DepartmentId: 1}
// 	if err := CreateTeam(&team); err != nil {
// 		t.Error("not created")
// 	}
// 	var count1 int64
// 	config.DB.Table("team").Count(&count1)

// 	_ = DeleteTeam(&team, strconv.FormatUint(uint64(team.Id), 10))

// 	var count2 int64
// 	config.DB.Table("team").Count(&count2)
// 	assert.Equal(t, count1-1, count2)
// }
