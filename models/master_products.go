package models

import "gorm.io/gorm"

type MasterProduct struct {
	gorm.Model
	ProductName string
	ProductCode string
	ProductDescription string
	CreatedBy string
	UpdatedBy string
	MasterModules []MasterModule
}