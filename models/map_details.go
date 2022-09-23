package models

import "gorm.io/gorm"

type MapDetail struct {
	gorm.Model
	MapName string
	MapDescription string
	CategoryName string
	MapStatus bool
	CreatedBy string
	UpdatedBy string
	Mappings []Mapping
	TenantMasterMappings []TenantMasterMapping
}