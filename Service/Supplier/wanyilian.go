package service

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
)

const Url = "http://openapi.winit.com.cn/cedpopenapi/service"
const TestUrl = "http://openapi.uat1.winit.com.cn/cedpopenapi/service"
const Token = "ed08b6f3-2597-41d2-9cbf-9e536b8cf545"

var wylInArg = WylInArg{
	SKU:               "",
	SPU:               "",
	CategoryID:        "",
	IsHavingInventory: "",
	Keywords:          "",
	PageParams: PageParams{
		PageNo:     1,
		PageSize:   200,
		TotalCount: 0,
	},
	UDefinedCategoryID: "",
	UserDefinedCode:    "",
	WarehouseCode:      "US0001",
	WarehouseName:      "",
}

var wylInJson = WylInJSON{
	Action:     "wanyilian.supplier.spu.querySPUList",
	AppKey:     "yabu007",
	Data:       wylInArg,
	Format:     "json",
	Language:   "zh_CN",
	Platform:   "SELLERERP",
	Sign:       "00000000000000000000000000000000",
	SignMethod: "md5",
	Timestamp:  "2018-08-23 14:59:03",
	//time.Now().Format("2006-01-02 15:04:05"),
	Version: "1.0",
}

const templ = `<!DOCTYPE html>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html;charset=utf-8" />
        <title>goEcommerce</title>
    </head>
    <body>
        <h1>产品数量 ：{{.PageParams.TotalCount}} </h1>
    </body>
</html>`

/*取得wyl的产品信息*/
func GetProductFromWyl(wylInArg WylInArg) (*WylResult, error) {
	// q := url.QueryEscape(strings.Join(terms, " "))
	/* 	wylInJson.Data.PageParams = page
	   	wylInJson.Data.WarehouseCode = "EWD" */
	wylInJson.Data = wylInArg

	wylInJson.Data.UpdateStartDate = "2018-08-23 14:59:03"
	wylInJson.Data.UpdateEndDate = "2018-09-23 14:59:03"
	wylInJson.Sign = getSign()
	fmt.Printf("get product from wyl pagesize is %v pageNo is %v code is %v \n",
		wylInJson.Data.PageParams.PageSize, wylInJson.Data.PageParams.PageNo, wylInJson.Data.WarehouseCode)

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(wylInJson)

	fmt.Printf("get product from wyl send json is %s\n", b)
	resp, err := http.Post(Url, "application/json; charset=utf-8", b)

	if err != nil {
		return nil, err
	}

	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("get product from wyl body is %s\n", string(body))
	fmt.Printf("get product from wyl status is %s\n", resp.Status)
	fmt.Printf("get product from wyl body is %v\n", resp.Body)

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {

		return nil, fmt.Errorf("get product from wyl error: %s" + resp.Status)
	}
	var wylResult WylResult
	if err := json.NewDecoder(resp.Body).Decode(&wylResult); err != nil {
		fmt.Printf("get product from wyl body msg is %v, code is %v, err is %v \n",
			wylResult.Msg, wylResult.Code, err)

		return nil, err
	}

	fmt.Printf("get product from wyl body totalCount is %v pagesize is %v pageno is %v\n",
		wylResult.Data.PageParams.TotalCount, wylResult.Data.PageParams.PageSize, wylResult.Data.PageParams.PageNo)

	fmt.Printf("get product from wyl body msg is %v, code is %v, wylResule data is %v\n",
		wylResult.Msg, wylResult.Code, "wylResult.Data")

	//ParseWylResult(&wylResult)
	return &wylResult, nil
}

/*取得wyl的仓库信息*/
func GetWarehouseFromWyl() (*WarehouseResult, error) {

	//查询仓库
	wylInJson.Action = "wanyilian.platform.queryWarehouse"
	wylInJson.Sign = getSign()
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(wylInJson)

	fmt.Printf("get warehouse from wyl send json is %s\n", b)
	resp, err := http.Post(Url, "application/json; charset=utf-8", b)

	if err != nil {
		return nil, err
	}

	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("get warehouse from wyl body is %s\n", string(body))
	fmt.Printf("get warehouse from wyl status is %s\n", resp.Status)

	fmt.Printf("get warehouse from wyl body is %v\n", resp.Body)

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {

		return nil, fmt.Errorf("get warehouse from wyl error: %s" + resp.Status)
	}
	var whResult WarehouseResult
	if err := json.NewDecoder(resp.Body).Decode(&whResult); err != nil {
		fmt.Printf("get warehouse from wyl body msg is %v, code is %v, err is %v \n",
			whResult.Msg, whResult.Code, err)

		return nil, err
	}

	fmt.Printf("get warehouse from wyl body msg is %v, code is %v, wylResule data is %v\n",
		whResult.Msg, whResult.Code, whResult.Data)

	ParseWylWarehouseResult(&whResult)
	return &whResult, nil
}

