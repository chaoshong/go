package service

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	sql "github.com/chaoshong/go/dao"
	models "github.com/chaoshong/go/model"
	"github.com/tealeg/xlsx"

	"github.com/chaoshong/go/inframe/db"
)

func ReadOrderExcel(filePath string) {

	excelFile := "/users/hgz/downloads/orderlittleboss.xlsx"
	excelFile = filePath
	excel, err := xlsx.OpenFile(excelFile)
	if err != nil {
		fmt.Printf("Open excel file err: %s \n excel file name : %s", err, excelFile)
	}
	fmt.Println("open excel file : \n", excelFile)

	var count int
	var columns = new(Columns)
	var orderLbs []models.OrderLittleBoss
	var order models.OrderLittleBoss
	rows, err := db.PostgreDb.Find(&order).Rows()
	checkerr(err, "get table rows")
	columns.tableColumns, err = rows.Columns()
	for _, column := range columns.tableColumns {
		fmt.Println(" table column is \n", column)
	}
	checkerr(err, "get table columns name")
	for _, sheet := range excel.Sheets {
		fmt.Println("sheet name is : \n", sheet.Name)

		columns.xlsxColumns = sheet.Rows[0].Cells
		columns.parseColumns()
		for i, rows := range sheet.Rows {
			if i == 0 {
				continue
			}

			order = models.OrderLittleBoss{}
			v := reflect.ValueOf(&order).Elem()
			t := reflect.TypeOf(&order).Elem()
			for key, value := range columns.useColumns {
				//fmt.Printf("key is %d and value is %s : \n", key, value)

				for i := 0; i < t.NumField(); i++ {
					if value != nil {

						tag := t.Field(i).Tag.Get("gorm")
						tag1 := strings.Split(tag, ":")

						//fmt.Println("order  : \n", i, t.Field(i).Name, t.Field(i).Tag, tag, tag1)
						if len(tag1) >= 2 && tag1[1] != "" {
							// if tag1[1] == "商品主图" {
							// 	continue
							// }
							if strings.ToLower(tag1[1]) == strings.ToLower(value[0]) {

								if v.Field(i).Kind() == reflect.String {
									v.Field(i).SetString(rows.Cells[key].String())
								} else if v.Field(i).Kind() == reflect.Float64 {
									fv, _ := rows.Cells[key].Float()
									v.Field(i).SetFloat(fv)
								} else if v.Field(i).Kind() == reflect.Int {
									iv, _ := rows.Cells[key].Int64()
									v.Field(i).SetInt(iv)
								} else if v.Field(i).Kind() == reflect.Struct {
									//fmt.Printf("date time start \n")
									tv, _ := time.Parse("2006-01-02 15:04:05", rows.Cells[key].String())

									v.Field(i).Set(reflect.ValueOf(tv))
									//fmt.Printf("date time is %f ，value is %s , order value is %s  \n", tv, reflect.ValueOf(tv), v.Field(i))

								}
							}
						}
					}

				}

			}
			//fmt.Println("order is  : \n", order)

			orderLbs = append(orderLbs, order)
			count++
		}
		sql.WriteOrdersLittleBoss(orderLbs)
	}
	fmt.Println("count : \n", count)

}

type Columns struct {
	xlsxColumns  []*xlsx.Cell
	tableColumns []string
	useColumns   [][]string
}

