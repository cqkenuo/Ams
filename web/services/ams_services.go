package services

import "Ams/model"

type AmsServicesInterFace interface {
	ListDomains(pageNum, pageNo int) []*model.Domains
	ListChildDomains(parentId int) []*model.Domains
	DeleteDomain(id int) bool
}
