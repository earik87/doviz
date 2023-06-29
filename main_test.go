package main

import "testing"

func TestScrape(t *testing.T){

	upperLimit := 50
	lowerLimit := 12
	got := scrape()

	if got < upperLimit && got > lowerLimit {
		t.Errorf("Test fail! upperLimit: '%d', lowerLimit: '%d', got: '%d'", upperLimit, lowerLimit, got)
	}

}
