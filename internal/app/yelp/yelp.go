package yelp

import (
	"fmt"
	"context"
	"net/http"
	"encoding/json"
	"net/url"
)

var (
	API_HOST      = "https://api.yelp.com"
	SEARCH_PATH   = "/v3/businesses/search"
	BUSINESS_PATH = "/v3/businesses/"
)

type YelpClient struct {
	apiKey string
}

func NewYelpClient(apikey string) *YelpClient {
	return &YelpClient{apikey}
}

func (y *YelpClient) GetBusinesses(term, location string) YelpSearchResult {
	options := YelpSearchOptions{
		Term:     term,
		Location: location,
	}
	res, err := y.DoSearch(options)
	if err != nil {
		fmt.Println(err)
	}

	return res
}

func (y *YelpClient) DoSearch(options YelpSearchOptions) (YelpSearchResult, error) {
	url := fmt.Sprintf("%s%s", API_HOST, SEARCH_PATH)
	res := YelpSearchResult{}
	err := Get(url, y.apiKey, options, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
