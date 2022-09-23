package models

type FieldMapping struct{
	ID uint
	TemplateFieldName string `gorm:"type:varchar(100)"`
	CustomerFieldName string `gorm:"type:varchar(100)"`
	Mappings []Mapping
}