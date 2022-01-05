package models

import (
	"manager/config"
)

type Team struct {
	Id             int    `json:"id"`
	TeamName       string `json:"teamname"`
	DepartmentName string `json:"departmentname"`
	CompanyName    string `json:"companyname"`
}

func (c *Team) TableName() string {
	return "Team"
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
