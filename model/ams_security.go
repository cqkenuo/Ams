package model

import "Ams/config"

// Domains [...]
type Domains struct {
	ID     int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null"`
	Domain string `gorm:"unique;column:domain;type:varchar(120);not null"`
	DNS    string `gorm:"column:dns;type:json"`
	Title  string `gorm:"column:title;type:varchar(255)"`
	Status int    `gorm:"column:status;type:int(11) unsigned"`
	Fid    int    `gorm:"index;column:fid;type:int(11) unsigned;not null"`
}

func AddDomainRow(domain string, fid int) *Domains {
	db := GetAppDB(*config.LoadConfig())
	row := &Domains{Domain: domain, Fid: fid}
	db.Create(row)
	return row
}
