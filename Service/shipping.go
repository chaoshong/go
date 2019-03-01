package Service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	sql "github.com/chaoshong/go/Databases"
	models "github.com/chaoshong/go/Models"
	"github.com/tealeg/xlsx"
)

func InitPostCodeZone() {
	excelFile := "/users/hgz/downloads/FT邮费标准2018.02.07.xlsx"
	excel, err := xlsx.OpenFile(excelFile)
	if err != nil {
		fmt.Printf("Open excel file err: %s \n excel file name : %s", err, excelFile)
	}
	fmt.Println("open excel file : \n", excelFile)

	var count int
	var columns = new(Columns)
	var postzones []models.PostZone
	var postzone models.PostZone
	var postFees []models.PostFee
	var postFee models.PostFee
	rows, err := sql.Db.Find(&postzone).Rows()
	checkerr(err, "get table rows")
	columns.tableColumns, err = rows.Columns()
	for _, column := range columns.tableColumns {
		fmt.Printf(" table column is %s \n", column)
	}
	checkerr(err, "get table columns name")
	for _, sheet := range excel.Sheets {
		fmt.Println("sheet name is : \n", sheet.Name)
		if sheet.Name == "DOMESTIC RATES" {
			var bPostFee bool
			sullierDate := time.Date(2018, 2, 7, 0, 0, 0, 0, time.Local)
			postFee = models.PostFee{Country: "AU", Supplier: "FT", SupplierDate: sullierDate}
			for j, rows := range sheet.Rows {
				fmt.Printf("Start init post Fee and row is %d and rows nil is %s \n", j, rows.Cells[0].Value)
				if j > 10 && rows.Cells[0].Value == "" {
					break
				}
				if rows.Cells[0].Value == "Destination Zone" {
					columns.xlsxColumns = rows.Cells
					fmt.Printf("weight title is %s and row is %d \n", columns.xlsxColumns, j)
					bPostFee = true
					continue
				}
				if bPostFee == true {

					//fmt.Printf("Start init post Fee and row is %d \n", j)
					postFee.Destination = rows.Cells[0].String()
					if strings.Contains(postFee.Destination, "REMOTE AREA") {
						postFee.RemoteArea = "1"
						bPostFee = false
					}
					// for i, _ := range columns.xlsxColumns {
					// 	fmt.Printf("columns name is %d and value is %s \n", i, columns.xlsxColumns[i].String())
					// }
					for i, value := range rows.Cells {
						if i == 0 {
							continue
						}
						if value.String() == "" {
							continue
						}
						fmt.Printf("cell number is %d and value is %s \n", i, value)
						if strings.Contains(columns.xlsxColumns[i].String(), "500g") {
							//fmt.Printf("columns name is 500g and value is %s \n", value)
							postFee.BasicFee, _ = value.Float()
							postFee.PerKgFee = 0.0
							postFee.StartWeight = 0
							postFee.EndWeight = 500
						} else if strings.Contains(columns.xlsxColumns[i].String(), "501g") {
							//fmt.Printf("columns name is 501g and value is %s \n", value.String())
							postFee.BasicFee, _ = value.Float()
							postFee.PerKgFee = 0.0
							postFee.StartWeight = 501
							postFee.EndWeight = 1000
						} else if strings.Contains(columns.xlsxColumns[i].String(), "1.01kg") {
							//fmt.Printf("columns name is 1.01kg and value is %s \n", rows.Cells[i].String())
							postFee.BasicFee, _ = rows.Cells[i].Float()
							postFee.PerKgFee = 0.0
							postFee.StartWeight = 1001
							postFee.EndWeight = 2000
						} else if strings.Contains(columns.xlsxColumns[i].String(), "2.01kg") {
							//fmt.Printf("columns name is 2.01kg and value is %s \n", rows.Cells[i].String())
							postFee.BasicFee, _ = rows.Cells[i].Float()
							postFee.PerKgFee = 0.0
							postFee.StartWeight = 2001
							postFee.EndWeight = 3000
						} else if strings.Contains(columns.xlsxColumns[i].String(), "3.01kg") {
							//fmt.Printf("columns name is 3.01kg and value is %s \n", rows.Cells[i].String())

							postFee.BasicFee, _ = rows.Cells[i].Float()
							postFee.PerKgFee = 0.0
							postFee.StartWeight = 3001
							postFee.EndWeight = 4000
						} else if strings.Contains(columns.xlsxColumns[i].String(), "4.01kg") {
							//fmt.Printf("columns name is 4.01kg and value is %s \n", rows.Cells[i].String())

							postFee.BasicFee, _ = rows.Cells[i].Float()
							postFee.PerKgFee = 0.0
							postFee.StartWeight = 3001
							postFee.EndWeight = 4000

						} else if strings.Contains(columns.xlsxColumns[i].String(), "5.01kg") {
							//fmt.Printf("columns name is 5.01kg and value is %s \n", rows.Cells[i].String())

							postFee.BasicFee, _ = rows.Cells[i].Float()
							postFee.PerKgFee = 0.0
							postFee.StartWeight = 5001
							postFee.EndWeight = 10000
						} else if strings.Contains(columns.xlsxColumns[i].String(), "10.01kg") {
							//fmt.Printf("columns name is 10.01kg and value is %s \n", rows.Cells[i].String())

							postFee.BasicFee, _ = rows.Cells[i].Float()
							postFee.PerKgFee = 0.0
							postFee.StartWeight = 10001
							postFee.EndWeight = 15000
						} else if strings.Contains(columns.xlsxColumns[i].String(), "15.01kg") {
							//fmt.Printf("columns name is 15.01kg and value is %s \n", rows.Cells[i].String())

							postFee.BasicFee, _ = rows.Cells[i].Float()
							postFee.PerKgFee = 0.0
							postFee.StartWeight = 15001
							postFee.EndWeight = 22000
						} else {

							continue
						}
						//fmt.Printf("post Fee is : %s \n", postFee)
						postFees = append(postFees, postFee)
					}
				}

			}
			fmt.Println("postfee is : \n", postFees)
			sql.WritePostFee(postFees)
		}
		if sheet.Name == "POST CODE" {
			continue
			var bRemoteArea bool
			var bNotRA bool
			sullierDate := time.Date(2018, 2, 7, 0, 0, 0, 0, time.Local)
			postzone = models.PostZone{Country: "AU", Supplier: "FT", SupplierDate: sullierDate}
			for _, rows := range sheet.Rows {
				if rows.Cells[0].String() == "Remote Area PostCode List" {
					bRemoteArea = true
					continue
				}
				if strings.Contains(rows.Cells[0].String(), "excluded from the Remote Area") {
					bRemoteArea = false
					bNotRA = true
					continue
				}
				if bRemoteArea == true {
					fmt.Println("Start init postcode zone \n")
					postzone.Destination = rows.Cells[0].String()
					postcode := rows.Cells[1].String()
					//fmt.Printf("Destination is %s and postcode is %s \n", postzone.Destination, postcode)
					postcodes := strings.Split(postcode, ",")
					for _, value := range postcodes {
						pcode := strings.Split(value, "-")
						if len(pcode) > 1 {

							postzone.StartCode, _ = strconv.Atoi(strings.TrimSpace(pcode[0]))
							postzone.EndCode, _ = strconv.Atoi(strings.TrimSpace(pcode[1]))
						} else {
							postzone.StartCode, _ = strconv.Atoi(strings.TrimSpace(pcode[0]))
							postzone.EndCode, _ = strconv.Atoi(strings.TrimSpace(pcode[0]))
						}
						postzone.RemoteArea = "1"
						//fmt.Printf("post zone is : %s \n", postzone)
						postzones = append(postzones, postzone)

					}
					sql.WritePostCode(postzones)

				}
				if bNotRA == true {
					fmt.Println("Start init postcode zone not remote area\n")

					postcode := rows.Cells[0].String()
					//fmt.Printf("not remote area postcode is %s \n", postcode)
					postcodes := strings.Split(postcode, ",")
					for _, value := range postcodes {

						postzone.StartCode, _ = strconv.Atoi(strings.TrimSpace(value))
						postzone.EndCode, _ = strconv.Atoi(strings.TrimSpace(value))
						postzone.RemoteArea = "0"
						postzone.Destination = sql.GetPostCodeDest(postzone, "1").Destination
						//fmt.Printf("not remote area post zone is : %s \n", postzone)
						postzones = append(postzones, postzone)

					}
					sql.WritePostCode(postzones)
				}

				count++
			}
			//fmt.Println("postzones is : \n", postzones)
			fmt.Println("count : \n", count)
		}
	}
	fmt.Println("count : \n", count)
}