/*wyl加密签名方法*/
func getSign() string {
	w := wylInJson
	//data, _ := json.Marshal(w.Data)
	sign := Token + "action" + w.Action + "app_key" + w.AppKey + "data" + getSignData(WylInArg(w.Data)) + "format" + w.Format +
		"platform" + "sign_method" + w.SignMethod + "timestamp" + "version" + w.Version + Token
	//fmt.Printf("sign string is %s\n", sign)
	writer := md5.New()
	io.WriteString(writer, sign)
	sign = fmt.Sprintf("%x", writer.Sum(nil))
	//fmt.Printf("sign code is %s\n", sign)
	return strings.ToUpper(sign)
}

/*wyl 数据报文生成顺序*/
func getSignData(wd WylInArg) string {
	return "{\"SKU\":\"" + wd.SKU + "\",\"SPU\":\"" + wd.SPU + "\",\"categoryID\":\"" + wd.CategoryID +
		"\",\"isHavingInventory\":\"" + wd.IsHavingInventory + "\",\"keywords\":\"" + wd.Keywords +
		"\",\"pageParams\":{\"pageNo\":" + strconv.Itoa(wd.PageParams.PageNo) + ",\"pageSize\":" +
		strconv.Itoa(wd.PageParams.PageSize) + ",\"totalCount\":" + strconv.Itoa(wd.PageParams.TotalCount) +
		"},\"uDefinedCategoryID\":\"" + wd.UDefinedCategoryID +
		"\",\"updateEndDate\":\"" + wd.UpdateEndDate + "\",\"updateStartDate\":\"" + wd.UpdateStartDate +
		"\",\"userDefinedCode\":\"" + wd.UserDefinedCode + "\",\"warehouseCode\":\"" + wd.WarehouseCode +
		"\",\"warehouseName\":\"" + wd.WarehouseName + "\"}"
}
func ParseWylResult(wr *WylResult) {
	//pNum := PageParams(wr.PageParams)
	//fmt.Printf("%d total %d size %d No \n", pNum.TotalCount, pNum.PageSize, pNum.PageNo)

	SPU := []SPUList(wr.Data.SPUList)
	for i, va := range SPU {
		fmt.Printf(" %d %s %s %s\n", i, va.SPU, va.UserDefinedCode, va.ChineseName)
		for j, sku := range va.SKUList {
			fmt.Printf("   %d %s %s %d\n", j, sku.SKU, sku.RandomSKU, sku.SupplyInventory)
		}
	}
}
func ParseWylWarehouseResult(whResult *WarehouseResult) {

	WH := []WarehouseList(whResult.Data)
	for i, wh := range WH {
		fmt.Printf(" %d %s %s %s\n", i, wh.WarehouseCode, wh.WarehouseName, wh.WarehouseAddress)
	}
}

func WriteSPUListExcel(SPU []SPUList) {

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("Sheet1")
	//checkerr(err, "write SPU excel")
	fmt.Println("start write SPU excel")
	for k, spu := range SPU {
		row = sheet.AddRow()
		/* 	v := reflect.ValueOf(spu)
		count := v.NumField() */

		SkuList := spu.SKUList
		if k < 5 {
			fmt.Println("SKUList: %v", SkuList)
		}
		for _, sku := range SkuList {
			vSku := reflect.ValueOf(sku)
			countSKU := vSku.NumField()

			for i := 0; i < countSKU; i++ {
				cell = row.AddCell()
				f := vSku.Field(i)
				cell.Value = f.String()
			}
			/* for j := 0; j < count; j++ {
				cell = row.AddCell()
				f := v.Field(j)
				cell.Value = f.String()
			} */

		}
	}
	_ = file.Save(" wyl product.xlsx")
	//service.checkerr(err, "save excel file")
}
