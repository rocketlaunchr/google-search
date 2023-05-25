// Copyright 2020-21 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package googlesearch_test

import (
	"context"
	"testing"

	googlesearch "github.com/chrisjoyce911/google-search"
)

var ctx = context.Background()

func TestSearch(t *testing.T) {

	q := "Hello World"

	opts := googlesearch.SearchOptions{
		Limit: 20,
	}

	returnLinks, err := googlesearch.Search(ctx, q, opts)
	if err != nil {
		t.Errorf("something went wrong: %v", err)
		return
	}

	if len(returnLinks) == 0 {
		t.Errorf("no results returned: %v", returnLinks)
	}

	noURL := 0
	noTitle := 0
	noDesc := 0

	for _, res := range returnLinks {
		if res.URL == "" {
			noURL++
		}

		if res.Title == "" {
			noTitle++
		}

		if res.Description == "" {
			noDesc++
		}
	}

	if noURL == len(returnLinks) || noTitle == len(returnLinks) || noDesc == len(returnLinks) {
		t.Errorf("google dom changed")
	}
}
