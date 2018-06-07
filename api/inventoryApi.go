package Api

import (
	"encoding/json"
	"net/http"
	"time"

	sql "../Database"
	models "../Models"
)

func ListStock(w http.ResponseWriter, req *http.Request) {
	var stocks []models.Stock
	var date time.Time
	_ = json.NewDecoder(req.Body).Decode(&date)
	stocks, _ = sql.GetStocks(date)
	json.NewEncoder(w).Encode(stocks)
}
