package yelp

import (
	"encoding/json"
)

type YelpSearchResult struct {
	Businesses []struct {
		Name        string  `json:"name"`
		Rating      float64 `json:"rating"`
		Price       string  `json:"price"`
		Description string  `json:"description"`
		Location    struct {
			Address1 string `json:"address1"`
			City     string `json:"city"`
			State    string `json:"state"`
			ZipCode  string `json:"zip_code"`
			Country  string `json:"country"`
		} `json:"location"`
	} `json:"businesses"`
	Total int `json:"total"`
}

func NewYelpSearchResult(body []byte) (*YelpSearchResult, error) {
	var res YelpSearchResult
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
