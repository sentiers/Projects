package config

import (
	"gorm.io/gorm"
)

// database that will be used across different packages
var DB *gorm.DB
