package googlesearch

import "golang.org/x/time/rate"

// RateLimit sets a global limit to how many requests to Google Search can be made in a given time interval.
// The default is unlimited (but obviously Google Search will block you temporarily if you do too many
// calls too quickly).
//
// See: https://godoc.org/golang.org/x/time/rate#NewLimiter
var RateLimit = rate.NewLimiter(rate.Inf, 0)
