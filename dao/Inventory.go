package dao

import (
	"github.com/chaoshong/go/inframe/db"
	"github.com/chaoshong/go/model"
)

type InventoryDao struct{}

var DefaultInvDao = InventoryDao{}

//执行库存数据库存储操作
func (*InventoryDao) Insert(inventory *model.Inventory) {
	db.PostgreDb.Save(&inventory)
}

//查找库存数据库操作
func (*InventoryDao) FindAll(inventory *model.Inventory, where string) {
	db.PostgreDb.Find(&inventory, where)
}
