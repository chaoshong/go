package model

import (
	"time"
)

type PostZone struct {
	Id           int
	Zone         string
	Destination  string `gorm:"column:destination"`
	StartCode    int    `gorm:"column:startcode"`
	EndCode      int    `gorm:"column:endcode"`
	RemoteArea   string `gorm:"column:remotearea"`
	Supplier     string
	Country      string
	SupplierDate time.Time `gorm:"column:SupplierDate"`
	DeleteAt     time.Time `gorm:"column:DeleteAt"`
}

func (PostZone) TableName() string {
	return "PostZone"
}

type PostFee struct {
	Id           int
	Zone         string
	Destination  string  `gorm:"column:Destination"`
	StartWeight  float64 `gorm:"column:StartWeight"`
	EndWeight    float64 `gorm:"column:EndWeight"`
	BasicFee     float64 `gorm:"column:BasicFee"`
	PerKgFee     float64 `gorm:"column:PerKgFee"`
	RemoteArea   string  `gorm:"column:RemoteArea"`
	Supplier     string
	Country      string
	SupplierDate time.Time `gorm:"column:SupplierDate"`
	DeleteAt     time.Time `gorm:"column:DeleteAt"`
}

func (PostFee) TableName() string {
	return "PostFee"
}
