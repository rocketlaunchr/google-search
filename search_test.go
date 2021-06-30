// Copyright 2020-21 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package googlesearch_test

import (
	"context"
	"testing"

	"github.com/rocketlaunchr/google-search"
)

var ctx = context.Background()

func TestSearch(t *testing.T) {

	q := "Hello World"

	opts := googlesearch.SearchOptions{
		Limit:     20,
		ProxyAddr: "socks://127.0.0.1:7890",
	}

	returnLinks, err := googlesearch.Search(ctx, q, opts)
	if err != nil {
		t.Errorf("something went wrong: %v", err)
		return
	}

	if len(returnLinks) == 0 {
		t.Errorf("no results returned: %v", returnLinks)
	}
}
