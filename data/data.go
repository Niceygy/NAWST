package data

import (
	"io"
	"log"
)

var LoadedEntries []WebViewEntry

func CreateNewEntry(body io.ReadCloser) {
	new := WebViewEntry{}
	new.FromJSONBody(body)
	LoadedEntries = append(LoadedEntries, new)
	log.Printf("Got new entry:\nDomain=%s\nPath=%s, Country=%s, HasVisitedBefore=%t", new.domain, new.path, new.country, new.hasVisitedBefore)
}

func LoadEntriesForDomain(domain string) []WebViewEntry {
	var result []WebViewEntry

	for _, e := range LoadedEntries {
		if e.domain == domain {
			result = append(result, e)
		}
	}

	return result
}
