package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	sql "github.com/chaoshong/go/Databases"
	models "github.com/chaoshong/go/Models"
	service "github.com/chaoshong/go/Service"
	serviceS "github.com/chaoshong/go/Service/Supplier"
)

func main() {

	sql.Init()
	//serviceS.GetProductFromWyl()
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/wylview", handleWylview)
	http.HandleFunc("/wyl", handleWyl)
	http.HandleFunc("/wylwarehouse", handleWylWarehouse)
	http.HandleFunc("/wylgetSPU", handleWylSPUList)

	server.ListenAndServe()

	// http.HandleFunc("/stock/", stock)
	// Init()

	// var count int
	// stocks = service.ReadStockExcel()
	// sql.WriteStocks(stocks)
	// stocks, count = sql.GetStocks(time.Date(2018, 6, 5, 0, 0, 0, 0, time.Local))

	// service.WriteSoldeazyExcel(stocks, count)

}

func updateOrder(filePath string) {
	var ordersLs []models.OrderLittleBoss

	var count int
	service.ReadOrderExcel(filePath)
	ordersLs, count = sql.GetOrderLittleBoss("")

	service.WriteOrderBalance(ordersLs, count)
	orderBln, count := sql.GetOrderBalance("0")
	service.WriteOrderIsale(orderBln, count)
	sql.UpdateBalanceHandle(orderBln)
}
func updateProduct() {

	service.ReadExcel()
}

func handleWylview(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("client.html")
	t.Execute(w, nil)

	http.Redirect(w, r, "/wyl", http.StatusOK)
}
func handleWyl(w http.ResponseWriter, r *http.Request) {

	wylInArg := serviceS.WylInArg{}
	wylInArg.PageParams.PageSize = 200
	wylInArg.PageParams.PageNo = 1
	r.ParseForm()
	fmt.Printf("URL Wyl %v  %v  %v\n",
		r.FormValue("pageSize"), r.FormValue("pageNo"), r.FormValue("warehouseCode"))
	wylInArg.PageParams.PageSize, _ = strconv.Atoi(r.FormValue("pageSize"))
	wylInArg.PageParams.PageNo, _ = strconv.Atoi(r.FormValue("pageNo"))
	wylInArg.WarehouseCode = r.FormValue("warehouseCode")

	wylResult, _ := serviceS.GetProductFromWyl(wylInArg)

	fmt.Fprintf(w, "%v\n", wylResult)
	sql.WriteWylSPUList(wylResult.Data.SPUList)
}

func handleWylWarehouse(w http.ResponseWriter, r *http.Request) {

	whResult, _ := serviceS.GetWarehouseFromWyl()

	fmt.Fprintf(w, "%v\n", whResult)
	sql.WriteWylWarehouseList(whResult)
}

func handleWylSPUList(w http.ResponseWriter, r *http.Request) {
	SPU := sql.GetWylSPUList()
	fmt.Fprintf(w, "%v\n", SPU)
	serviceS.WriteSPUListExcel(SPU)
}
