package models

type MasterTemplate struct{
	ID uint
	MasterModuleID uint
	TemplateURL string
	TemplateName string
	TemplateCategory string
	TenantMasterMappings []TenantMasterMapping
}