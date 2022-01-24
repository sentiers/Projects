package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/sentiers/api-server/config"
	"github.com/sentiers/api-server/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

// Company Controller ==============================================
func TestGetAllCompanies(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.GET("/company", GetAllCompanies)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/company", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateCompany(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.POST("/company", CreateCompany)

	company := models.Company{CompanyName: "controlcompany"}
	payload, _ := json.Marshal(&company)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/company", bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	//config.DB.Unscoped().Where("company_name = ?", company.CompanyName).Delete(&models.Company{})
}

func TestGetCompanyById(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.GET("/company/:id", GetCompanyById)

	company := models.Company{CompanyName: "controlcompany"}
	if err := models.CreateCompany(&company); err != nil {
		t.Error("not created")
	}
	str := "/company/" + strconv.FormatUint(uint64(company.Id), 10)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", str, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestUpdateCompany(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.PUT("/company/:id", UpdateCompany)

	company := models.Company{CompanyName: "controlcompany"}
	if err := models.CreateCompany(&company); err != nil {
		t.Error("not created")
	}
	str := "/company/" + strconv.FormatUint(uint64(company.Id), 10)

	newcompany := models.Company{CompanyName: "newcontrolcompany"}
	payload, _ := json.Marshal(&newcompany)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", str, bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestDeleteCompany(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.DELETE("/company/:id", DeleteCompany)

	company := models.Company{CompanyName: "controlcompany"}
	if err := models.CreateCompany(&company); err != nil {
		t.Error("not created")
	}
	str := "/company/" + strconv.FormatUint(uint64(company.Id), 10)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", str, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

// Department Controller ==============================================
func TestGetAllDepartments(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.GET("/department", GetAllDepartments)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/department", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateDepartment(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.POST("/department", CreateDepartment)

	department := models.Department{DepartmentName: "controldepartment", CompanyId: 1}
	payload, _ := json.Marshal(&department)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/department", bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetDepartmentById(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.GET("/department/:id", GetDepartmentById)

	department := models.Department{DepartmentName: "controldepartment", CompanyId: 1}
	if err := models.CreateDepartment(&department); err != nil {
		t.Error("not created")
	}
	str := "/department/" + strconv.FormatUint(uint64(department.Id), 10)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", str, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestUpdateDepartment(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.PUT("/department/:id", UpdateDepartment)

	department := models.Department{DepartmentName: "controldepartment", CompanyId: 1}
	if err := models.CreateDepartment(&department); err != nil {
		t.Error("not created")
	}
	str := "/department/" + strconv.FormatUint(uint64(department.Id), 10)

	newdepartment := models.Department{DepartmentName: "newcontroldepartment", CompanyId: 1}
	payload, _ := json.Marshal(&newdepartment)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", str, bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestDeleteDepartment(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.DELETE("/department/:id", DeleteDepartment)

	department := models.Department{DepartmentName: "controldepartment", CompanyId: 1}
	if err := models.CreateDepartment(&department); err != nil {
		t.Error("not created")
	}
	str := "/department/" + strconv.FormatUint(uint64(department.Id), 10)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", str, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

// Team Controller ==============================================

func TestGetAllTeams(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.GET("/team", GetAllTeams)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/team", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateTeam(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.POST("/team", CreateTeam)

	team := models.Team{TeamName: "controlteam", DepartmentId: 1}
	payload, _ := json.Marshal(&team)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/team", bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetTeamById(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.GET("/team/:id", GetTeamById)

	team := models.Team{TeamName: "controlteam", DepartmentId: 1}
	if err := models.CreateTeam(&team); err != nil {
		t.Error("not created")
	}
	str := "/team/" + strconv.FormatUint(uint64(team.Id), 10)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", str, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestUpdateTeam(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.PUT("/team/:id", UpdateTeam)

	team := models.Team{TeamName: "controlteam", DepartmentId: 1}
	if err := models.CreateTeam(&team); err != nil {
		t.Error("not created")
	}
	str := "/team/" + strconv.FormatUint(uint64(team.Id), 10)

	newteam := models.Team{TeamName: "newcontrolteam", DepartmentId: 1}
	payload, _ := json.Marshal(&newteam)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", str, bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestDeleteTeam(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.DELETE("/team/:id", DeleteTeam)

	team := models.Team{TeamName: "controlteam", DepartmentId: 1}
	if err := models.CreateTeam(&team); err != nil {
		t.Error("not created")
	}
	str := "/team/" + strconv.FormatUint(uint64(team.Id), 10)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", str, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

// Employee Controller ==============================================
func TestGetAllEmployees(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.GET("/employee", GetAllEmployees)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/employee", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateEmployee(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.POST("/employee", CreateEmployee)

	employee := models.Employee{EmployeeName: "controlemployee", Email: "control@email.com", PhoneNumber: "010-1234-5678"}
	payload, _ := json.Marshal(&employee)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/employee", bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetEmployeeById(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.GET("/employee/:id", GetEmployeeById)

	employee := models.Employee{EmployeeName: "controlemployee", Email: "control@email.com", PhoneNumber: "010-1234-5678"}
	if err := models.CreateEmployee(&employee); err != nil {
		t.Error("not created")
	}
	str := "/employee/" + strconv.FormatUint(uint64(employee.Id), 10)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", str, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestUpdateEmployee(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.PUT("/employee/:id", UpdateEmployee)

	employee := models.Employee{EmployeeName: "controlemployee", Email: "control@email.com", PhoneNumber: "010-1234-5678"}
	if err := models.CreateEmployee(&employee); err != nil {
		t.Error("not created")
	}
	str := "/employee/" + strconv.FormatUint(uint64(employee.Id), 10)

	newemployee := models.Employee{EmployeeName: "newcontrolemployee", Email: "control@email.com", PhoneNumber: "010-1234-5678"}
	payload, _ := json.Marshal(&newemployee)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", str, bytes.NewBuffer(payload))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestDeleteEmployee(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.DELETE("/employee/:id", DeleteEmployee)

	employee := models.Employee{EmployeeName: "controlemployee", Email: "control@email.com", PhoneNumber: "010-1234-5678"}
	if err := models.CreateEmployee(&employee); err != nil {
		t.Error("not created")
	}
	str := "/employee/" + strconv.FormatUint(uint64(employee.Id), 10)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", str, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

// =========================================

func TestGetEmployeeByName(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.GET("/employee/name/:name", GetEmployeeByName)

	employee := models.Employee{EmployeeName: "controlemployee", Email: "control@email.com", PhoneNumber: "010-1234-5678"}
	if err := models.CreateEmployee(&employee); err != nil {
		t.Error("not created")
	}
	str := "/employee/name/" + employee.EmployeeName

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", str, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetEmployeeByDate(t *testing.T) {
	config.InitDatabase()

	router := gin.Default()
	router.GET("/employee/date/:date", GetEmployeeByDate)

	employee := models.Employee{EmployeeName: "controlemployee", Email: "control@email.com", PhoneNumber: "010-1234-5678"}
	if err := models.CreateEmployee(&employee); err != nil {
		t.Error("not created")
	}
	str := "/employee/date/" + employee.CreatedAt.String()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", str, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
