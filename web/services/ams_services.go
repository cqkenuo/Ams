package services

import "Ams/model"

type AmsServicesInterFace interface {
	ListDomains(pageNum, pageNo int) []*model.Domains
	ListChildDomains(parentId int) []*model.Domains
	DeleteDomain(id int) bool
}

type AmsWebService struct {
	
}

func (a *AmsWebService) ListDomains(pageNum, pageNo int) []*model.Domains {
	return []*model.Domains{{Domain: "qaaaa"}}
}

func (a *AmsWebService) ListChildDomains(parentId int) []*model.Domains {
	panic("implement me")
}

func (a *AmsWebService) DeleteDomain(id int) bool {
	panic("implement me")
}

func NewWebService()AmsServicesInterFace{
	return &AmsWebService{}
}