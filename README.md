# google-search [![GoDoc](http://godoc.org/github.com/rocketlaunchr/google-search?status.svg)](http://godoc.org/github.com/rocketlaunchr/google-search) [![Go Report Card](https://goreportcard.com/badge/github.com/rocketlaunchr/google-search)](https://goreportcard.com/report/github.com/rocketlaunchr/google-search)

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
 },
 (googlesearch.Result) {
  Rank: (int) 4,
  URL: (string) (len=73) "https://www.kijiji.ca/b-cars-trucks/gta-greater-toronto-area/c174l1700272",
  Title: (string) (len=57) "Great Deals on New or Used Cars and Trucks Near Me in ...",
  Description: (string) (len=162) "Find new, used and salvaged cars & trucks for sale locally in Toronto (GTA) : Toyota, Honda, BMW, Mercedes Benz, Chrysler, Nissan and it is all about driving\u00a0..."
 },
 (googlesearch.Result) {
  Rank: (int) 5,
  URL: (string) (len=59) "https://ca.cargurus.com/Cars/spt_cheap_cars-Toronto_L414276",
  Title: (string) (len=45) "Cheap Cars For Sale in Toronto, ON - CarGurus",
  Description: (string) (len=146) "Listings 1 - 15 of 2747 - Search used cheap cars listings to find the best Toronto, ON deals. We analyze hundreds of thousands of used cars daily."
 },
 (googlesearch.Result) {
  Rank: (int) 6,
  URL: (string) (len=50) "http://autocatch.com/toronto-ontario/used-cars.htm",
  Title: (string) (len=46) "New and Used Cars For Sale Toronto | AutoCatch",
  Description: (string) (len=157) "Find your next new and used cars in Toronto on AutoCatch.com. Search over 1415 new and used vehicle listings available in Toronto, Ontario and find your\u00a0..."
 },
 (googlesearch.Result) {
  Rank: (int) 7,
  URL: (string) (len=49) "https://www.ucda.ca/find-used-car-near-me/toronto",
  Title: (string) (len=24) "Used Cars Toronto - UCDA",
  Description: (string) (len=138) "With over 5000 member dealerships and more than 20,000 used cars for sales in Toronto, it's easy to find a perfect vehicle. All Makes\u00a0..."
 },
 (googlesearch.Result) {
  Rank: (int) 8,
  URL: (string) (len=34) "https://www.torontohonda.com/used/",
  Title: (string) (len=59) "Used Cars, SUVs, Trucks for Sale in Toronto | Toronto Honda",
  Description: (string) (len=136) "Shop our selection of used vehicles for sale at our dealership in Toronto. Contact us today to apply for financing or book a test drive!"
 },
 (googlesearch.Result) {
  Rank: (int) 9,
  URL: (string) (len=60) "https://www.hertzcarsalestoronto.com/all-inventory/index.htm",
  Title: (string) (len=46) "Search All Used Cars - Hertz Car Sales Toronto",
  Description: (string) (len=165) "Browse our extensive online inventory of affordable used cars that are available in Toronto. Our certified vehicles are sold with a warranty. Shop our inventory\u00a0..."
 },
 (googlesearch.Result) {
  Rank: (int) 10,
  URL: (string) (len=126) "https://www.auto123.com/en/used-cars/all-inventories/ontario/toronto/all-years/all-makes/all-models/all-bodytypes/all-sellers/",
  Title: (string) (len=59) "Used cars for sale in Toronto - Second hand vehicles in ...",
  Description: (string) (len=139) "111 results - Used vehicles for sale in Toronto in the car classifieds of Auto123.com, the most comprehensive automotive website in Canada."
 },
 (googlesearch.Result) {
  Rank: (int) 11,
  URL: (string) (len=24) "https://www.autorama.ca/",
  Title: (string) (len=48) "Used Car Dealership in Toronto, GTA & North York",
  Description: (string) (len=147) "Specializing in Japanese Used Cars and SUVs, used European Import Cars and ... cars for sale come with a FREE CARFAX Canada Vehicle History Report."
 }
}
```

## Warning

The implementation relies on Google's search page DOM being constant. If it changes, then the implementation will break. Having said that, it hasn't changed for years.

In the event it changes, this package will be updated as soon as possible.

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
