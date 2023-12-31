package yelp

import (
	"testing"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

func init() {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("../config/config_test.yaml")
	if err != nil {
		panic(err)
	}

}

func TestYelpClientSearch(t *testing.T) {
	y := NewYelpClient(config.String("yelp.api_key"))
	res := y.GetBusinesses("food", "San Francisco")
	if res.Total == 0 {
		t.Errorf("Expected non-zero results, got %d", res.Total)
	}
}

func TestYelpClientDoSearch(t *testing.T) {
	y := NewYelpClient(config.String("yelp.api_key"))
	options := YelpSearchOptions{
		Term:     "food",
		Location: "San Francisco",
	}
	res, err := y.DoSearch(options)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if res.Total == 0 {
		t.Errorf("Expected non-zero results, got %d", res.Total)
	}
}
