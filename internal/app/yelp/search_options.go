package yelp

type YelpSearchOptions struct {
	Term     string
	Location string
}

func NewYelpSearchOptions(term, location string) *YelpSearchOptions {
	return &YelpSearchOptions{term, location}
}
