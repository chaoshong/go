package domain

import (
	"github.com/chaoshong/go/dao"
	"github.com/chaoshong/go/model"
)

type InventoryDomain struct{}

var DefaultInvDomain = InventoryDomain{}

func (i *InventoryDomain) GetInventory() *model.Inventory {
	dao.DefaultInvDao.FindAll(&model.Inventory, "")
}

func (i *InventoryDomain) CreateInventory(*model.Inventory) {
	dao.DefaultInvDao.Insert(&model.Inventory)
}
