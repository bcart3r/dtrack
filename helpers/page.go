package helpers

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const BaseLink = "https://www.sports-reference.com/cfb/"
const SearchLink = BaseLink + "search/search.fcgi?hint=&search="

func GetPage(link string) *goquery.Document {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("unable to fetch website")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}
