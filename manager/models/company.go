package models

import (
	"manager/config"
)

type Company struct {
	Id          int    `json:"id"`
	CompanyName string `json:"companyname"`
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
