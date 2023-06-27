package main

import (
    "fmt"
    "github.com/gocolly/colly"
)

func scrape() {
    c := colly.NewCollector()

    c.OnHTML("div.text-xl.font-semibold.text-white", func(e *colly.HTMLElement) {
		fmt.Println("Dolar/TRY Paritesi:", e.Text)
	})

    c.Visit("https://kur.doviz.com/serbest-piyasa/amerikan-dolari")
}

func main() {
    scrape()
}
