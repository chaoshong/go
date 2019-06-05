package db

import (
	"fmt"

	"github.com/chaoshong/go/inframe/config"
	"github.com/chaoshong/go/model"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var PostgreDb *gorm.DB

func init() {
	var err error
	fmt.Println("init DB\n")
	dbType, dbString := config.GetDBInfo()
	fmt.Println("db type is :%s , db string is : %s", dbType, dbString)
	PostgreDb, err = gorm.Open(dbType, dbString)
	if err != nil {
		panic(err)
	}
	PostgreDb.AutoMigrate(
		&model.Products{},
		&model.Inventory{},
		&model.SKU{},
	)
	//&service.WarehouseList{}, &service.SPUList{}, &service.SKUList{}, &service.ItemPropertyList{}, &service.UDefinedCategoryList{},
	//&service.CategoryList{})
	//&models.Stock{}, &models.Employee{}, &models.OrderBalance{}, &models.OrderLittleBoss{},
	//&models.OrderIsale{}, &models.PostFee{}, &models.PostZone{}, &models.ProductCommon{},

	fmt.Println("create table\n")
}
