package models

import (
	"strconv"
	"testing"

	"github.com/sentiers/api-server/v2/config"

	"github.com/go-playground/assert/v2"
)

// Company Model Test ======================================
func TestCreateCompany(t *testing.T) {
	config.InitDatabase()
	company := Company{CompanyName: "createcompany"}
	if err := CreateCompany(&company); err != nil {
		t.Error("not created")
	}
	var company2 Company
	config.DB.Find(&company2, "company_name = ?", "createcompany")
	assert.Equal(t, company.CompanyName, company2.CompanyName)
}

func TestGetCompanyById(t *testing.T) {
	config.InitDatabase()

	company := Company{CompanyName: "getcompany"}
	if err := CreateCompany(&company); err != nil {
		t.Error("not created")
	}
	if err := GetCompanyById(&company, strconv.FormatUint(uint64(company.Id), 10)); err != nil {
		t.Error("not found")
	}
}

func TestUpdateCompany(t *testing.T) {
	config.InitDatabase()

	company := Company{CompanyName: "updatecompany"}
	if err := CreateCompany(&company); err != nil {
		t.Error("not created")
	}

	company.CompanyName = "updatedcompany"
	_ = UpdateCompany(&company)
	assert.Equal(t, "updatedcompany", company.CompanyName)
}

func TestDeleteCompany(t *testing.T) {
	config.InitDatabase()

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
	config.InitDatabase()
	department := Department{DepartmentName: "createdepartment", CompanyId: 1}
	if err := CreateDepartment(&department); err != nil {
		t.Error("not created")
	}
	var department2 Department
	config.DB.Find(&department2, "department_name = ?", "createdepartment")
	assert.Equal(t, department.DepartmentName, department2.DepartmentName)
}

func TestGetDepartmentById(t *testing.T) {
	config.InitDatabase()

	department := Department{DepartmentName: "getdepartment", CompanyId: 1}
	if err := CreateDepartment(&department); err != nil {
		t.Error("not created")
	}

	if err := GetDepartmentById(&department, strconv.FormatUint(uint64(department.Id), 10)); err != nil {
		t.Error("not found")
	}

}

func TestUpdateDepartment(t *testing.T) {
	config.InitDatabase()

	department := Department{DepartmentName: "updatedepartment", CompanyId: 1}
	if err := CreateDepartment(&department); err != nil {
		t.Error("not created")
	}

	department.DepartmentName = "updateddepartment"
	_ = UpdateDepartment(&department)
	assert.Equal(t, "updateddepartment", department.DepartmentName)
}

func TestDeleteDepartment(t *testing.T) {
	config.InitDatabase()

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
	config.InitDatabase()
	team := Team{TeamName: "createteam", DepartmentId: 1}
	if err := CreateTeam(&team); err != nil {
		t.Error("not created")
	}
	var team2 Team
	config.DB.Find(&team2, "team_name = ?", "createteam")
	assert.Equal(t, team.TeamName, team2.TeamName)
}

func TestGetTeamById(t *testing.T) {
	config.InitDatabase()

	team := Team{TeamName: "getteam", DepartmentId: 1}
	if err := CreateTeam(&team); err != nil {
		t.Error("not created")
	}
	if err := GetTeamById(&team, strconv.FormatUint(uint64(team.Id), 10)); err != nil {
		t.Error("not found")
	}
}

func TestUpdateTeam(t *testing.T) {
	config.InitDatabase()

	team := Team{TeamName: "updateteam", DepartmentId: 1}
	if err := CreateTeam(&team); err != nil {
		t.Error("not created")
	}

	team.TeamName = "updatedteam"
	_ = UpdateTeam(&team)
	assert.Equal(t, "updatedteam", team.TeamName)
}

func TestDeleteTeam(t *testing.T) {
	config.InitDatabase()

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
func TestCreateEmployee(t *testing.T) {
	config.InitDatabase()
	employee := Employee{EmployeeName: "createemployee", Email: "create@email.com", PhoneNumber: "010-1234-5678"}
	if err := CreateEmployee(&employee); err != nil {
		t.Error("not created")
	}
	var employee2 Employee
	config.DB.Find(&employee2, "employee_name = ?", "createemployee")
	assert.Equal(t, employee.EmployeeName, employee2.EmployeeName)
}

func TestGetEmployeeById(t *testing.T) {
	config.InitDatabase()

	employee := Employee{EmployeeName: "getemployee", Email: "get@email.com", PhoneNumber: "010-1234-5678"}
	if err := CreateEmployee(&employee); err != nil {
		t.Error("not created")
	}

	if err := GetEmployeeById(&employee, strconv.FormatUint(uint64(employee.Id), 10)); err != nil {
		t.Error("not found")
	}

}

func TestUpdateEmployee(t *testing.T) {
	config.InitDatabase()

	employee := Employee{EmployeeName: "updateemployee", Email: "update@email.com", PhoneNumber: "010-1234-5678"}
	if err := CreateEmployee(&employee); err != nil {
		t.Error("not created")
	}

	employee.EmployeeName = "updatedemployee"
	_ = UpdateEmployee(&employee)
	assert.Equal(t, "updatedemployee", employee.EmployeeName)
}

func TestDeleteEmployee(t *testing.T) {
	config.InitDatabase()

	employee := Employee{EmployeeName: "deleteemployee", Email: "delete@email.com", PhoneNumber: "010-1234-5678"}
	if err := CreateEmployee(&employee); err != nil {
		t.Error("not created")
	}
	var count1 int64
	config.DB.Table("employee").Count(&count1)

	_ = DeleteEmployee(&employee, strconv.FormatUint(uint64(employee.Id), 10))

	var count2 int64
	config.DB.Table("employee").Count(&count2)
	assert.Equal(t, count1-1, count2)
}

// =========================================
func TestGetEmployeeByName(t *testing.T) {
	config.InitDatabase()
	employee := Employee{EmployeeName: "testname", Email: "get@email.com", PhoneNumber: "010-1234-5678"}
	if err := CreateEmployee(&employee); err != nil {
		t.Error("not created")
	}
	var employees []Employee
	if err := GetEmployeeByName(&employees, "testname"); err != nil {
		t.Error("not found")
	}
}

func TestGetEmployeeByDate(t *testing.T) {
	config.InitDatabase()
	employee := Employee{EmployeeName: "todayname", Email: "get@email.com", PhoneNumber: "010-1234-5678"}
	if err := CreateEmployee(&employee); err != nil {
		t.Error("not created")
	}
	var employees []Employee
	if err := GetEmployeeByDate(&employees, employee.CreatedAt.String()); err != nil {
		t.Error("not found")
	}
}
