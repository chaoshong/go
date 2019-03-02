package service

type CoinInfo struct {
	Rank         string `json:"rank"`
	Name         string `json:"name"`
	CurrentCount string `json:"current_count"`
	CurrentPrice string `json:"current_price"`
	CurrentMark  string `json:"Current_mark"`
	Count        string `json:"count"`
	Change       string `json:"change"`
}
