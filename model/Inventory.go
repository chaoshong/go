package model

import (
	"time"
)

type Inventory struct {
	Id         int
	SupplierId int
	StockNum   int
	StockDays  time.Time
	Sku        string
	ItemTitle  string
	CreateAt   time.Time
	UpdateAt   time.Time
	DeleteAt   time.Time
}
