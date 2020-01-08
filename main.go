package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"

	"github.com/wagnerfonseca/scraper-fundamentus/model"
)

func main() {

	fileName := "fundamentus.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fileName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(header())

	// Instantiate default collector
	c := colly.NewCollector()

	// Create another collector to scrape symbols details
	detailCollector := c.Clone()

	m := make(map[string]map[string]string)
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

			symbol := &model.Symbol{}

			_, ok := m[s]
			if ok {
				for k, v := range m[s] {
					model.Build(symbol, k, v)
				}
				writer.Write(buildRow(symbol))
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

		if e.ChildText("td:nth-child(5)") != "" {
			add(m, symbol, strings.Replace(e.ChildText("td:nth-child(5)"), "?", "", -1), e.ChildText("td:nth-child(6)"))
		}
	})

	c.Visit("https://www.fundamentus.com.br/detalhes.php")

	log.Printf("Scraping finished, check file %q for results\n", fileName)

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

// buildRow build row for csv file
func buildRow(symbol *model.Symbol) []string {
	return []string{
		symbol.Papel,
		symbol.Cotacao,
		symbol.DataUtlCot,
		symbol.Empresa,
		symbol.Setor,
		symbol.Subsetor,
		symbol.VolMed,
		symbol.ValorMercado,
		symbol.NroAcoes,
		symbol.PL,
		symbol.PVP,
		symbol.VPA,
		symbol.PEBIT,
		symbol.Ativo,
		symbol.PAtivCircLiq,
		symbol.PAtivos,
		symbol.PatrimLiq,
		symbol.ReceitaLiquida,
		symbol.LucroLiquido,
		symbol.EBIT,
		symbol.EVEBITDA,
		symbol.EVEBIT,
		symbol.MargLiquida,
		symbol.MargBruta,
		symbol.MargEBIT,
		symbol.PCapGiro,
		symbol.DivYield,
		symbol.DividaLiquida,
		symbol.DividaBruta,
		symbol.ROIC,
		symbol.ROE,
		symbol.LiquidezCorrente,
		symbol.DivBrPatrim,
		symbol.GiroAtivos,
	}
}

// header build the header to csv file
func header() []string {
	return []string{
		"Papel",
		"Cotacao",
		"Data Utl Cot.",
		"Empresa",
		"Setor",
		"Subsetor",
		"Vol. Med.",
		"Valor Mercado",
		"Nro. Acoes",
		"P/L",
		"P/VP",
		"VPA",
		"P/EBIT",
		"Ativo",
		"P/Ativ Circ Liq",
		"P/Ativos",
		"Patrim Liquido",
		"Receita Liquida",
		"Lucro Liquido",
		"EBIT",
		"EVEBITDA",
		"EVEBIT",
		"Marg. Liquida",
		"Marg. Bruta",
		"Marg. EBIT",
		"P/CapGiro",
		"Div. Yield",
		"Divida Liquida",
		"Divida Bruta",
		"ROIC",
		"ROE",
		"Liquidez Corrente",
		"Div. Br Patrim",
		"Giro Ativos",
	}
}
