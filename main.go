package main

import (
	"log"
	"net/http"
	"time"

	api "./Api"
	sql "./Database"
	models "./Models"
	service "./Service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// server := http.Server{
	// 	Addr: "127.0.0.1:8080",
	// }
	// http.HandleFunc("/stock/", stock)
	// server.ListenAndServe()
	// Init()

	// initServer()

	sql.Init()

	router := mux.NewRouter()
	router.HandleFunc("/stock", api.ListStock).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}), handlers.AllowedOrigins([]string{"*"}))(router)))
	// Db.Create(stocks)
	//WriteStocks(stocks)

	updateStock()
	//InitPostCodeZone()
	// updateOrder("/users/hgz/downloads/orderlittleboss.xlsx")
	//updateProduct()
	// api.ListStock(new http.ResponseWriter,new http.Request)
}

func updateStock() {
	var stocks []models.Stock
	var count int
	stocks = service.ReadStockExcel()
	sql.WriteStocks(stocks)
	stocks, count = sql.GetStocks(time.Date(2018, 6, 5, 0, 0, 0, 0, time.Local))

	service.WriteSoldeazyExcel(stocks, count)
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
