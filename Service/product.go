package Service

import (
	"fmt"
	"strings"

	"reflect"

	sql "github.com/chaoshong/goecommerce/Databases"
	models "github.com/chaoshong/goecommerce/Models"
	"github.com/tealeg/xlsx"
)

func ReadExcel() {

	excelFile := "/users/hgz/downloads/产品开发目录.xlsx"
	excel, err := xlsx.OpenFile(excelFile)
	if err != nil {
		fmt.Printf("Open excel file err: %s \n excel file name : %s", err, excelFile)
	}
	fmt.Println("open excel file : \n", excelFile)

	var count int
	var columns = new(Columns)
	var productcommon models.ProductCommon
	//tablename := "ProductCommon"
	rows, err := sql.Db.Find(&productcommon).Rows()
	checkerr(err, "get table rows")
	columns.tableColumns, err = rows.Columns()
	for _, column := range columns.tableColumns {
		fmt.Println(" table column is \n", column)
	}
	checkerr(err, "get table columns name")
	for _, sheet := range excel.Sheets {
		//fmt.Println("sheet name is : \n", sheet.Name)
		if sheet.Name == "澳洲赛尔产品" {
			columns.xlsxColumns = sheet.Rows[0].Cells
			columns.parseColumns()
			//tmp := 0
			for i, rows := range sheet.Rows {
				if i == 0 {
					continue
				}

				productcommon = models.ProductCommon{}
				v := reflect.ValueOf(&productcommon).Elem()
				t := reflect.TypeOf(&productcommon).Elem()
				for key, value := range columns.useColumns {
					fmt.Printf("key is %d and value is %s : \n", key, value)

					for i := 0; i < t.NumField(); i++ {
						if value != nil {
							if strings.ToLower(t.Field(i).Name) == strings.ToLower(value[0]) {
								//fmt.Println("productcommon  : \n", i, t.Field(i).Name, v.Field(i).Type(), v.Field(i).Interface())

								//fmt.Println("use Columns  : \n", key, t.Field(i).Name, value, rows.Cells[key].String())
								v.Field(i).SetString(rows.Cells[key].String())
							}
						}

					}

				}
				//fmt.Println("productcommon is  : \n", productcommon)

				sql.Db.Create(&productcommon)
				count++
			}

		}
	}
	fmt.Println("count : \n", count)

}
