package googlesearch

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gocolly/colly/v2"
)

// GoogleResult represents a single result from Google Search.
type GoogleResult struct {

	// Rank is the order number of the search result.
	Rank int `json:"rank"`

	// URL of result.
	URL string `json:"url"`

	// Title of result.
	Title string `json:"title"`

	// Description of the result.
	Description string `json:"description"`
}

// GoogleDomains represents localized Google homepages. The 2 letter country code is based on ISO 3166-1 alpha-2.
//
// PR's are welcome.
//
// See: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
var GoogleDomains = map[string]string{
	"us": "https://www.google.com/search?q=",
	"gb": "https://www.google.co.uk/search?q=",
	"ru": "https://www.google.ru/search?q=",
	"fr": "https://www.google.fr/search?q=",
	"au": "https://www.google.com.au/search?q=",
	"nz": "https://www.google.co.nz/search?q=",
}

// SearchOptions modifies how the Search function behaves.
type SearchOptions struct {

	// CountryCode sets the ISO 3166-1 alpha-2 code of the localized Google Search homepage to use.
	// The default is "us", which will return results from https://www.google.com.
	CountryCode string

	// LanguageCode sets the language code.
	// Default: en
	LanguageCode string

	// Limit sets how many results to fetch (at maximum).
	Limit int

	// Start sets from what rank the new result set should return.
	Start int

	// UserAgent sets the UserAgent of the http request.
	// Default: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36"
	UserAgent string
}

// Search returns a list of search results from Google.
func Search(ctx context.Context, searchTerm string, opts ...SearchOptions) ([]GoogleResult, error) {

	c := colly.NewCollector(colly.MaxDepth(0))
	if len(opts) == 0 {
		opts = append(opts, SearchOptions{})
	}

	if opts[0].UserAgent == "" {
		c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36"
	} else {
		c.UserAgent = opts[0].UserAgent
	}

	var lc string
	if opts[0].LanguageCode == "" {
		lc = "en"
	} else {
		lc = opts[0].LanguageCode
	}

	results := []GoogleResult{}
	var rErr error
	rank := 1

	c.OnRequest(func(r *colly.Request) {
		if err := ctx.Err(); err != nil {
			r.Abort()
			rErr = err
			return
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		rErr = err
	})

	c.OnHTML("div.g", func(e *colly.HTMLElement) {

		sel := e.DOM

		for i := range sel.Nodes {
			item := sel.Eq(i)

			rDiv := item.Find("div.r")

			linkHref, _ := rDiv.Find("a").Attr("href")
			linkText := strings.TrimSpace(linkHref)
			titleText := strings.TrimSpace(rDiv.Find("h3").Text())

			sDiv := item.Find("div.s")

			descText := strings.TrimSpace(sDiv.Find("span.st").Text())

			if linkText != "" && linkText != "#" {
				result := GoogleResult{
					Rank:        rank,
					URL:         linkText,
					Title:       titleText,
					Description: descText,
				}
				results = append(results, result)
				rank += 1
			}
		}
	})

	url := url(searchTerm, opts[0].CountryCode, lc, opts[0].Limit, opts[0].Start)
	c.Visit(url)

	if rErr != nil {
		return nil, rErr
	}
	return results, nil
}

func url(searchTerm string, countryCode string, languageCode string, limit int, start int) string {
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	countryCode = strings.ToLower(countryCode)

	var url string

	if googleBase, found := GoogleDomains[countryCode]; found {
		if start == 0 {
			url = fmt.Sprintf("%s%s&hl=%s", googleBase, searchTerm, languageCode)
		} else {
			url = fmt.Sprintf("%s%s&hl=%s&start=%d", googleBase, searchTerm, languageCode, start)
		}
	} else {
		if start == 0 {
			url = fmt.Sprintf("%s%s&hl=%s", GoogleDomains["us"], searchTerm, languageCode)
		} else {
			url = fmt.Sprintf("%s%s&hl=%s&start=%d", GoogleDomains["us"], searchTerm, languageCode, start)
		}
	}

	if limit != 0 {
		url = fmt.Sprintf("%s&num=%d", url, limit)
	}

	return url
}
