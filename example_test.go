package googlesearch

import (
	"context"
	"fmt"
	"strings"
)

func ExampleSearch() {

	opt := SearchOptions{
		CountryCode: "au",
	}

	serp, err := Search(context.Background(), "First Aid Course Australia Wide First Aid", opt)

	if err != nil {
		fmt.Print(err.Error())
	}

	for _, resault := range serp {
		if strings.Contains(resault.URL, "australiawidefirstaid.com.au") {
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

	serp, err := Search(context.Background(), "First Aid Course Australia Wide First Aid", opt)

	if err != nil {
		fmt.Print(err.Error())
	}

	for _, resault := range serp {
		if strings.Contains(resault.URL, "australiawidefirstaid.com.au") {
			fmt.Println("Australia Wide First Aid (https://www.australiawidefirstaid.com.au/) found in the serp")
			break
		}
	}

	// Output: Australia Wide First Aid (https://www.australiawidefirstaid.com.au/) found in the serp

}
