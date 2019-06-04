package model

//产品基本信息表
type Products struct {
	ID                int     `gorm: "primary_key"`
	ProductCategoryId int     `gorm: "not null"`
	SPU               string  `gorm: "type:char(10),NOT NULL"`
	Cname             string  `gorm: "type:varchar(127)"`
	Ename             string  `gorm: "type:varchar(127)"`
	Title             string  `gorm: "type:varchar(127)"`
	Price             float64 `gorm: "type:decimal(12,2) "`
	PriceUnit         string  `gorm: "type:varchar(31)"`
	Description       string  `gorm: "type:text"`
	Image             string  `gorm: "type:text"`
	Inventory         int
	KeyWords          string `gorm: "type:varchar(255)"`
	Tags              string `gorm: "type:varchar(127)"`
	State             int    `gorm: "not null"`
}
