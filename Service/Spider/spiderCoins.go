package service

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func queryExampleTwo() {
	doc, err := goquery.NewDocument("https://www.feixiaohao.com/")
	if err != nil {
		log.Fatal(err)
	}
	var allData []CoinInfo
	doc.Find("table tbody tr").Each(func(i int, selector *goquery.Selection) {
		var data CoinInfo
		Rank := selector.Find("td").Eq(0).Text()
		CoinName := strings.TrimSpace(selector.Find("td").Eq(1).Text())
		CurrentCount := selector.Find("td").Eq(2).Text()
		CurrentPrice := selector.Find("td").Eq(3).Text()
		CurrentMark := selector.Find("td").Eq(4).Text()
		Count := selector.Find("td").Eq(5).Text()
		Change := strings.TrimSpace(selector.Find("td").Eq(6).Text())
		data = CoinInfo{
			Rank:         Rank,
			Name:         CoinName,
			CurrentCount: CurrentCount,
			CurrentPrice: CurrentPrice,
			CurrentMark:  CurrentMark,
			Count:        Count,
			Change:       Change,
		}
		allData = append(allData, data)
	})
	Data, _ := json.MarshalIndent(allData, "", "  ")
	fmt.Println(string(Data))
}
