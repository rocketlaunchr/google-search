package googlesearch

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

func ExampleSearch() {

	opt := SearchOptions{
		CountryCode: "au",
	}

	//lint:ignore SA1012 ignore this bare essentials by passing nil for context and removing context package (despite not being idiomatic go).
	serp, err := Search(nil, "First Aid Course Australia Wide First Aid", opt)

	if err != nil {
		fmt.Print(err.Error())
	}

	for _, result := range serp {
		if strings.Contains(result.URL, "australiawidefirstaid.com.au") {
			fmt.Println("Australia Wide First Aid (https://www.australiawidefirstaid.com.au/) found in the serp")
			break
		}
	}

	// Output: Australia Wide First Aid (https://www.australiawidefirstaid.com.au/) found in the serp

}

/*
Example of how to set the useragent
*/
func ExampleUserAgent() {

	// whatismybrowser.com maintains a database of UserAgents
	// https://www.whatismybrowser.com/guides/the-latest-user-agent/chrome

	opt := SearchOptions{
		CountryCode: "au",
		UserAgent:   "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
	}

	//lint:ignore SA1012 ignore this bare essentials by passing nil for context and removing context package (despite not being idiomatic go).
	serp, err := Search(nil, "First Aid Course Australia Wide First Aid", opt)

	if err != nil {
		fmt.Print(err.Error())
	}

	for _, result := range serp {
		if strings.Contains(result.URL, "australiawidefirstaid.com.au") {
			fmt.Println("Australia Wide First Aid (https://www.australiawidefirstaid.com.au/) found in the serp")
			break
		}
	}

	// Output: Australia Wide First Aid (https://www.australiawidefirstaid.com.au/) found in the serp

}

/*
Example of how to set a Rate Limit
*/
func ExampleRateLimit() {

	ctx := context.Background()

	RateLimit.SetLimit(rate.Every(5 * time.Second)) // Interval
	RateLimit.SetBurst(1)                           // Requests per Interval

	err := RateLimit.Wait(ctx)
	if err != nil {
		fmt.Print(err.Error())
	}

	for i := 0; i < 1; i++ {
		serp, err := Search(ctx, "Australia Wide First Aid")

		if err != nil {
			fmt.Print(err.Error())
		}

		if len(serp) > 0 {
			fmt.Println("Resaults found")
		}

	}

	// Output:
	// Resaults found
}

/*
Example of how to get the Related Searches
*/

func ExampleRelatedSearch() {

	opt := SearchOptions{
		CountryCode: "au",
	}

	_, topics, err := RelatedSearch(context.Background(), "Google", opt)

	if err != nil {
		fmt.Print(err.Error())
	}

	for _, topic := range topics {
		if strings.Contains(topic, "google") {
			fmt.Println("Related search found")
			break
		}
	}

	// Output: Related search found

}
