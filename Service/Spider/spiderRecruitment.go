package service

import (
	"fmt"

	"github.com/gocolly/colly"
)

func GetRecruitment() {
	c := colly.NewCollector()
	//colly.AllowedDomains("go-colly.org"))
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.Visit("http://51job.com/")
}
