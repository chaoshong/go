package model

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/tealeg/xlsx"
)

type Stock struct {
	Id          int
	Sku         string `varchar(100)`
	ParentTitle string `varchar(255)`
	Category    string `varchar(100)`
	SaleDays    int
	SaleNums    int
	StockDay    time.Time
}

func (Stock) TableName() string {
	return "Inventory"
}

type SoldeazyStock struct {
	SKU         string
	Warehouse   string
	Stock_Level int
}
type Employee struct {
	Id        int
	Name      string
	ShortName string
	Platform  string
}

type Product struct {
	Id int
}

type Columns struct {
	xlsxColumns  []*xlsx.Cell
	tableColumns []string
	useColumns   [][]string
}
type row struct {
	insertID int64
	sql      string
	value    map[string]string
}

type ProductCommon struct {
	gorm.Model
	CnName                 string `gorm:"column:CnName"`
	Category               string
	PhotoFile              string `gorm:"column:PhotoFile"`
	SKU                    string `gorm:"column:SKU"`
	Title                  string
	Description            string
	EnName                 string `gorm:"column:EnName"`
	Color                  string
	CostPrice              float64 `gorm:"column:CostPrice"`
	SalesPrice             float64 `gorm:"column:SalesPrice"`
	ProductPrice           float64 `gorm:"column:ProductPrice"`
	ShippingCost           float64 `gorm:"column:ShippingCost"`
	TotalCost              float64 `gorm:"column:TotalCost"`
	ProfitRate             float64 `gorm:"column:ProfitRate"`
	PlatformPrice          float64 `gorm:"column:PlatformPrice"`
	MonthSale              string  `gorm:"column:MonthSale"`
	PlatformProfitRate     string  `gorm:"column:PlatformProfitRate"`
	PlatformCategory       string  `gorm:"column:PlatformCategory"`
	PlatformCategoryNumber string  `gorm:"column:PlatformCategoryNumber"`
	PlatformLink           string  `gorm:"column:PlatformLink"`
	Supplier               string
	Weight                 float64
	Length                 string
	Width                  string
	Height                 string
	PackageType            string `gorm:"column:PackageType"`
}

func (ProductCommon) TableName() string {
	return "ProductCommon"
}

type ProductSupplier struct {
	gorm.Model
}
