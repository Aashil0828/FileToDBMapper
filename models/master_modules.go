package models

import "gorm.io/gorm"

type MasterModule struct {
	gorm.Model
	MasterProductID uint
	ModuleName string
	ModuleCode string
	ModuleDescription string
	CreatedBy string
	UpdatedBy string
	MasterTemplates []MasterTemplate
}