package model

import "time"

type EntityModel struct {
	ID         int       `gorm:"primaryKey;autoIncrement;unique" json:"id"`
	Status     int       `gorm:"not null; default:0" json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Ext        string    `json:"ext"`
}

type EntityWithNameModel struct {
	ID         int       `gorm:"primaryKey;autoIncrement;unique" json:"id"`
	Status     int       `gorm:"not null; default:0" json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Ext        string    `json:"ext"`
	Name       string    `gorm:"type:varchar(100)" json:"name"`
}
