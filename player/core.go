package player

import (
	"dtrack/page"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Player struct {
	name  string
	games []*goquery.Selection
	stats []map[string]string
}

func GetPlayer(fname, lname, team string) string {
	p := page.Get(page.SearchLink + fname + "+" + lname + "&pid=&idx=")
	var doc string

	p.Find(".search-item").Each(func(_ int, s *goquery.Selection) {
		if s.Find(".search-item-league").Text() == team {
			purl := s.Find(".search-item-url").Text()
			link := page.BaseLink + strings.Replace(purl, "/cfb/", "", -1)

			//remove .html from the link for concatenation purposes
			doc = strings.Replace(link, ".html", "", -1)
		}
	})

	return doc
}

func GetGameLog(link, year string) []*goquery.Selection {
	var g []*goquery.Selection

	doc := page.Get(link + "/gamelog/" + year + "/")

	fmt.Println("in game log")
	games := doc.Find("tbody")
	games.Find("tr").Each(func(_ int, s *goquery.Selection) {
		g = append(g, s)
	})

	return g
}
