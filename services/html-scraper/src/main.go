package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"net/http"
)

func main(){
	//b := LoadPage().Body
	//
	//r, _ := ioutil.ReadAll(b)

	//fmt.Println(string(r))

	ScrapeHTML()
}

func LoadPage() *http.Response {
	url := "https://www.donedeal.ie/cars-for-sale/honda-accord-2011/27264659"
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	return resp
}

func ScrapeHTML() {

	log.Println("Scrape html")

	c :=  colly.NewCollector()

	c.OnHTML("main script", func(e *colly.HTMLElement) {


		f := e.DOM.Find(".price")



		fmt.Println("dsadasa" , f.Length())

		//fmt.Println("Text" , e.Text)



	})


	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://www.donedeal.ie/cars-for-sale/honda-accord-2011/27264659")
}
//
//func traverseDIVTree(div *colly.HTMLElement) {
//
//	if div.
//
//}