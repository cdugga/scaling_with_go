package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/http"
)

func main(){
	ScrapeHTML()
}

func LoadPage() *http.Response {
	url := "https://www.bridgewebs.com/cgi-bin/bwoo/bw.cgi?club=bridgeclublive&pid=display_past"
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	return resp
}

func ScrapeHTML() {

	log.Println("Scrape html")

	c :=  colly.NewCollector( )

	c.OnHTML("table tbody tr", func(e *colly.HTMLElement) {
		e.ForEach("td.els", func(_ int, row *colly.HTMLElement) {

			//if row.Attr("class") == ".elw2ph" || row.Attr("class") == ".els elw2ph" {
			//	fmt.Println("Stuff" ,e.Text)
			fmt.Println("Class..." , e.ChildText(("els")))
			//}

		})
	})

	c.Visit("https://www.bridgewebs.com/cgi-bin/bwoo/bw.cgi?club=bridgeclublive&pid=display_past")
}
