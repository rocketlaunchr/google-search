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
	"fmt"
	"github.com/rocketlaunchr/google-search"
)

func main() {
	fmt.Println(googlesearch.Search(nil, "cars for sale in Toronto, Canada"))
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

## :warning: Warning

The implementation relies on Google's search page DOM being constant. From time to time, Google changes their DOM and thus breaks the implementation.

In the event it changes, this package will be updated as soon as possible.

Also note, that if you call this function too quickly, Google detects that it is being scraped and produces a [recaptcha](https://www.google.com/recaptcha/intro/v3.html) which interferes with the scraping. **Don't call it in quick succession.** It may take some time before Google unblocks you.

You can try the built-in [rate-limiter](https://godoc.org/github.com/rocketlaunchr/google-search#RateLimit).

<details>
  <summary>Further Details</summary>
  
<svg width="100" height="100" xmlns="http://www.w3.org/2000/svg">
<foreignObject width="100" height="100">
  <div xmlns="http://www.w3.org/1999/xhtml">
  <div style="font-family: arial, sans-serif; background-color: #fff; color: #000; padding:20px; font-size:18px;" onload="e=document.getElementById('captcha');if(e){e.focus();}">
  <div style="max-width:400px;">


  <div style="font-size:13px;">
  [HTTP STATUS CODE: 429 &mdash; Too Many Requests]
  <b>About this page</b><br><br>

  Our systems have detected unusual traffic from your computer network.  This page checks to see if it&#39;s really you sending the requests, and not a robot.  <a href="#" onclick="document.getElementById('infoDiv').style.display='block';">Why did this happen?</a><br><br>

  <div id="infoDiv" style="display:none; background-color:#eee; padding:10px; margin:0 0 15px 0; line-height:1.4em;">
  This page appears when Google automatically detects requests coming from your computer network which appear to be in violation of the <a href="//www.google.com/policies/terms/">Terms of Service</a>. The block will expire shortly after those requests stop.  In the meantime, solving the above CAPTCHA will let you continue to use our services.<br><br>This traffic may have been sent by malicious software, a browser plug-in, or a script that sends automated requests.  If you share your network connection, ask your administrator for help &mdash; a different computer using the same IP address may be responsible.  <a href="//support.google.com/websearch/answer/86640">Learn more</a><br><br>Sometimes you may be asked to solve the CAPTCHA if you are using advanced terms that robots are known to use, or sending requests very quickly.
  </div>

  IP address: xxx.xx.xxx.xx<br>Time: 2021-01-13T05:27:34Z<br>URL: https://www.google.com/search?q=Hello+World&amp;hl=en&amp;num=20<br>
  </div>
  </div>
  </div>
  </div>
</foreignObject>
</svg>
</details>


## Credits

Special thanks to [Edmund Martin](https://edmundmartin.com/scraping-google-with-golang/).


Other useful packages
------------

- [awesome-svelte](https://github.com/rocketlaunchr/awesome-svelte) - Resources for killing react
- [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) - Statistics and data manipulation
- [dbq](https://github.com/rocketlaunchr/dbq) - Zero boilerplate database operations for Go
- [electron-alert](https://github.com/rocketlaunchr/electron-alert) - SweetAlert2 for Electron Applications
- [igo](https://github.com/rocketlaunchr/igo) - A Go transpiler with cool new syntax such as fordefer (defer for for-loops)
- [mysql-go](https://github.com/rocketlaunchr/mysql-go) - Properly cancel slow MySQL queries
- [react](https://github.com/rocketlaunchr/react) - Build front end applications using Go
- [remember-go](https://github.com/rocketlaunchr/remember-go) - Cache slow database queries
- [testing-go](https://github.com/rocketlaunchr/testing-go) - Testing framework for unit testing
