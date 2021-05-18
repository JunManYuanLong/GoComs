package model

import "ict.com/public.v1/model"

type Project struct {
	model.EntityModel
	Name        string  `gorm:"unique; comment:'项目名'" json:"name"`
	Description string  `gorm:"type:varchar(1000); comment:'项目描述'" json:"description"`
	Logo        string  `gorm:"type:varchar(2048); comment:'项目logo'" json:"logo"`
	Weight      int     `gorm:"default:1; comment:'权重'" json:"weight""`
	Category    string  `gorm:"type:varchar(20);comment:'项目类别'" json:"category"`
	Level       int     `gorm:"comment:'项目级别 公司级、部门级、条线级'" json:"level"` //Plevel 改动
	PM          int     `gorm:"comment:'项目负责人'" json:"pm"`               // proowner 改动
	TD          int     `gorm:"comment:'技术负责人'" json:"td"`               // techowner 改动
	Background  string  `gorm:"type:text;comment:'项目背景'" json:"background"`
	Worth       string  `gorm:"type:text;comment:'项目价值'" json:"worth"` // provalue 改动
	Target      string  `gorm:"type:text;comment:'项目目标'" json:"target"`
	Milestone   string  `gorm:"type:text;comment:'里程碑'" json:"milestone"`   // milepost 改动
	Budget      float64 `gorm:"comment:'预算'" json:"budget"`                 // cost 改动 类型 改动
	IsShow      int     `gorm:"comment:'项目是否展示 1 展示 0 不展示'" json:"is_show"` // show_ignore 改动
}
