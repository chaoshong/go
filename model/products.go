package model

//产品基本信息表
type Products struct {
	Id                int     `gorm:"primary_key"`
	ProductCategoryId int     `gorm:"not null;comment:'产品类别id'"`
	SPU               string  `gorm:"type:varchar(100);not null;comment:'产品SPU'"`
	Cname             string  `gorm:"type:varchar(127);comment:'中文名字'"`
	Ename             string  `gorm:"type:varchar(127);comment:'英文名字'"`
	Title             string  `gorm:"type:varchar(127);comment:'英文title'"`
	Price             float64 `gorm:"type:decimal(12,2);comment:'产品采购价格'"`
	PriceUnit         string  `gorm:"type:varchar(31);comment:'产品采购价格币种'"`
	Description       string  `gorm:"type:text;comment:'产品描述'"`
	Image             string  `gorm:"type:text;comment:'产品主图'"`
	Inventory         int     `gorm:"default:0;comment:'产品库存'"`
	KeyWords          string  `gorm:"type:varchar(255);COMMENT '产品关键词'"`
	Tags              string  `gorm:"type:varchar(127);COMMENT:'产品标签'"`
	State             int     `gorm:"not null;default:0;COMMENT:'产品状态，默认0'"`
}

type SKU struct {
	Id               int     `gorm:"primary_key"`
	SPUId            int     `gorm:"not null;comment:'关联产品SPUid'"`
	SKU              string  `gorm:"type:varchar(100);not null;comment:'产品SKU'"`
	Ename            string  `gorm:"type:varchar(127);comment:'英文名字'"`
	SalePrice        float64 `gorm:"type:decimal(12,2);comment:'产品销售价格'"`
	SalePriceUnit    string  `gorm:"type:varchar(31);comment:'产品销售价格币种'"`
	ShippingCost     float64 `gorm:"type:decimal(12,2);comment:'产品运费'"`
	ShippingCostUnit string  `gorm:"type:varchar(31);comment:'产品运费币种'"`
	SuppilerId       int     `gorm:"not null;default:1;comment:'产品供应商id'"`
	Weight           float64 `gorm:"type:decimal(12,2);comment:'产品重量 g'"`
	Height           float64 `gorm:"type:decimal(12,2);comment:'产品高度 cm'"`
	Width            float64 `gorm:"type:decimal(12,2);comment:'产品宽度 cm'"`
	Length           float64 `gorm:"type:decimal(12,2);comment:'产品长度 cm'"`
	Inventory        int     `gorm:"default:0;comment:'产品库存'"`
}
