package dao

import (
	"Ams/model"
	"github.com/jinzhu/gorm"
)

type JoinDao struct {
	
}

func (j JoinDao) Add(domain *model.Domains) *model.Domains {
	panic("implement me")
}

func (j JoinDao) Del(id int) bool {
	panic("implement me")
}

func (j JoinDao) Update(_model *model.Domains, updates map[string]interface{}) bool {
	panic("implement me")
}

func (j JoinDao) GetDB() *gorm.DB {
	panic("implement me")
}

