// Copyright 2020-21 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package googlesearch

import (
	"context"
	"testing"
)

var ctx = context.Background()

func TestSearch(t *testing.T) {

	q := "Hello World"

	opts := SearchOptions{
		Limit: 20,
	}

	returnLinks, err := Search(ctx, q, opts)
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

func TestRelatedSearch(t *testing.T) {

	q := "Hello World"

	opts := SearchOptions{
		Limit: 20,
	}

	returnLinks, related, err := RelatedSearch(ctx, q, opts)
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

	if len(related) < 1 {
		t.Errorf("google dom changed")
	}

	if noURL == len(returnLinks) || noTitle == len(returnLinks) || noDesc == len(returnLinks) {
		t.Errorf("google dom changed")
	}
}

func Test_base(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want string
	}{
		{
			name: "Full url",
			url:  "https://www.google.com.au/search?q=",
			want: "https://www.google.com.au/search?q=",
		},
		{
			name: "No base",
			url:  "com.au/search?q=",
			want: "https://www.google.com.au/search?q=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base(tt.url); got != tt.want {
				t.Errorf("base() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildUrl(t *testing.T) {

	tests := []struct {
		name         string
		searchTerm   string
		countryCode  string
		languageCode string
		limit        int
		start        int
		want         string
	}{

		{
			name:         "With Spaces",
			searchTerm:   "Test Example",
			countryCode:  "au",
			languageCode: "en",
			limit:        0,
			start:        0,
			want:         "https://www.google.com.au/search?q=Test+Example&hl=en",
		},

		{
			name:         "Start at 3",
			searchTerm:   "Test Example",
			countryCode:  "au",
			languageCode: "en",
			limit:        0,
			start:        3,
			want:         "https://www.google.com.au/search?q=Test+Example&hl=en&start=3",
		},

		{
			name:         "Non countryCode",
			searchTerm:   "Test Example",
			countryCode:  "xx",
			languageCode: "en",
			limit:        0,
			start:        0,
			want:         "https://www.google.com/search?q=Test+Example&hl=en",
		},
		{
			name:         "Non countryCode Start at 3",
			searchTerm:   "Test Example",
			countryCode:  "xx",
			languageCode: "en",
			limit:        0,
			start:        3,
			want:         "https://www.google.com/search?q=Test+Example&hl=en&start=3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildUrl(tt.searchTerm, tt.countryCode, tt.languageCode, tt.limit, tt.start); got != tt.want {
				t.Errorf("buildUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getStart(t *testing.T) {

	tests := []struct {
		name string
		uri  string
		want int
	}{
		{
			name: "No start set",
			uri:  "https://www.google.com/search?q=Test+Example&hl=en",
			want: 0,
		},
		{
			name: "Start at t",
			uri:  "https://www.google.com/search?q=Test+Example&hl=en&start=3",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStart(tt.uri); got != tt.want {
				t.Errorf("getStart() = %v, want %v", got, tt.want)
			}
		})
	}
}
