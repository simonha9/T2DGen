package yelp

import (
	"fmt"

	"github.com/JustinBeckwith/go-yelp/yelp"
)

type YelpClient struct {
	client *yelp.Client
	apiKey string
}

func NewYelpClient(apikey string) *YelpClient {
	options := &yelp.AuthOptions{
		AccessToken: apikey,
	}

	return &YelpClient{
		client: yelp.New(options, nil),
		apiKey: apikey,
	}
}

func (y *YelpClient) GetBusinesses(term, location string) yelp.SearchResult {
	options := yelp.SearchOptions{
		GeneralOptions: &yelp.GeneralOptions{
			Term: term,
		},
		LocationOptions: &yelp.LocationOptions{
			Location: location,
		},
	}
	res, err := y.client.DoSearch(options)
	if err != nil {
		fmt.Println(err)
	}

	return res
}
