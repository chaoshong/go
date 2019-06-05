package model

import (
	"time"
)

type Inventory struct {
	Id         uint
	SupplierId uint
	StockNum   int
	StockDays  int
	Sku        string
	ItemTitle  string
	CreateAt   time.Time
	UpdateAt   time.Time
	DeleteAt   time.Time
}

type InventorySerializer struct {
	Id        uint      `json:"id"`
	StockNum  int       `json:stockNum`
	StockDays int       `json:stockDays`
	Sku       string    `json:SKU`
	Title     string    `json:title`
	CreateAt  time.Time `json:"createAt"`
}

func (i *Inventory) Serializer() InventorySerializer {
	return InventorySerializer{
		Id:        i.Id,
		StockNum:  i.StockNum,
		StockDays: i.StockDays,
		Sku:       i.Sku,
		Title:     i.ItemTitle,
		CreateAt:  i.CreateAt,
	}
}
