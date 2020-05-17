package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func createDOM(response *http.Response) *goquery.Document {
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}
	return document
}

func parseDOM(document *goquery.Document) {
	document.Find(".we-customer-review.lockup.ember-view p").Each(processElement)
}

func processElement(index int, element *goquery.Selection) {
	log.Println(element.Text())
}