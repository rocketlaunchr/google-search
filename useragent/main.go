package main

import (
	"bytes"
	"go/format"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {

	var userAgent []string
	collector := colly.NewCollector(
		colly.AllowedDomains("whatismybrowser.com", "www.whatismybrowser.com"),
	)

	collector.OnHTML("table > tbody > tr", func(h *colly.HTMLElement) {
		h.ForEach("td:nth-child(2) > ul", func(_ int, el *colly.HTMLElement) {
			h.ForEach("li", func(_ int, li *colly.HTMLElement) {
				userAgent = append(userAgent, strings.TrimSpace(li.ChildText("span")))
			})
		})
	})

	collector.Visit("https://www.whatismybrowser.com/guides/the-latest-user-agent/chrome")

	if len(userAgent[0]) > 0 {

		tmp := template.Must(template.New("const").Parse(`	// This file is programmatically generated
	// Do not manualy update
	// go run useragent/main.go
	
	package googlesearch

const defaultAgent			 = 				"{{.}}"
`))

		f, err := os.Create("useragent.go")
		if err != nil {
			log.Panicln(err.Error())
		}
		defer f.Close()

		var tpl bytes.Buffer

		err = tmp.Execute(&tpl, userAgent[0])
		if err != nil {
			log.Panicln(err.Error())
		}

		formatted, err := format.Source(tpl.Bytes())
		if err != nil {
			log.Fatal(err)
		}
		f.Write(formatted)
	}
}
