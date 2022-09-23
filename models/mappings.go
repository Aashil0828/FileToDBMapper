package models

import "gorm.io/gorm"

type Mapping struct {
	ID             uint
	MapDetailID    uint
	FieldMappingID uint
	DeletedAt      gorm.DeletedAt
}