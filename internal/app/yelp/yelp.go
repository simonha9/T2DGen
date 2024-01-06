package yelp

import (
	"context"
	"fmt"
	"net/http"
)

const (
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

func (y *YelpClient) GetBusinesses(term, location string) *YelpSearchResult {
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

func (y *YelpClient) DoSearch(options YelpSearchOptions) (*YelpSearchResult, error) {
	url := fmt.Sprintf("%s%s", API_HOST, SEARCH_PATH)
	res := &YelpSearchResult{}
	httpres, err := y.DoRequest(context.Background(), url, "GET")
	if err != nil {
		res, err = NewYelpSearchResult(httpres)
		if err != nil {
			return &YelpSearchResult{}, err
		}
	}
	return res, err
}

func (y *YelpClient) DoRequest(ctx context.Context, url string, method string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", y.apiKey))

	return http.DefaultClient.Do(req)
}