func (c *Columns) parseColumns() {

	c.useColumns = make([][]string, len(c.xlsxColumns))
	for key, value := range c.xlsxColumns {
		// fmt.Printf("excel key is %s and value is %s \n", key, value)
		column := value.String()
		columnName := strings.Split(column, "|")
		var status = false
		for _, value := range c.tableColumns {

			if strings.ToLower(columnName[0]) == strings.ToLower(value) {

				c.useColumns[key] = columnName
				status = true
				//fmt.Printf("columns userc key is %d, value is %s and coValue is %s, value is %s \n", key, c.useColumns[key], columnName[0], value)
			}
			// if key == 2 {
			//fmt.Printf("coValue is %s, value is %s \n", strings.ToLower(columnName[0]), strings.ToLower(value))

			// }

		}
		if status == false {
			//fmt.Printf("columns userc key is %d, value is %s and coValue is %s, value is %s \n", key, c.useColumns[key], columnName[0], value)
		}
	}

	fmt.Println("Column use is :\n", c.useColumns)
}
func WriteOrderBalance(ordersBl []models.OrderLittleBoss, count int) {
	var ordersBalance []models.OrderBalance
	var orderBalance models.OrderBalance

	tbln := reflect.TypeOf(&orderBalance).Elem()
	vbln := reflect.ValueOf(&orderBalance).Elem()

	for _, value := range ordersBl {
		orderBalance = models.OrderBalance{}
		tLb := reflect.TypeOf(&value).Elem()
		vLb := reflect.ValueOf(&value).Elem()
		for i := 0; i < vbln.NumField(); i++ {
			tin := tbln.Field(i).Tag.Get("order")
			for j := 0; j < tLb.NumField(); j++ {
				if tin == tLb.Field(j).Tag.Get("order") && tin != "" {

					if vLb.Field(j).Kind() == reflect.Int {
						vbln.Field(i).SetInt(vLb.Field(j).Int())
						//fmt.Printf("order little boss value is %s tag is %s , order balance value is %s tag is %s \n", vLb.Field(j), tLb.Field(j).Tag.Get("order"), vbln.Field(i), tin)

					} else if vLb.Field(j).Kind() == reflect.String {
						vbln.Field(i).SetString(vLb.Field(j).String())
						//fmt.Printf("order little boss value is %s tag is %s , order balance value is %s tag is %s \n", vLb.Field(j), tLb.Field(j).Tag.Get("order"), vbln.Field(i), tin)

					} else if vLb.Field(j).Kind() == reflect.Float64 {
						vbln.Field(i).SetFloat(vLb.Field(j).Float())
						//fmt.Printf("order little boss value is %s tag is %s , order balance value is %s tag is %s \n", vLb.Field(j), tLb.Field(j).Tag.Get("order"), vbln.Field(i), tin)

					} else if vLb.Field(j).Kind() == reflect.Struct {
						//fmt.Printf("start payment date \n")
						st := vLb.Field(j).Interface()
						//tv, _ := time.Parse("2006-01-02 15:04:05", st)
						vbln.Field(i).Set(reflect.ValueOf(st))
						//fmt.Printf("date time is %s ,value is %s , order value is %s  \n", st, vLb.Field(j), vbln.Field(i))

					}

					//fmt.Printf("order little boss value is %s name is %s , order balance value is %s name is %s , i is %d , j is %d\n", vLb.Field(j), tLb.Field(j).Name, vbln.Field(i), tbln.Field(i).Name, i, j)

				}
			}

			//if tin == "OrderDate"||tin=="Email"||tin=="Company"||tin=="PackagingGroupTag"||tin=="OrderItemTitle"||tin=="Unitcost"
		}
		//fmt.Println("orderbalance is \n", orderBalance)
		//fmt.Println("order little boss is \n", value)
		orderBalance = ParseOrder(orderBalance)
		ordersBalance = append(ordersBalance, orderBalance)
	}

	sql.WriteOrdersBalance(ordersBalance)
}

func WriteOrderIsale(ordersBln []models.OrderBalance, count int) {
	fmt.Printf("start Write isale order \n")
	var ordersIsale []models.OrderIsale
	var orderIsale models.OrderIsale
	tIsale := reflect.TypeOf(&orderIsale).Elem()
	vIsale := reflect.ValueOf(&orderIsale).Elem()
	fmt.Println("order balance is \n", ordersBln)
	for _, value := range ordersBln {
		if value.Supplier != "iSale" {
			continue
		}
		orderIsale = models.OrderIsale{Source: "DS-eBay"}
		tbln := reflect.TypeOf(&value).Elem()
		vbln := reflect.ValueOf(&value).Elem()
		for i := 0; i < vIsale.NumField(); i++ {
			tin := tIsale.Field(i).Tag.Get("order")
			for j := 0; j < tbln.NumField(); j++ {
				if tin == tbln.Field(j).Tag.Get("order") && tin != "" {

					//fmt.Printf("parse order isale field name is %s , type is %s , order bln type is %s \n", tin, vIsale.Field(i).Kind(), vbln.Field(j).Kind())
					if vIsale.Field(i).Kind() == reflect.String {
						vIsale.Field(i).SetString(vbln.Field(j).String())
					} else if vIsale.Field(i).Kind() == reflect.Int {
						vIsale.Field(i).SetInt(vbln.Field(j).Int())
					} else if vIsale.Field(i).Kind() == reflect.Float64 {
						vIsale.Field(i).SetFloat(vbln.Field(j).Float())
					} else if vIsale.Field(i).Kind() == reflect.Struct {
						//fmt.Printf("start payment date \n")
						st := vbln.Field(j).Interface()
						//tv, _ := time.Parse("2006-01-02 15:04:05", st)
						vIsale.Field(i).Set(reflect.ValueOf(st))
						//fmt.Printf("date time is %s ,value is %s , order value is %s  \n", st, vbln.Field(j), vIsale.Field(i))

					}
				}
			}
		}
		//fmt.Println("order Isale is \n", orderIsale)
		fmt.Println("order balance is \n", value)

		ordersIsale = append(ordersIsale, orderIsale)
	}

	sql.WriteOrdersIsale(ordersIsale)
	WriteIsaleExcel(ordersIsale)
}

