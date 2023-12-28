package yelp

import (
	"fmt"
)

type YelpClient struct {
}

func NewYelpClient() *YelpClient {
	return &YelpClient{}
}

func (y *YelpClient) GetBusinesses(term, location string) {
	fmt.Println("yelp get businesses")
}
