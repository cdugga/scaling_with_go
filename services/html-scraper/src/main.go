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

	c.OnHTML("td[onclick].els", func(e *colly.HTMLElement) {

		exists := e.Attr("class") == "els"

		if  exists {
			fmt.Println("table..." , e.Text)
		}

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://www.bridgewebs.com/cgi-bin/bwoo/bw.cgi?club=bridgeclublive&pid=display_past")
}
