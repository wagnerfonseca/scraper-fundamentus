package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"

	"github.com/wagnerfonseca/scraper-fundamentus/model"
)

func main() {

	// Instantiate default collector
	c := colly.NewCollector()

	// Create another collector to scrape symbols details
	detailCollector := c.Clone()

	m := make(map[string]map[string]string)
	// Check if a symbol was verified
	check := make([]string, 0)
	var symbol string = ""

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	// On every a HTML element which has name attribute call callback
	c.OnHTML(`a[href]`, func(e *colly.HTMLElement) {
		// Activate detailCollector if the link contains details
		symbolURL := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Index(symbolURL, "detalhes.php?papel=") != -1 {
			log.Println("==>", symbolURL)
			detailCollector.Visit(symbolURL)
			s := strings.Replace(symbolURL, "https://www.fundamentus.com.br/detalhes.php?papel=", "", -1)

			symbol := &model.Symbol{
				Papel: s,
			}

			if len(check) > 0 {
				for k, v := range m[s] {
					model.Build(symbol, k, v)
					check = check[:0]
				}
			}
			_, ok := m[s]
			if ok {
				for k, v := range m[s] {
					model.Build(symbol, k, v)
				}
			} else {
				check = append(check, s)
			}
		}
	})

	// Extract details of the symbol
	detailCollector.OnHTML("table tr", func(e *colly.HTMLElement) {
		if strings.Contains(e.ChildText("td:first-child"), "Papel") {
			symbol = strings.Replace(e.ChildText("td:nth-child(2)"), "?", "", -1)
		}

		if e.ChildText("td:first-child") != "" {
			add(m, symbol, strings.Replace(e.ChildText("td:first-child"), "?", "", -1), e.ChildText("td:nth-child(2)"))
		}

		if e.ChildText("td:nth-child(3)") != "" {
			add(m, symbol, strings.Replace(e.ChildText("td:nth-child(3)"), "?", "", -1), e.ChildText("td:nth-child(4)"))
		}
	})

	c.Visit("https://www.fundamentus.com.br/detalhes.php")

	// Display collector's statistics
	log.Println(c)
}

// add adding new data
func add(m map[string]map[string]string, symbol, key, value string) {
	mm, ok := m[symbol]
	if !ok {
		mm = make(map[string]string)
		m[symbol] = mm
	}
	mm[key] = value
}
