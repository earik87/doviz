package internal

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func GetEurTryParity() float64 {
	c := colly.NewCollector()

	var eurTryParity string
	c.OnHTML("div.text-xl.font-semibold.text-white", func(e *colly.HTMLElement) {
		eurTryParity = e.Text
	})

	err := c.Visit("https://kur.doviz.com/serbest-piyasa/euro")
	eurTryParity = strings.Replace(eurTryParity, ",", ".", -1)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	val, err := strconv.ParseFloat(eurTryParity, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return val
}
