// Copyright 2020-22 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package googlesearch

import (
	"errors"

	"golang.org/x/time/rate"
)

// ErrBlocked indicates that Google has detected that you were scraping and temporarily blocked you.
// The duration of the block is unspecified.
//
// See: https://github.com/rocketlaunchr/google-search#warning-warning
var ErrBlocked = errors.New("google block")

// RateLimit sets a global limit to how many requests to Google Search can be made in a given time interval.
// The default is unlimited (but obviously Google Search will block you temporarily if you do too many
// calls too quickly).
//
// See: https://godoc.org/golang.org/x/time/rate#NewLimiter
var RateLimit = rate.NewLimiter(rate.Inf, 0)
