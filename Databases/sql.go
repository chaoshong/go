package Databases

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"reflect"

	models "github.com/chaoshong/go/Models"
	service "github.com/chaoshong/go/Service/Supplier"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var Db *gorm.DB

func Init() {
	var err error
	fmt.Println("open\n")
	Db, err = gorm.Open("postgres", "user=postgres password=hgz dbname=postgres sslmode=disable")
	// dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
	// 	"root",
	// 	"hgz",
	// 	"127.0.0.1:3306",
	// 	"ecom")
	// Db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("open success\n")
	Db.AutoMigrate(&service.WarehouseList{}, &service.SPUList{}, &service.SKUList{}, &service.ItemPropertyList{}, &service.UDefinedCategoryList{},
		&service.CategoryList{})
	//&models.Stock{}, &models.Employee{}, &models.OrderBalance{}, &models.OrderLittleBoss{},
	//&models.OrderIsale{}, &models.PostFee{}, &models.PostZone{}, &models.ProductCommon{},

	fmt.Println("create table\n")

}

func Column(tableName string) {
	stock := models.Stock{}
	query := "select * from " + tableName + " limit 1"
	fmt.Println("query \n", query)
	columns := Db.NewScope(stock).Fields()
	// checkerr(err1, "get row")
	fmt.Println("get table columns\n", reflect.TypeOf(columns))
	// column, err := row.Columns()
	// checkerr(err, "get fields")
	for column := range columns {
		fmt.Println("field:\n", reflect.TypeOf(column))
	}
}

func WriteStocks(stocks []models.Stock) {
	// for i, stock := range stocks {
	// 	if i < 100 {
	// 		fmt.Println("stock : \n", stock)
	// 	}
	// }

	// stock := Stock{0, "653", "test", "1", 0, 0, time.Now()}
	fmt.Println("start insert table : \n")
	for _, stock := range stocks {
		if stock.Sku != "" || stock.ParentTitle != "" {
			stockDb := Db.Where(&models.Stock{Sku: stock.Sku, SaleNums: stock.SaleNums, StockDay: stock.StockDay}).First(&stock)
			//fmt.Println("find table : \n", stockDb)
			if stockDb != nil {
				Db.Create(&stock)
			}
		}
	}

}

func GetStocks(stockDay time.Time) (stocks []models.Stock, count int) {

	Db.Where(&models.Stock{StockDay: stockDay}).Find(&stocks).Count(&count)

	return stocks, count
}

func GetEmployee() (employee []models.Employee) {

	Db.Find(&employee)

	return employee
}

func WriteEmployee(em []models.Employee) {
	fmt.Println("start insert employee : \n")
	for _, e := range em {

		Db.Create(&e)

	}
}
func WriteOrdersLittleBoss(ordersLbs []models.OrderLittleBoss) {

	fmt.Println("start insert table OrdersLittleBoss : \n")
	for _, ordersLb := range ordersLbs {
		if ordersLb.OrderId != "" || ordersLb.ListingSKu != "" || ordersLb.Quantity != 0 {

			Db.FirstOrCreate(&ordersLb, models.OrderLittleBoss{ConsigneeName: ordersLb.ConsigneeName, OrderId: ordersLb.OrderId, ListingSKu: ordersLb.ListingSKu, Quantity: ordersLb.Quantity})

		}
	}

}

func GetOrderLittleBoss(orderHandled string) (orderlb []models.OrderLittleBoss, count int) {
	Db.Where(&models.OrderLittleBoss{OrderHandled: orderHandled}).Find(&orderlb).Count(&count)

	return orderlb, count
}
func GetOrderBalance(orderHandled string) (orderbln []models.OrderBalance, count int) {
	Db.Where(&models.OrderBalance{OrderHandled: orderHandled}).Find(&orderbln).Count(&count)

	return orderbln, count
}
func UpdateBalanceHandle(orderbln []models.OrderBalance) {
	Db.Model(&orderbln).Update("OrderHandled", "1")
}

func WriteOrdersIsale(ordersIsale []models.OrderIsale) {

	fmt.Println("start insert table OrderIsale : \n")
	for _, orderIsale := range ordersIsale {
		if orderIsale.ListingSKu != "" || orderIsale.Quantity != 0 {
			//orderIsale := Db.Where(&Stock{Sku: stock.Sku, SaleNums: stock.SaleNums, StockDay: stock.StockDay}).First(&stock)
			//fmt.Println("find table : \n", stockDb)
			//if stockDb != nil {
			Db.FirstOrCreate(&orderIsale, models.OrderIsale{ConsigneeName: orderIsale.ConsigneeName, OrderId: orderIsale.OrderId, ListingSKu: orderIsale.ListingSKu, Quantity: orderIsale.Quantity})
			//}
		}
	}

}
func WriteOrdersBalance(ordersBalance []models.OrderBalance) {

	fmt.Println("start insert table OrderBalance : \n")
	for _, value := range ordersBalance {
		if value.OrderId != "" || value.SalesSku != "" || value.Qty != 0 {
			//orderIsale := Db.Where(&Stock{Sku: stock.Sku, SaleNums: stock.SaleNums, StockDay: stock.StockDay}).First(&stock)
			//fmt.Println("find table : \n", stockDb)
			//if stockDb != nil {
			Db.FirstOrCreate(&value, models.OrderBalance{FullName: value.FullName, OrderId: value.OrderId, SalesSku: value.SalesSku, Qty: value.Qty})
			//}
		}
	}

}

