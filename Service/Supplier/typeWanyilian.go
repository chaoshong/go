package service

import "time"

type PageParams struct {
	PageSize   int `json:"pageSize"`
	PageNo     int `json:"pageNo"`
	TotalCount int `json:"totalCount"`
}
type WylInArg struct {
	SPU                string
	SKU                string
	UserDefinedCode    string     `json:"userDefinedCode"`
	CategoryID         string     `json:"categoryID"`
	UDefinedCategoryID string     `json:"uDefinedCategoryID"`
	Keywords           string     `json:"keywords"`
	WarehouseName      string     `json:"warehouseName"`
	WarehouseCode      string     `json:"warehouseCode"`
	IsHavingInventory  string     `json:"isHavingInventory"`
	PageParams         PageParams `json:"pageParams"`
	UpdateStartDate    string     `json:"updateStartDate"`
	UpdateEndDate      string     `json:"updateEndDate"`
}

/*WylInJSON 入参Json结构*/
type WylInJSON struct {
	Action     string   `json:"action"`
	AppKey     string   `json:"app_key"`
	Data       WylInArg `json:"data"`
	Format     string   `json:"format"`
	Language   string   `json:"language"`
	Platform   string   `json:"flatform"`
	Sign       string   `json:"sign"`
	SignMethod string   `json:"sign_method"`
	Timestamp  string   `json:"stimestamp"`
	Version    string   `json:"version"`
}

type WylResult struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		SPUList    []SPUList
		PageParams PageParams `json:"pageParams"`
	} `json:"data"`
}
type SPUList struct {
	SPU                  string
	UserDefinedCode      string                 `json:"userDefinedCode"`
	ChineseName          string                 `json:"chineseName"`
	EnglishName          string                 `json:"englishName"`
	CategoryList         []CategoryList         `json:"categoryList"`
	UDefinedCategoryList []UDefinedCategoryList `json:"UDefinedCategoryList"`
	Keywords             string                 `json:"keywords"`
	Description          string                 `json:"description"`
	WarehouseName        string                 `json:"warehouseName"`
	WarehouseCode        string                 `json:"warehouseCode"`
	CategoryID           float64                `json:"categoryID"`
	SKUList              []SKUList
	EbayListingSubsidy   float64   `json:"ebayListingSubsidy"`
	ImgList              []string  `json:"imgList" gorm:"-"`
	Img                  string    `json:"-"`
	CreatedAt            time.Time `json:"-"`
	ID                   uint      `json:"-"`
}
type SKUList struct {
	SKU              string
	RandomSKU        string             `json:"randomSKU"`
	SupplyInventory  int                `json:"supplyInventory"`
	Sepcification    string             `json:"sepcification"`
	Rebate           int                `json:"imgrebateList"`
	SupplyPrice      float64            `json:"supplyPrice"`
	SettlePrice      float64            `json:"settlePrice"`
	MinRetailPrice   float64            `json:"minRetailPrice"`
	RetailPriceCode  string             `json:"retailPriceCode"`
	Weight           float64            `json:"weight"`
	Length           float64            `json:"length"`
	Width            float64            `json:"width"`
	Height           float64            `json:"height"`
	ItemPropertyList []ItemPropertyList `json:"itemPropertyList"`
	CreatedAt        time.Time          `json:"-"`
	SPUListID        uint               `json:"-"`
	ID               uint               `json:"-"`
}

type ItemPropertyList struct {
	key   string `json:"key"`
	Value string `json:"value"`
	SKUID uint   `json:"-"`
	ID    uint   `json:"-"`
}
type UDefinedCategoryList struct {
	UDefinedCategoryID   int    `json:"uDefinedCategoryID"`
	UDefinedCategoryName string `json:"uDefinedCategoryName"`
	UDfatherCategoryID   int    `json:"uDfatherCategoryID"`
	SPUListID            uint   `json:"-"`
	ID                   uint   `json:"-"`
}

type CategoryList struct {
	CategoryID       int    `json:"categoryID"`
	CategoryName     string `json:"categoryName"`
	FatherCategoryID int    `json:"fatherCategoryID"`
	SPUListID        uint   `json:"-"`
	ID               uint   `json:"-"`
}
type WarehouseResult struct {
	Code string          `json:"code"`
	Msg  string          `json:"msg"`
	Data []WarehouseList `json:"data"`
}
type WarehouseList struct {
	WarehouseCode    string    `json:"warehouseCode"`
	WarehouseName    string    `json:"warehouseName"`
	WarehouseAddress string    `json:"warehouseAddress"`
	CreatedAt        time.Time `json:"-"`
	ID               uint      `json:"-"`
}
