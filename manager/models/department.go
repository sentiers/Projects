package models

import (
	"manager/config"
)

type Department struct {
	Id             int    `json:"id"`
	DepartmentName string `json:"departmentname"`
	CompanyName    string `json:"companyname"`
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