func WriteIsaleExcel(ordersIsale []models.OrderIsale) {

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	checkerr(err, "write order isale excel")
	for i, value := range ordersIsale {
		row = sheet.AddRow()
		vo := reflect.ValueOf(&value).Elem()
		to := reflect.TypeOf(&value).Elem()

		for j := 0; j < vo.NumField(); j++ {
			if j == 0 || j == vo.NumField()-1 {
				continue
			}
			if i == 0 {

				tag := strings.Split(to.Field(j).Tag.Get("gorm"), ":")
				cell = row.AddCell()
				cell.Value = tag[len(tag)-1]
			} else {
				cell = row.AddCell()
				if vo.Field(j).Kind() == reflect.Int {
					cell.Value = strconv.FormatInt(vo.Field(j).Int(), 10)
				} else if vo.Field(j).Kind() == reflect.String {
					cell.Value = vo.Field(j).String()
				}

			}

		}
		if i == 0 {
			continue
		}
		for k, n := range sheet.Rows[0].Cells {

			cells := row.Cells
			if n.Value == "OrderDate" || n.Value == "Email" || n.Value == "Company" || n.Value == "PackagingGroupTag" || n.Value == "OrderItemTitle" || n.Value == "Unitcost" || n.Value == "OrderId" {
				cells[k].Value = ""
			}
			if n.Value == "PostageServiceTag" {
				service := strings.Split(value.SelectedLogistics, "_")
				cells[k].Value = service[len(service)-1]
			}
		}

	}
	err = file.Save("北京德昭圆明 " + time.Now().Format("20060102") + ".xlsx")
	checkerr(err, "save excel file")
}

func ParseOrder(order models.OrderBalance) (rtnOrder models.OrderBalance) {
	//fmt.Printf("start parse order , fieldName is %s , order is %s \n", fieldName, order)

	skus := strings.Split(order.SalesSku, "-")
	sku := skus[len(skus)-1]
	//fmt.Printf("parse order , skus is %s,sku is %s \n", skus, sku)
	qty := order.Qty
	if strings.Contains(sku, "*") {
		tmp := strings.Split(sku, "*")
		sku = strings.TrimSpace(tmp[0])
		tQty, _ := strconv.Atoi(tmp[1])
		qty = tQty * qty
	}
	order.ProductSKU = sku
	order.ProductQty = qty
	order.Supplier = ParseSupplier(sku)
	costorder := sql.GetCostOrder(order)
	order.ProductStockNum = costorder.ProductStockNum
	order.ProductStockDate = costorder.ProductStockDate
	order.CostPrice = costorder.CostPrice
	order.ShippingCostBySystem = costorder.ShippingCostBySystem
	order.OrderHandled = "0"
	//fmt.Printf("finish parse order , sku is %s , qty is %d \n", order.ProductSKU, order.ProductQty)
	return order
}

func ParseSupplier(sku string) (supplier string) {
	_, err := strconv.Atoi(sku)
	if strings.Index(strings.ToLower(sku), "m2c") == 0 {
		supplier = "M2C"
	} else if len(sku) == 10 && err == nil {
		supplier = "iSale"
	} else if strings.Index(strings.ToLower(sku), "w") == 0 {
		supplier = "WYL"
	}
	return supplier
}
