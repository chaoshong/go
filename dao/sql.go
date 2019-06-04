package dao

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"reflect"

	service "github.com/chaoshong/go/Service/Supplier"
	"github.com/chaoshong/go/model"

	"github.com/chaoshong/go/inframe/db"
)

func Column(tableName string) {
	stock := model.Stock{}
	query := "select * from " + tableName + " limit 1"
	fmt.Println("query \n", query)
	columns := db.PostgreDb.NewScope(stock).Fields()
	// checkerr(err1, "get row")
	fmt.Println("get table columns\n", reflect.TypeOf(columns))
	// column, err := row.Columns()
	// checkerr(err, "get fields")
	for column := range columns {
		fmt.Println("field:\n", reflect.TypeOf(column))
	}
}

func WriteStocks(stocks []model.Stock) {
	// for i, stock := range stocks {
	// 	if i < 100 {
	// 		fmt.Println("stock : \n", stock)
	// 	}
	// }

	// stock := Stock{0, "653", "test", "1", 0, 0, time.Now()}
	fmt.Println("start insert table : \n")
	for _, stock := range stocks {
		if stock.Sku != "" || stock.ParentTitle != "" {
			stockDb := db.PostgreDb.Where(&model.Stock{Sku: stock.Sku, SaleNums: stock.SaleNums, StockDay: stock.StockDay}).First(&stock)
			//fmt.Println("find table : \n", stockDb)
			if stockDb != nil {
				db.PostgreDb.Create(&stock)
			}
		}
	}

}

func GetStocks(stockDay time.Time) (stocks []model.Stock, count int) {

	db.PostgreDb.Where(&model.Stock{StockDay: stockDay}).Find(&stocks).Count(&count)

	return stocks, count
}

func GetEmployee() (employee []model.Employee) {

	db.PostgreDb.Find(&employee)

	return employee
}

func WriteEmployee(em []model.Employee) {
	fmt.Println("start insert employee : \n")
	for _, e := range em {

		db.PostgreDb.Create(&e)

	}
}
func WriteOrdersLittleBoss(ordersLbs []model.OrderLittleBoss) {

	fmt.Println("start insert table OrdersLittleBoss : \n")
	for _, ordersLb := range ordersLbs {
		if ordersLb.OrderId != "" || ordersLb.ListingSKu != "" || ordersLb.Quantity != 0 {

			db.PostgreDb.FirstOrCreate(&ordersLb, model.OrderLittleBoss{ConsigneeName: ordersLb.ConsigneeName, OrderId: ordersLb.OrderId, ListingSKu: ordersLb.ListingSKu, Quantity: ordersLb.Quantity})

		}
	}

}

func GetOrderLittleBoss(orderHandled string) (orderlb []model.OrderLittleBoss, count int) {
	db.PostgreDb.Where(&model.OrderLittleBoss{OrderHandled: orderHandled}).Find(&orderlb).Count(&count)

	return orderlb, count
}
func GetOrderBalance(orderHandled string) (orderbln []model.OrderBalance, count int) {
	db.PostgreDb.Where(&model.OrderBalance{OrderHandled: orderHandled}).Find(&orderbln).Count(&count)

	return orderbln, count
}
func UpdateBalanceHandle(orderbln []model.OrderBalance) {
	db.PostgreDb.Model(&orderbln).Update("OrderHandled", "1")
}

func WriteOrdersIsale(ordersIsale []model.OrderIsale) {

	fmt.Println("start insert table OrderIsale : \n")
	for _, orderIsale := range ordersIsale {
		if orderIsale.ListingSKu != "" || orderIsale.Quantity != 0 {
			//orderIsale := Db.Where(&Stock{Sku: stock.Sku, SaleNums: stock.SaleNums, StockDay: stock.StockDay}).First(&stock)
			//fmt.Println("find table : \n", stockDb)
			//if stockDb != nil {
			db.PostgreDb.FirstOrCreate(&orderIsale, model.OrderIsale{ConsigneeName: orderIsale.ConsigneeName, OrderId: orderIsale.OrderId, ListingSKu: orderIsale.ListingSKu, Quantity: orderIsale.Quantity})
			//}
		}
	}

}
func WriteOrdersBalance(ordersBalance []model.OrderBalance) {

	fmt.Println("start insert table OrderBalance : \n")
	for _, value := range ordersBalance {
		if value.OrderId != "" || value.SalesSku != "" || value.Qty != 0 {
			//orderIsale := Db.Where(&Stock{Sku: stock.Sku, SaleNums: stock.SaleNums, StockDay: stock.StockDay}).First(&stock)
			//fmt.Println("find table : \n", stockDb)
			//if stockDb != nil {
			db.PostgreDb.FirstOrCreate(&value, model.OrderBalance{FullName: value.FullName, OrderId: value.OrderId, SalesSku: value.SalesSku, Qty: value.Qty})
			//}
		}
	}

}

