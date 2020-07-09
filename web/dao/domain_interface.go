package dao

import (
	"Ams/model"
	"github.com/jinzhu/gorm"
)

type JoinDomain interface {
	Add(domain *model.Domains) *model.Domains
	Del(id int) bool
	Update(_model *model.Domains, updates map[string]interface{}) bool
	GetDB() *gorm.DB
}
