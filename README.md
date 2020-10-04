<p align="right">
  <a href="http://godoc.org/github.com/rocketlaunchr/google-search"><img src="http://godoc.org/github.com/rocketlaunchr/google-search?status.svg" /></a>
  <a href="https://goreportcard.com/report/github.com/rocketlaunchr/google-search"><img src="https://goreportcard.com/badge/github.com/rocketlaunchr/google-search" /></a>
</p>

<p align="center">
<img src="https://github.com/rocketlaunchr/google-search/raw/master/screenshot.png" alt="google home page" />
</p>

Quickly scrape Google Search Results.

⭐ **the project to show your appreciation.**

## Example

```go
package main

import (
	"context"
	"fmt"
	"github.com/rocketlaunchr/google-search"
)

func main() {
	ctx := context.Background()
	fmt.Println(googlesearch.Search(ctx, "cars for sale in Toronto, Canada"))
}
```

## Results:

```go
([]googlesearch.Result) (len=11 cap=16) {
 (googlesearch.Result) {
  Rank: (int) 1,
  URL: (string) (len=42) "https://www.autotrader.ca/cars/on/toronto/",
  Title: (string) (len=51) "New & Used Cars for sale in Toronto | autoTRADER.ca",
  Description: (string) ""
 },
 (googlesearch.Result) {
  Rank: (int) 2,
  URL: (string) (len=42) "https://www.autotrader.ca/cars/on/toronto/",
  Title: (string) (len=51) "New & Used Cars for sale in Toronto | autoTRADER.ca",
  Description: (string) ""
 },
 (googlesearch.Result) {
  Rank: (int) 3,
  URL: (string) (len=50) "https://www.carpages.ca/ontario/toronto/used-cars/",
  Title: (string) (len=31) "Used Cars Toronto | Carpages.ca",
  Description: (string) (len=337) "13518 results - Used Cars, Trucks and SUVs for Sale in Toronto, ON. 2009 Acura TL. AWD, Navi, Camera, Leather, 3/Y warranty availabl. 2010 Chevrolet Traverse. 2LT. 2017 Jaguar F-PACE. 35T-AWD-NAVI-CAMERA-PANO ROOF-CPO WARRANTY. 2005 Audi A6. $2,495. 2007 Audi A4. 2.0T. 2012 Audi Q7. 3.0L Premium Plus. 2005 Ford F-250. 2010 Nissan Cube."
 }
}
```

## Warning

The implementation relies on Google's search page DOM being constant. From time to time, Google changes their DOM and thus breaks the implementation.

In the event it changes, this package will be updated as soon as possible.

Also note, that if you call this function too quickly, Google detects that it is being scraped and produces a [recaptcha](https://www.google.com/recaptcha/intro/v3.html) which interferes with the scraping. **Don't call it in quick succession.**



## Credits

Special thanks to [Edmund Martin](https://edmundmartin.com/scraping-google-with-golang/).


Other useful packages
------------

- [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) - Statistics and data manipulation
- [dbq](https://github.com/rocketlaunchr/dbq) - Zero boilerplate database operations for Go
- [electron-alert](https://github.com/rocketlaunchr/electron-alert) - SweetAlert2 for Electron Applications
- [igo](https://github.com/rocketlaunchr/igo) - A Go transpiler with cool new syntax such as fordefer (defer for for-loops)
- [mysql-go](https://github.com/rocketlaunchr/mysql-go) - Properly cancel slow MySQL queries
- [react](https://github.com/rocketlaunchr/react) - Build front end applications using Go
- [remember-go](https://github.com/rocketlaunchr/remember-go) - Cache slow database queries