func GetCostOrder(order models.OrderBalance) (rtnOrder models.OrderBalance) {
	fmt.Println("start get table OrderBalance by getcostorder : \n")
	stock := models.Stock{}
	product := models.ProductCommon{}
	startcode, _ := strconv.Atoi(strings.TrimSpace(order.Postcode))
	//fmt.Printf("get cost order startcode is %d, order is %s ,err is %s \n", startcode, order.Postcode, err)
	endcode := startcode
	postzone := models.PostZone{}
	postfee := models.PostFee{}
	Db.Where(&models.Stock{Sku: order.ProductSKU}).Find(&stock)
	Db.Where(&models.ProductCommon{SKU: order.ProductSKU}).Find(&product)
	//fmt.Printf("get cost order stock num is %d, stock date is %s ,product price is %d \n", stock.SaleNums, stock.StockDay, product.ProductPrice)
	postzone = GetPostCodeDest(models.PostZone{StartCode: startcode, EndCode: endcode}, "0")

	if postzone.Destination == "" {
		postzone = GetPostCodeDest(models.PostZone{StartCode: startcode, EndCode: endcode}, "1")
	}
	if postzone.RemoteArea == "" {
		postzone.RemoteArea = "0"
	}

	weight := product.Weight * float64(order.ProductQty)
	Db.Where("\"StartWeight\"<= ? and \"EndWeight\">=? and \"RemoteArea\" =?", weight, weight, postzone.RemoteArea).First(&postfee)
	//fmt.Printf("get cost order postzone is %s,  \n post fee is %s \n", postzone, postfee)
	rtnOrder = models.OrderBalance{}
	rtnOrder.ProductStockDate = stock.StockDay
	rtnOrder.ProductStockNum = stock.SaleNums
	rtnOrder.CostPrice = product.ProductPrice * float64(order.ProductQty)
	rtnOrder.ShippingCostBySystem = postfee.BasicFee + postfee.PerKgFee*weight/1000
	//fmt.Printf("get cost order ,order is %s \n", rtnOrder)
	return rtnOrder
}

func GetPostCodeDest(postcode models.PostZone, remotearea string) (returnpc models.PostZone) {
	//fmt.Printf("postcode is %s \n", &postcode)

	//Db.Find(&postcode, "'StartCode' <= ? and 'EndCode' >= ? and 'RemoteArea' = ?", postcode.StartCode, postcode.EndCode, "1")
	Db.Where("StartCode <= ? and EndCode >= ? and RemoteArea=?", postcode.StartCode, postcode.EndCode, remotearea).First(&returnpc)

	//fmt.Printf("sql is %s \n", returnpc)
	return returnpc
}

func GetPostFee(postcode models.PostZone, remotearea string) (returnpf models.PostFee) {
	//fmt.Printf("postcode is %s \n", &postcode)

	//Db.Find(&postcode, "'StartCode' <= ? and 'EndCode' >= ? and 'RemoteArea' = ?", postcode.StartCode, postcode.EndCode, "1")
	Db.Where("StartCode <= ? and EndCode >= ? and RemoteArea=?", postcode.StartCode, postcode.EndCode, remotearea).First(&returnpf)

	//fmt.Printf("sql is %s \n", returnpc)
	return returnpf
}

func WritePostCode(postZones []models.PostZone) {

	fmt.Println("start insert table PostZone : \n")
	for _, postZone := range postZones {
		if postZone.StartCode != 0 || postZone.Supplier != "" {
			Db.FirstOrCreate(&postZone, models.PostZone{Destination: postZone.Destination, StartCode: postZone.StartCode, EndCode: postZone.EndCode, Supplier: postZone.Supplier, SupplierDate: postZone.SupplierDate})

		}
	}

}

func WritePostFee(postFees []models.PostFee) {

	fmt.Println("start insert table PostFee : \n")
	for _, postFee := range postFees {
		if postFee.BasicFee != 0 || postFee.Supplier != "" {
			Db.FirstOrCreate(&postFee, models.PostFee{Destination: postFee.Destination, StartWeight: postFee.StartWeight, EndWeight: postFee.EndWeight, Supplier: postFee.Supplier, SupplierDate: postFee.SupplierDate, BasicFee: postFee.BasicFee})

		}
	}

}

func WriteWylSPUList(wylSPU []service.SPUList) {

	fmt.Println("start insert table of WylSPU: \n")
	for _, SPU := range wylSPU {
		SPU.CreatedAt = time.Now()

		// fmt.Printf("%v", SPU.ImgList)
		for _, img := range SPU.ImgList {
			SPU.Img += " " + img
		}
		Db.Create(&SPU)
	}
}

func WriteWylWarehouseList(whResult *service.WarehouseResult) {

	fmt.Println("start insert table of WylWylWareHouse: \n")
	for _, WareHouse := range whResult.Data {

		fmt.Printf("%v", WareHouse)
		WareHouse.CreatedAt = time.Now()
		Db.Create(&WareHouse)

	}
}

func GetWylSPUList() (SPU []service.SPUList) {
	fmt.Println("start get WylSPU: \n")
	SPU = []service.SPUList{}
	Db.Where("warehouse_code=?", "EWD").Find(&SPU)
	SKU := []service.SKUList{}
	Db.Model(&SPU).Related(&SKU, "spu_list_id").Find(&SKU)

	for _, data := range SKU {
		fmt.Println("WylSKU: %v\n", data)
	}
	return SPU
}
