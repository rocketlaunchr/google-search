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
