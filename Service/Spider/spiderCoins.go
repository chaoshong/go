package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	goquery "github.com/PuerkitoBio/goquery"
)

func QueryExampleTwo() {
	res, err := http.Get("https://www.feixiaohao.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
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