func GetCostOrder(order model.OrderBalance) (rtnOrder model.OrderBalance) {
	fmt.Println("start get table OrderBalance by getcostorder : \n")
	stock := model.Stock{}
	product := model.ProductCommon{}
	startcode, _ := strconv.Atoi(strings.TrimSpace(order.Postcode))
	//fmt.Printf("get cost order startcode is %d, order is %s ,err is %s \n", startcode, order.Postcode, err)
	endcode := startcode
	postzone := model.PostZone{}
	postfee := model.PostFee{}
	db.PostgreDb.Where(&model.Stock{Sku: order.ProductSKU}).Find(&stock)
	db.PostgreDb.Where(&model.ProductCommon{SKU: order.ProductSKU}).Find(&product)
	//fmt.Printf("get cost order stock num is %d, stock date is %s ,product price is %d \n", stock.SaleNums, stock.StockDay, product.ProductPrice)
	postzone = GetPostCodeDest(model.PostZone{StartCode: startcode, EndCode: endcode}, "0")

	if postzone.Destination == "" {
		postzone = GetPostCodeDest(model.PostZone{StartCode: startcode, EndCode: endcode}, "1")
	}
	if postzone.RemoteArea == "" {
		postzone.RemoteArea = "0"
	}

	weight := product.Weight * float64(order.ProductQty)
	db.PostgreDb.Where("\"StartWeight\"<= ? and \"EndWeight\">=? and \"RemoteArea\" =?", weight, weight, postzone.RemoteArea).First(&postfee)
	//fmt.Printf("get cost order postzone is %s,  \n post fee is %s \n", postzone, postfee)
	rtnOrder = model.OrderBalance{}
	rtnOrder.ProductStockDate = stock.StockDay
	rtnOrder.ProductStockNum = stock.SaleNums
	rtnOrder.CostPrice = product.ProductPrice * float64(order.ProductQty)
	rtnOrder.ShippingCostBySystem = postfee.BasicFee + postfee.PerKgFee*weight/1000
	//fmt.Printf("get cost order ,order is %s \n", rtnOrder)
	return rtnOrder
}

func GetPostCodeDest(postcode model.PostZone, remotearea string) (returnpc model.PostZone) {
	//fmt.Printf("postcode is %s \n", &postcode)

	//Db.Find(&postcode, "'StartCode' <= ? and 'EndCode' >= ? and 'RemoteArea' = ?", postcode.StartCode, postcode.EndCode, "1")
	db.PostgreDb.Where("StartCode <= ? and EndCode >= ? and RemoteArea=?", postcode.StartCode, postcode.EndCode, remotearea).First(&returnpc)

	//fmt.Printf("sql is %s \n", returnpc)
	return returnpc
}

func GetPostFee(postcode model.PostZone, remotearea string) (returnpf model.PostFee) {
	//fmt.Printf("postcode is %s \n", &postcode)

	//Db.Find(&postcode, "'StartCode' <= ? and 'EndCode' >= ? and 'RemoteArea' = ?", postcode.StartCode, postcode.EndCode, "1")
	db.PostgreDb.Where("StartCode <= ? and EndCode >= ? and RemoteArea=?", postcode.StartCode, postcode.EndCode, remotearea).First(&returnpf)

	//fmt.Printf("sql is %s \n", returnpc)
	return returnpf
}

func WritePostCode(postZones []model.PostZone) {

	fmt.Println("start insert table PostZone : \n")
	for _, postZone := range postZones {
		if postZone.StartCode != 0 || postZone.Supplier != "" {
			db.PostgreDb.FirstOrCreate(&postZone, model.PostZone{Destination: postZone.Destination, StartCode: postZone.StartCode, EndCode: postZone.EndCode, Supplier: postZone.Supplier, SupplierDate: postZone.SupplierDate})

		}
	}

}

func WritePostFee(postFees []model.PostFee) {

	fmt.Println("start insert table PostFee : \n")
	for _, postFee := range postFees {
		if postFee.BasicFee != 0 || postFee.Supplier != "" {
			db.PostgreDb.FirstOrCreate(&postFee, model.PostFee{Destination: postFee.Destination, StartWeight: postFee.StartWeight, EndWeight: postFee.EndWeight, Supplier: postFee.Supplier, SupplierDate: postFee.SupplierDate, BasicFee: postFee.BasicFee})

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
		db.PostgreDb.Create(&SPU)
	}
}

func WriteWylWarehouseList(whResult *service.WarehouseResult) {

	fmt.Println("start insert table of WylWylWareHouse: \n")
	for _, WareHouse := range whResult.Data {

		fmt.Printf("%v", WareHouse)
		WareHouse.CreatedAt = time.Now()
		db.PostgreDb.Create(&WareHouse)

	}
}

func GetWylSPUList() (SPU []service.SPUList) {
	fmt.Println("start get WylSPU: \n")
	SPU = []service.SPUList{}
	db.PostgreDb.Where("warehouse_code=?", "EWD").Find(&SPU)
	SKU := []service.SKUList{}
	db.PostgreDb.Model(&SPU).Related(&SKU, "spu_list_id").Find(&SKU)

	for _, data := range SKU {
		fmt.Println("WylSKU: %v\n", data)
	}
	return SPU
}
