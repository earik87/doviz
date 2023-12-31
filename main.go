package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func getUsdTryParity() {
	c := colly.NewCollector()

	c.OnHTML("div.text-xl.font-semibold.text-white", func(e *colly.HTMLElement) {
		dolarUsdParite := e.Text
		fmt.Println("USD/TRY paritesi:", dolarUsdParite)
	})

	c.Visit("https://kur.doviz.com/serbest-piyasa/amerikan-dolari")
}

func main() {
	getUsdTryParity()
}
