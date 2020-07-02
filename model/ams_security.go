package model

// Domains [...]
type Domains struct {
	ID     int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null"`
	Domain string `gorm:"column:domain;type:varchar(120);not null"`
	DNS    string `gorm:"column:dns;type:json"`
	Title  string `gorm:"column:title;type:varchar(255)"`
	Status int    `gorm:"column:status;type:int(11) unsigned"`
	Fid    int    `gorm:"index;column:fid;type:int(11) unsigned;not null"`
}

func(Domains)TableName()string{
	return "domains"
}