package Service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	sql "../Database"
	models "../Models"
	"github.com/tealeg/xlsx"
)

func ReadStockExcel() (stock []models.Stock) {
	excelFile := "/users/hgz/downloads/2018-06-05 SKU 库存更新.xlsx"
	excel, err := xlsx.OpenFile(excelFile)
	if err != nil {
		fmt.Printf("Open excel file err: %s \n excel file name : %s", err, excelFile)
	}
	fmt.Println("open excel file : \n", excelFile)

	var stocks []models.Stock
	var count int
	for _, sheet := range excel.Sheets {
		// stocks Stock{}
		for _, row := range sheet.Rows {

			stock := models.Stock{}

			// err := row.ReadStruct(&stock)
			// checkerr(err, "readstruct")
			if row.Cells != nil {
				stock.StockDay = time.Date(2018, 6, 5, 0, 0, 0, 0, time.Local)
				for i, cell := range row.Cells {
					//stock = cell.String()
					log.Printf("ID %d value is : %v\n", i, cell)
					if cell.Value != "" {
						switch {
						case i == 0:
							stock.Category = cell.Value
						case i == 1:
							stock.Sku = cell.Value
						case i == 2:
							stock.ParentTitle = cell.Value
						// case i == 3:
						// 	stock.SaleDays, _ = strconv.Atoi(cell.Value)
						case i == 3:
							stock.SaleNums, _ = strconv.Atoi(cell.Value)

						}

					}

				}
			}
			count++
			stocks = append(stocks, stock)

		}
	}
	fmt.Println("count : \n", count)
	// for i, stock := range stocks {
	// 	if i < 100 {
	// 		fmt.Println("stock : \n", stock)
	// 	}
	// }

	return stocks

}

func WriteJson(ourput interface{}) {
	stockJson, err := os.OpenFile("./stock.json", os.O_WRONLY, 0644)
	checkerr(err, "open json")
	n, _ := stockJson.Seek(0, os.SEEK_END)
	output, err := json.MarshalIndent(ourput, "", "\t\t")
	checkerr(err, "write json")
	stockJson.WriteAt([]byte(output), n)
	defer stockJson.Close()
}

func WriteSoldeazyExcel(stocks []models.Stock, count int) {
	em := sql.GetEmployee()
	fmt.Println("Stock count :\n", count)

	em = append(em, models.Employee{ShortName: "AU"})
	// for _, e := range em {
	// 	fmt.Println("employee :\n", e)
	// }
	var BBem []models.Employee
	var MSem []models.Employee
	var SLem []models.Employee
	for _, e := range em {
		switch {
		case e.Platform == "BB":
			BBem = append(BBem, e)
			e.ShortName = "h" + e.ShortName
			BBem = append(BBem, e)
		case e.Platform == "MS":
			MSem = append(MSem, e)
			e.ShortName = "h" + e.ShortName
			MSem = append(MSem, e)
		case e.Platform == "SL":
			SLem = append(SLem, e)
			e.ShortName = "h" + e.ShortName
			SLem = append(SLem, e)
		default:
			// fmt.Println("employee in switch:\n", e)
			BBem = append(BBem, e)
			MSem = append(MSem, e)
			SLem = append(SLem, e)
			e.ShortName = "h" + e.ShortName
			BBem = append(BBem, e)
			MSem = append(MSem, e)
			SLem = append(SLem, e)
		}
	}
	// for _, e := range BBem {
	// 	fmt.Println("employee to BB", e)
	// }
	// for _, e := range MSem {
	// 	fmt.Println("employee to MS", e)
	// }
	// for _, e := range SLem {
	// 	fmt.Println("employee to SL", e)
	// }
	WriteStockExcel(stocks, count, BBem)
	WriteStockExcel(stocks, count, MSem)
	WriteStockExcel(stocks, count, SLem)
}

func WriteStockExcel(stocks []models.Stock, count int, em []models.Employee) {

	var seStocks []models.SoldeazyStock

	wh := "AU-sale"
	var platform string

	for _, e := range em {
		// fmt.Println("employee name :\n", e)
		var emName string
		if e.ShortName != "AU" {
			emName = e.ShortName + "-AU-"
		} else {
			emName = "AU-"
		}
		if e.Platform != "" {
			platform = e.Platform
		}

		for _, stock := range stocks {
			seStock := models.SoldeazyStock{}
			seStock.SKU = emName + stock.Sku
			seStock.Warehouse = wh
			seStock.Stock_Level = stock.SaleNums
			count++
			seStocks = append(seStocks, seStock)
		}
	}

	// for _, seStock := range seStocks {
	// 	fmt.Println("stock of soldeazy :\n", seStock)
	// }
	fmt.Println("stock count :\n", count)

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	checkerr(err, "write stock excel")
	for i, se := range seStocks {
		row = sheet.AddRow()
		if i == 0 || i == 19991 {
			cell = row.AddCell()
			cell.Value = "stock_model"
			cell = row.AddCell()
			cell.Value = "stock_warehouse"
			cell = row.AddCell()
			cell.Value = "stock_stklevel_normal"

			row = sheet.AddRow()
			cell = row.AddCell()
			cell.Value = "SKU"
			cell = row.AddCell()
			cell.Value = "Warehouse"
			cell = row.AddCell()
			cell.Value = "Stock Level"

			row = sheet.AddRow()
		}

		cell = row.AddCell()
		cell.Value = se.SKU
		cell = row.AddCell()
		cell.Value = se.Warehouse
		cell = row.AddCell()
		cell.Value = strconv.Itoa(se.Stock_Level)
		if i == 19990 {
			err = file.Save(platform + " 20180605-1.xlsx")
			checkerr(err, "save excel file")
			file = xlsx.NewFile()
			sheet, err = file.AddSheet("Sheet1")
			checkerr(err, "write stock excel")
		}

	}
	err = file.Save(platform + " 20180605.xlsx")
	checkerr(err, "save excel file")
}

func InitEmployee() {
	var em []models.Employee
	// em = append(em, Employee{2, "杨晓婷", "yxt", ""})
	// em = append(em, Employee{3, "陈青丽", "cql", ""})
	// em = append(em, Employee{4, "张春燕", "zcy", "MS"})
	// em = append(em, Employee{5, "杨晓敏", "yxm", ""})
	// em = append(em, Employee{6, "张思琪", "zsq", "SL"})
	// em = append(em, Employee{7, "李晓团", "lxt", "SL"})
	// em = append(em, Employee{8, "谭诗慧", "tsh", "SL"})
	// em = append(em, Employee{9, "幸凡", "xxj", ""})
	// em = append(em, Employee{10, "黄中超", "hzc", "BB"})
	// em = append(em, Employee{11, "石求娟", "sqj", ""})
	// em = append(em, Employee{12, "杨晓玲", "yxl", "BB"})
	// em = append(em, Employee{13, "董选集", "dxj", ""})
	// em = append(em, Employee{14, "邓艳婷", "dyt", ""})
	// em = append(em, Employee{15, "谈春晖", "tch", "SL"})
	sql.WriteEmployee(em)
}
