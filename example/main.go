package main

import (
	"context"
	"fmt"

	googlesearch "github.com/chrisjoyce911/google-search"
)

func main() {

	fmt.Println(googlesearch.Search(context.Background(), "First Aid Course Australia Wide First Aid"))
}
