package yelp

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
	url := buildURL(fmt.Sprintf("%s%s", API_HOST, SEARCH_PATH), options)
	res := &YelpSearchResult{}
	fmt.Println("url", url)
	httpres, err := y.DoRequest(context.Background(), url, "GET")

	defer httpres.Body.Close()

	body, err := ioutil.ReadAll(httpres.Body)

	if err != nil {
		return &YelpSearchResult{}, err
	}

	res, err = NewYelpSearchResult(body)
	fmt.Println("res", res)
	if err != nil {
		return &YelpSearchResult{}, err
	}
	return res, err
}

func buildURL(base string, options YelpSearchOptions) string {
	encodedLocation := url.QueryEscape(options.Location)
	encodedTerm := url.QueryEscape(options.Term)
	url := fmt.Sprintf("%s%s?location=%s&term=%s", API_HOST, SEARCH_PATH, encodedLocation, encodedTerm)
	return url
}

func (y *YelpClient) DoRequest(ctx context.Context, url string, method string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", y.apiKey))

	return http.DefaultClient.Do(req)
}
