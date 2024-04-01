package main

import (
	"context"

	"github.com/nbd-wtf/go-nostr"
)

func rejectEventsFromUsersNotInWhitelist(ctx context.Context, event *nostr.Event) (reject bool, msg string) {
	if isPublicKeyInWhitelist(event.PubKey) {
		return false, ""
	}
	if event.Kind == 1985 {
		// we accept reports from anyone (will filter them for relevance in the next function)
		return false, ""
	}
	return true, "not authorized"
}

var supportedKinds = []uint16{
	0,
	1,
	3,
	5,
	6,
	8,
	16,
	1063,
	1984,
	1985,
	9735,
	10000,
	10001,
	10002,
	10003,
	10004,
	10005,
	10006,
	10007,
	10015,
	10030,
	30000,
	30002,
	30003,
	30004,
	30008,
	30009,
	30015,
	30030,
	30078,
	30311,
	31922,
	31923,
	31924,
	31925,
}

func validateAndFilterReports(ctx context.Context, event *nostr.Event) (reject bool, msg string) {
	if event.Kind == 1985 {
		if e := event.Tags.GetFirst([]string{"e", ""}); e != nil {
			// event report: check if the target event is here
			res, _ := sys.StoreRelay().QuerySync(ctx, nostr.Filter{IDs: []string{(*e)[1]}})
			if len(res) == 0 {
				return true, "we don't know anything about the target event"
			}
		} else if p := event.Tags.GetFirst([]string{"p", ""}); p != nil {
			// pubkey report
			if !isPublicKeyInWhitelist((*p)[1]) {
				return true, "target pubkey is not a user of this relay"
			}
		} else {
			return true, "invalid report"
		}
	}

	return false, ""
}

func removeAuthorsNotWhitelisted(ctx context.Context, filter *nostr.Filter) {
	if n := len(filter.Authors); n > len(whitelist)*11/10 {
		// this query was clearly badly constructed, so we will not bother even looking
		filter.Limit = -1 // this causes the query to be short cut
	} else if n > 0 {
		// otherwise we go through the authors list and remove the irrelevant ones
		newAuthors := make([]string, 0, n)
		for i := 0; i < n; i++ {
			k := filter.Authors[i]
			if _, ok := whitelist[k]; ok {
				newAuthors = append(newAuthors, k)
			}
		}
		filter.Authors = newAuthors

		if len(newAuthors) == 0 {
			filter.Limit = -1 // this causes the query to be short cut
		}
	}
}
